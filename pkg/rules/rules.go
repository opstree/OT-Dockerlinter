package rules

type Rule struct {
	RegularExpression string
	Description       string
	Level             string
	Recommendation    string
}

var DockerFileRules = []Rule{
	{
		RegularExpression: "RUN sudo apt-get update",
		Description:       "sudo command is being used",
		Level:             "FATAL",
		Recommendation:    "Avoid using sudo commands, use gosu instead",
	},
	{
		RegularExpression: "FROM .*:latest",
		Description:       "Latest tag is being used",
		Level:             "WARN",
		Recommendation:    "Avoid using latest tag in Dockerfiles",
	},
}
