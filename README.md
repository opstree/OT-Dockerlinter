<p align="left">
  <img src="./static/ot-dockerlinter.svg" height="150" width="150">
</p>

OT-Dockerlinter helps you in writing a Dockerfile with best practices. This tools can be integrated with your container native CI pipeline for Dockerfile's static code analysis and reporting.

### Supported Features

- Dockefile linting and reporting in different formats like table, json
- Integration with Jenkins(In Development Mode)
- Cross platform support is available
- Dockerfile best practices and recommendation

## Quickstart

A quickstart guide for installing, using and managing OT-Dockerlinter.

### Installation

OT-Dockerlinter installation packages can be found inside [Releases](https://github.com/opstree/OT-Dockerlinter/releases)

Supported Platforms:-

- Linux and Windows Platform with supported architecture types:-
  - AMD64
  - ARM63
  - AMD32Bit
  - ARM32Bit

For installation on debian and redhat based system, `.deb` and `.rpm` packages can be used.

### Available Options

There are help page available for ot-dockerlinter which can be called by `help` or `--help` command.

```shell
A tool for checking Dockerfile best practices.

Usage:
  ot-docker-linter [flags]
  ot-docker-linter [command]

Available Commands:
  audit       Runs ot-docker-linter audit
  help        Help about any command
  version     Prints the current version.

Flags:
  -h, --help                help for ot-docker-linter
      --log.format string   ot-docker-linter log format. (default "text")
      --log.level string    ot-docker-linter logging level. (default "info")

Use "ot-docker-linter [command] --help" for more information about a command.
```
