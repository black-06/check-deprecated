package checkdeprecated

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewCheckDeprecatedAnalyzer(patterns ...string) *analysis.Analyzer {
	_ = setPatterns(MarkDeprecatedCommentAnalyzer.Flags.Lookup("patterns"), patterns)
	return &analysis.Analyzer{
		Name: "check_deprecated",
		Doc:  "Using a deprecated function, variable, constant or field",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			deprecatedMap := pass.ResultOf[MarkDeprecatedCommentAnalyzer].(DeprecatedMap)
			pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).Preorder(
				[]ast.Node{&ast.Ident{}}, func(n ast.Node) {
					node := n.(*ast.Ident)
					obj := pass.TypesInfo.ObjectOf(node)
					if obj == nil || obj.Pos() == node.Pos() {
						return
					}
					if deprecated, ok := deprecatedMap[obj]; ok {
						pass.Reportf(node.Pos(), "using deprecated: %s", deprecated.Message)
					}
				},
			)
			return nil, nil
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer, MarkDeprecatedCommentAnalyzer},
	}
}

func NewCheckDeprecatedCommentAnalyzer(patterns ...string) *analysis.Analyzer {
	_ = setPatterns(MarkDeprecatedCommentAnalyzer.Flags.Lookup("patterns"), patterns)
	return &analysis.Analyzer{
		Name: "check_deprecated_comment",
		Doc:  "Check malformed deprecated comment",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			deprecatedMap := pass.ResultOf[MarkDeprecatedCommentAnalyzer].(DeprecatedMap)
			for object, deprecated := range deprecatedMap {
				if deprecated.MalformedHeader != "" {
					pass.Reportf(object.Pos(), "malformed deprecated header: %s", deprecated.MalformedHeader)
				}
			}
			return nil, nil
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer, MarkDeprecatedCommentAnalyzer},
	}
}
