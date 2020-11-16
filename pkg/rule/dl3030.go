package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3030 Use the -y switch to avoid manual input yum install -y <package>
func validateDL3030(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isYum, isInstalled bool
			length := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "yum":
					isYum = true
				case "install":
					if isYum {
						isInstalled = true
					}
				case "&&":
					isYum, isInstalled = false, false
				case "-y":
					rst = append(rst, ValidateResult{line: child.StartLine})
				default:
					if isInstalled && !yesPattern.MatchString(v) && length == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
					}
					isYum, isInstalled = false, false
				}
			}
		}
	}
	return rst, nil
}
