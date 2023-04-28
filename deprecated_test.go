package checkdeprecated

import (
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckDeprecated(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewCheckDeprecatedAnalyzer(), "deprecated")
}

func TestCheckDeprecatedComment(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewCheckDeprecatedCommentAnalyzer(), "deprecated_comment")
}

func TestCheck(t *testing.T) {
	checkers := []*analysis.Analyzer{
		NewCheckDeprecatedAnalyzer(),
		NewCheckDeprecatedCommentAnalyzer(),
	}
	err := analysis.Validate(checkers)
	if err != nil {
		t.Errorf("validate failed, got err: %v", err)
	}
}
