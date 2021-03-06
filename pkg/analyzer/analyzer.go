package linter

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"

	"dockerfile-inspector/pkg/rule"
)

// Analyzer implements Analyzer.
type Analyzer struct {
	rules []*rules.Rule
}

// NewAnalyzer generate a NewAnalyzer with rules to apply.
func NewAnalyzer(ignoreRules []string) Analyzer {
	return newAnalyzer(ignoreRules)
}

func newAnalyzer(ignoreRules []string) Analyzer {
	var filteredRules []*rules.Rule
	for _, k := range getMakeDiff(rules.RuleKeys, ignoreRules) {
		if rule, ok := rules.Rules[k]; ok {
			filteredRules = append(filteredRules, rule)
		}
	}
	return Analyzer{rules: filteredRules}
}

// Run apply docker best practice rules to docker ast.
func (a Analyzer) Run(node *parser.Node, fileName string) ([]rules.Result, error) {
	var rst []rules.Result
	rstChan := make(chan []rules.Result, len(a.rules))
	errChan := make(chan error, len(a.rules))

	for i := range a.rules {
		go func(r *rules.Rule) {
			vrst, err := r.ValidateFunc(node)
			if err != nil {
				errChan <- err
			} else {
				rstChan <- rules.CreateMessage(a.rules[i], vrst, fileName)
			}
		}(a.rules[i])
		select {
		case value := <-rstChan:
			rst = append(rst, value...)
		case err := <-errChan:
			return nil, err
		}
	}
	return rst, nil
}

// getMakeDifference is a function to create a difference set.
func getMakeDiff(xs, ys []string) []string {
	if len(xs) > len(ys) {
		return makeDiff(xs, ys)
	}
	return makeDiff(ys, xs)
}

// make set difference.
func makeDiff(xs, ys []string) []string {
	var set []string
	for i := range xs {
		if !isContain(ys, xs[i]) {
			set = append(set, xs[i])
		}
	}
	return set
}

// isContain is a function to check if s is in xs.
func isContain(xs []string, s string) bool {
	for i := range xs {
		if xs[i] == s {
			return true
		}
	}
	return false
}
