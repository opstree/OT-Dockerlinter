package rules

type Rule struct {
	RegularExpression string
	Description       string
	Level             string
	Recommendation    string
}

var DockerFileRules = []Rule{
	{
		RegularExpression: ".* sudo .*",
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
	{
		RegularExpression: "USER root",
		Description:       "Root user is being used",
		Level:             "FATAL",
		Recommendation:    "Root user should not be used in Dockerfiles",
	},
	{
		RegularExpression: ".* cd .*",
		Description:       "cd command is getting used to change directory",
		Level:             "WARN",
		Recommendation:    "Use WORKDIR instead of cd command for changing the directory",
	},
	{
		RegularExpression: "(.* update|.*upgrade)",
		Description:       "apt or yum command is getting for update and upgrade",
		Level:             "FATAL",
		Recommendation:    "Don't use update and upgrade command in RUN step",
	},
}
