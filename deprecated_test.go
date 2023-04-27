package checkdeprecated

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckDeprecated(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewCheckDeprecatedAnalyzer(), "deprecated")
}

func TestCheckDeprecatedComment(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewCheckDeprecatedCommentAnalyzer(), "deprecated_comment")
}
