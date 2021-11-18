# mvisonneau/mmds - Missed (AWS) Meta-Data (service)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mvisonneau/mmds)](https://pkg.go.dev/mod/github.com/mvisonneau/mmds)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/mmds)](https://goreportcard.com/report/github.com/mvisonneau/mmds)
[![test](https://github.com/mvisonneau/mmds/actions/workflows/test.yml/badge.svg)](https://github.com/mvisonneau/mmds/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/mvisonneau/mmds/badge.svg?branch=main)](https://coveralls.io/github/mvisonneau/mmds?branch=main)
[![release](https://github.com/mvisonneau/mmds/actions/workflows/release.yml/badge.svg)](https://github.com/mvisonneau/mmds/actions/workflows/release.yml)

`mmds` allows you to get some information which are currently missing from the [AWS meta-data service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)

## TL:DR

```bash
# Figure out the pricing model of the instance
~$ mmds pricing-model
spot
```

## Usage

```bash
~$ mmds
NAME:
   mmds - Missed (AWS) Meta-Data (service)

USAGE:
   mmds [global options] command [command options] [arguments...]

COMMANDS:
   pricing-model  get instance pricing-model
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level level    log level (debug,info,warn,fatal,panic) (default: "info") [$MMDS_LOG_LEVEL]
   --log-format format  log format (json,text) (default: "text") [$MMDS_LOG_FORMAT]
   --help, -h           show help (default: false)
```

## Install

Have a look onto the [latest release page](https://github.com/mvisonneau/mmds/releases/latest) and pick your flavor.

### Go

```bash
~$ go install github.com/mvisonneau/mmds/cmd/mmds@latest
```

### Homebrew (linux only)

```bash
~$ brew install mvisonneau/tap/mmds
```

### Scoop

```bash
~$ scoop bucket add https://github.com/mvisonneau/scoops
~$ scoop install mmds
```

### Binaries, DEB and RPM packages

For the following ones, you need to know which version you want to install, to fetch the latest available :

```bash
~$ export MMDS_VERSION=$(curl -s "https://api.github.com/repos/mvisonneau/mmds/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
```

```bash
# Binary (eg: freebsd/amd64)
~$ wget https://github.com/mvisonneau/mmds/releases/download/${MMDS_VERSION}/mmds_${MMDS_VERSION}_freebsd_amd64.tar.gz
~$ tar zxvf mmds_${MMDS_VERSION}_freebsd_amd64.tar.gz -C /usr/local/bin

# DEB package (eg: linux/386)
~$ wget https://github.com/mvisonneau/mmds/releases/download/${MMDS_VERSION}/mmds_${MMDS_VERSION}_linux_386.deb
~$ dpkg -i mmds_${MMDS_VERSION}_linux_386.deb

# RPM package (eg: linux/arm64)
~$ wget https://github.com/mvisonneau/mmds/releases/download/${MMDS_VERSION}/mmds_${MMDS_VERSION}_linux_arm64.rpm
~$ rpm -ivh mmds_${MMDS_VERSION}_linux_arm64.rpm
```

## Develop / Test

```bash
~$ make
all                            Test, builds and ship package for all supported platforms
build                          Build the binaries using local GOOS
clean                          Remove binary if it exists
coverage-html                  Generates coverage report and displays it in the browser
coverage                       Generates coverage report
fmt                            Format source code
gofumpt                        Test code syntax with gofumpt
gosec                          Test code for security vulnerabilities
help                           Displays this help
ineffassign                    Test code syntax for ineffassign
install                        Build and install locally the binary (dev purpose)
is-git-dirty                   Tests if git is in a dirty state
lint                           Run all lint related tests against the codebase
misspell                       Test code with misspell
prerelease                     Build & prerelease the binaries (edge)
release                        Build & release the binaries (stable)
revive                         Test code syntax with revive
setup                          Install required libraries/tools for build tasks
test                           Run the tests against the codebase
vet                            Test code syntax with go vet
```

## Contribute

Contributions are more than welcome! Feel free to submit a [PR](https://github.com/mvisonneau/mmds/pulls).
