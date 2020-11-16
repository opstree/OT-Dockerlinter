package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3031 is "Do not use yum update"
func validateDL3031(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isYum, isUpgrade bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "yum":
					isYum = true
				case "update":
					isUpgrade = true
				}
			}
			if isYum && isUpgrade {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
