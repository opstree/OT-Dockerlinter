package rules

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3029 Do not use --platform= with FROM
func validateDL3029(node *parser.Node) (rst []ValidateResult, err error) {
	regexDL3029 := regexp.MustCompile(`--platform`)
	for _, child := range node.Children {
		if child.Value == FROM && regexDL3029.MatchString(child.Next.Value) {
			rst = append(rst, ValidateResult{line: child.StartLine})
		}
	}
	return rst, nil
}
