# mvisonneau/mmds - Missed (AWS) Meta-Data (service)

[![GoDoc](https://godoc.org/github.com/mvisonneau/mmds?status.svg)](https://godoc.org/github.com/mvisonneau/mmds)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/mmds)](https://goreportcard.com/report/github.com/mvisonneau/mmds)
[![Docker Pulls](https://img.shields.io/docker/pulls/mvisonneau/mmds.svg)](https://hub.docker.com/r/mvisonneau/mmds/)
[![Build Status](https://travis-ci.org/mvisonneau/mmds.svg?branch=master)](https://travis-ci.org/mvisonneau/mmds)
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

VERSION:
   <devel>

COMMANDS:
     pricing-model  get instance pricing-model
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level level    log level (debug,info,warn,fatal,panic) (default: "info") [$MMDS_LOG_LEVEL]
   --log-format format  log format (json,text) (default: "text") [$MMDS_LOG_FORMAT]
   --help, -h           show help
   --version, -v        print the version
```

## Install

You can have a look at the [release page](https://github.com/mvisonneau/mmds/releases) of the project, we currently build it for **Linux**, **Darwin** and **Windows** platforms.

```
~$ wget https://github.com/mvisonneau/mmds/releases/download/0.1.0/mmds_linux_amd64 -O /usr/local/bin/mmds; chmod +x /usr/local/bin/mmds
```

You can also use the [docker version](https://hub.docker.com/r/mvisonneau/mmds):

```
~$ docker run -it --rm mvisonneau/mmds
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
