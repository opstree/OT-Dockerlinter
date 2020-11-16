package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3032 is "yum clean all missing after yum command."
func validateDL3032(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isYum, isClean bool
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "yum":
					isYum = true
				case "clean":
					isClean = true
				}
			}
			if isYum && !isClean {
				rst = append(rst, ValidateResult{line: child.StartLine})
			}
		}
	}
	return rst, nil
}
