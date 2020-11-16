package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

var regexDL3033 = regexp.MustCompile(`.+-.+`)

// validateDL3033 Specify version with yum install -y <package>-<version>
func validateDL3033(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			var isYum, isInstall bool
			l := len(rst)
			for _, v := range strings.Fields(child.Next.Value) {
				fmt.Println(v)
				switch v {
				case "yum":
					isYum = true
				case "install":
					if isYum {
						isInstall = true
					}
				case "&&":
					isYum, isInstall = false, false
					continue
				case "-y":
					isYum, isInstall = false, false
					continue
				case "update":
					isYum, isInstall = false, false
					continue
				case "upgrade":
					isYum, isInstall = false, false
					continue
				default:
					if isInstall && !regexDL3033.MatchString(v) && l == len(rst) {
						rst = append(rst, ValidateResult{line: child.StartLine})
						isYum, isInstall = false, false
					}
				}
			}
		}
	}
	return rst, nil
}
