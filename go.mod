module dockerfile-inspector

go 1.15

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23 // indirect
	github.com/landoop/tableprinter v0.0.0-20200805134727-ea32388e35c1
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/moby/buildkit v0.6.3
	github.com/olekukonko/tablewriter v0.0.4
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tidwall/pretty v1.0.2
	github.com/xlab/tablewriter v0.0.0-20160610135559-80b567a11ad5
)

replace (
	github.com/containerd/containerd v1.3.0-0.20190507210959-7c1e88399ec0 => github.com/containerd/containerd v1.3.0
	github.com/docker/docker v1.14.0-0.20190319215453-e7b5f7dbe98c => github.com/docker/docker v1.4.2-0.20200227233006-38f52c9fec82
	golang.org/x/crypto v0.0.0-20190129210102-0709b304e793 => golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
)
