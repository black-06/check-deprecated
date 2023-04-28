package checkdeprecated

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const standard = "Deprecated: "

var (
	commonPatterns = []string{
		"deprecated",
		"it's deprecated",
		"this type is deprecated",
		"this function is deprecated",
		"[[deprecated]]",
		"note: deprecated",
	}
	MarkDeprecatedCommentAnalyzer = &analysis.Analyzer{
		Name:       "mark_deprecated_comment",
		Doc:        "mark deprecated comment",
		Run:        deprecatedCommentRun,
		Flags:      deprecatedCommentFlags(),
		Requires:   []*analysis.Analyzer{inspect.Analyzer},
		ResultType: reflect.TypeOf(DeprecatedMap{}),
		FactTypes:  []analysis.Fact{&Deprecated{}},
	}
)

type DeprecatedMap map[types.Object]*Deprecated

type Deprecated struct {
	Message         string
	MalformedHeader string
}

func (*Deprecated) AFact() {}

func (p *Deprecated) String() string {
	return "message: " + p.Message + " ,malformed_header: " + p.MalformedHeader
}

func deprecatedCommentFlags() flag.FlagSet {
	options := flag.NewFlagSet("", flag.ExitOnError)
	options.String("patterns", "", "custom deprecated comment header")
	return *options
}

func setPatterns(f *flag.Flag, patterns []string) error {
	bytes, err := json.Marshal(patterns)
	if err != nil {
		return err
	}
	return f.Value.Set(string(bytes))
}

func parsePatterns(f *flag.Flag) ([]string, error) {
	bytes := []byte(f.Value.String())
	var patterns []string
	if err := json.Unmarshal(bytes, &patterns); err != nil {
		return nil, err
	}
	patterns = append(patterns, commonPatterns...)
	return patterns, nil
}

// deprecatedCommentRun finds all 'deprecated' comment
func deprecatedCommentRun(pass *analysis.Pass) (interface{}, error) { //nolint:gocyclo
	patterns, err := parsePatterns(pass.Analyzer.Flags.Lookup("patterns"))
	if err != nil {
		return nil, err
	}

	pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).Preorder(nil, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.FuncDecl:
			if deprecated := extract(patterns, node.Doc); deprecated != nil {
				pass.ExportObjectFact(pass.TypesInfo.ObjectOf(node.Name), deprecated)
			}
		case *ast.TypeSpec:
			if deprecated := extract(patterns, node.Doc); deprecated != nil {
				pass.ExportObjectFact(pass.TypesInfo.ObjectOf(node.Name), deprecated)
			}
		case *ast.Field:
			if deprecated := extract(patterns, node.Doc); deprecated != nil {
				for _, name := range node.Names {
					pass.ExportObjectFact(pass.TypesInfo.ObjectOf(name), deprecated)
				}
			}
		case *ast.ValueSpec:
			if deprecated := extract(patterns, node.Doc); deprecated != nil {
				for _, name := range node.Names {
					pass.ExportObjectFact(pass.TypesInfo.ObjectOf(name), deprecated)
				}
			}
		case *ast.ImportSpec:
			if deprecated := extract(patterns, node.Doc); deprecated != nil {
				pass.ExportObjectFact(pass.TypesInfo.ObjectOf(node.Name), deprecated)
			}
		case *ast.GenDecl:
			if node.Tok == token.VAR || node.Tok == token.CONST || node.Tok == token.TYPE {
				if deprecated := extract(patterns, node.Doc); deprecated != nil {
					for _, spec := range node.Specs {
						switch s := spec.(type) {
						case *ast.ValueSpec:
							for _, name := range s.Names {
								pass.ExportObjectFact(pass.TypesInfo.ObjectOf(name), deprecated)
							}
						case *ast.TypeSpec:
							pass.ExportObjectFact(pass.TypesInfo.ObjectOf(s.Name), deprecated)
						}
					}
				}
			}
		}
	})

	deprecatedMap := make(DeprecatedMap)
	for _, fact := range pass.AllObjectFacts() {
		deprecatedMap[fact.Object] = fact.Fact.(*Deprecated)
	}
	return deprecatedMap, nil
}

func extract(patterns []string, doc *ast.CommentGroup) *Deprecated {
	if doc == nil {
		return nil
	}
	for _, part := range strings.Split(doc.Text(), "\n\n") {
		if deprecated := extractFromStr(patterns, part); deprecated != nil {
			return deprecated
		}

		for _, line := range strings.Split(part, "\n") {
			if deprecated := extractFromStr(patterns, line); deprecated != nil {
				if deprecated.MalformedHeader == "" {
					deprecated.MalformedHeader = "`Deprecated: ` should be at the beginning of the paragraph"
				}
				return deprecated
			}
		}
	}
	return nil
}

func extractFromStr(patterns []string, str string) *Deprecated {
	if hasPrefixFold(str, standard) {
		prefix := str[:len(standard)]
		deprecated := Deprecated{
			Message:         clear(str[len(standard):]),
			MalformedHeader: "",
		}
		if deprecated.Message == "" {
			deprecated.Message = "(no comment)"
			deprecated.MalformedHeader = "it's deprecated, but no comment"
		}
		if prefix != standard {
			deprecated.MalformedHeader = fmt.Sprintf("use `Deprecated: ` (note the casing) instead of `%s`", prefix)
		}
		return &deprecated
	}
	for _, pattern := range patterns {
		if hasPrefixFold(str, pattern) {
			msg := clear(str[len(pattern):])
			if msg == "" {
				msg = "(no comment)"
			}
			return &Deprecated{Message: msg, MalformedHeader: "the proper format is `Deprecated: <text>`"}
		}
	}
	return nil
}

func clear(part string) string {
	runes := []rune(strings.ReplaceAll(part, "\n", " "))
	for len(runes) > 0 && (unicode.IsPunct(runes[0]) || unicode.IsSpace(runes[0])) {
		runes = runes[1:]
	}
	return string(runes)
}

func hasPrefixFold(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[:len(prefix)], prefix)
}
