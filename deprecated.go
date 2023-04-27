package checkdeprecated

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type CheckDeprecated struct {
	marker *analysis.Analyzer
}

func NewCheckDeprecatedAnalyzer(patterns ...string) *analysis.Analyzer {
	marker := NewMarkDeprecatedCommentAnalyzer(patterns...)
	return &analysis.Analyzer{
		Name: "check_deprecated",
		Doc:  "Using a deprecated function, variable, constant or field",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			deprecatedMap := pass.ResultOf[marker].(DeprecatedMap)
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
		Requires: []*analysis.Analyzer{inspect.Analyzer, marker},
	}
}
