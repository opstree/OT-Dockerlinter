package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3027 Do not use apt as it is meant to be an end-user tool, use apt-get or apt-cache instead.
func validateDL3027(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isApt bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt":
					isApt = true
				default:
					if isApt && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isApt = false
				}
			}
		}
	}
	return rst, nil
}
