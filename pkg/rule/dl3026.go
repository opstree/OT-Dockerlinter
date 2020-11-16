package rules

import (
	"os"
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3026 Use only an allowed registry in the FROM image
func validateDL3026(node *parser.Node) (rst []ValidateResult, err error) {
	trustedRegistry := os.Getenv("TRUSTED_REGISTRY")
	if trustedRegistry != "" {
		regexDL3026 := regexp.MustCompile(trustedRegistry + ".*")
		for _, child := range node.Children {
			if child.Value == FROM && regexDL3026.MatchString(child.Next.Value) != true {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
