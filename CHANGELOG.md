# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to  [0ver](https://0ver.org).

## [Unreleased]

### Changed

- ci: migrated to github actions
- go: upgraded to 1.17
- deps: updated all of them to their latest versions

## [0.1.2] - 2020-10-29

### Added

- Release linux arm64 binary
- gosec tests

### Changed

- Follow golang file structure
- Release artifacts using goreleaser
- Bumped all dependencies

## [0.1.1] - 2019-05-27

### Changed

- Release binaries are now automatically built and published from the CI
- Use go modules
- Monitor dependencies with dependabot
- Optimized Makefile
- Moved CI from `Travis` to `Drone`
- Rewrote LICENSE to markdown
- Use busybox as base docker image

## [0.1.0] - 2019-03-04

### Added

- Working state of the app
- get instance pricing-model information from lifecycle values
- got some tests in place
- Makefile
- LICENSE
- README

[Unreleased]: https://github.com/mvisonneau/mmds/compare/0.1.2...HEAD
[0.1.2]: https://github.com/mvisonneau/mmds/tree/0.1.2
[0.1.1]: https://github.com/mvisonneau/mmds/tree/0.1.1
[0.1.0]: https://github.com/mvisonneau/mmds/tree/0.1.0
