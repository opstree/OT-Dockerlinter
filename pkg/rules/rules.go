package rules

type Rule struct {
	RegularExpression string
	Description       string
	Level             string
	Recommendation    string
}

var DockerFileRules = []Rule{
	{
		RegularExpression: "FROM .*:latest",
		Description:       "Latest tag is being used",
		Level:             "WARN",
		Recommendation:    "Avoid using latest tag in Dockerfiles",
	},
}
