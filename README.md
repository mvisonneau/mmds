# mvisonneau/mmds - Missed (AWS) Meta-Data (service)

[![GoDoc](https://godoc.org/github.com/mvisonneau/mmds?status.svg)](https://godoc.org/github.com/mvisonneau/mmds)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/mmds)](https://goreportcard.com/report/github.com/mvisonneau/mmds)
[![Docker Pulls](https://img.shields.io/docker/pulls/mvisonneau/mmds.svg)](https://hub.docker.com/r/mvisonneau/mmds/)
[![Build Status](https://cloud.drone.io/api/badges/mvisonneau/mmds/status.svg)](https://cloud.drone.io/mvisonneau/mmds)
[![Coverage Status](https://coveralls.io/repos/github/mvisonneau/mmds/badge.svg?branch=master)](https://coveralls.io/github/mvisonneau/mmds?branch=master)

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
~$ go get -u github.com/mvisonneau/mmds
```

### Homebrew (linux only)

```bash
~$ brew install mvisonneau/tap/mmds
```

### Docker

```bash
~$ docker run -it --rm mvisonneau/mmds
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
~$ wget https://github.com/mvisonneau/strongbox/releases/download/${MMDS_VERSION}/strongbox_${MMDS_VERSION}_freebsd_amd64.tar.gz
~$ tar zxvf strongbox_${MMDS_VERSION}_freebsd_amd64.tar.gz -C /usr/local/bin

# DEB package (eg: linux/386)
~$ wget https://github.com/mvisonneau/strongbox/releases/download/${MMDS_VERSION}/strongbox_${MMDS_VERSION}_linux_386.deb
~$ dpkg -i strongbox_${MMDS_VERSION}_linux_386.deb

# RPM package (eg: linux/arm64)
~$ wget https://github.com/mvisonneau/strongbox/releases/download/${MMDS_VERSION}/strongbox_${MMDS_VERSION}_linux_arm64.rpm
~$ rpm -ivh strongbox_${MMDS_VERSION}_linux_arm64.rpm
```

## Develop / Test

If you use docker, you can easily get started using :

```bash
~$ make dev-env
# You should then be able to use go commands to work onto the project, eg:
~docker$ make fmt
~docker$ mmds
```

## Contribute

Contributions are more than welcome! Feel free to submit a [PR](https://github.com/mvisonneau/mmds/pulls).
