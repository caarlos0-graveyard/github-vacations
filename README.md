# github-vacations [![Build Status](https://travis-ci.org/caarlos0/github-vacations.svg?branch=master)](https://travis-ci.org/caarlos0/github-vacations) [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

Automagically ignore all notifications related to work when you are on vacations

Just put the binary somewhere, export a `GITHUB_TOKEN` environment variable,
and put it in your crontab:

```crontab
* * * * * GITHUB_TOKEN="xyz" /path/to/github-vacations SomeOrg > /dev/null 2>&1
```

Enjoy your vacations! 🏖

## Install

```console
brew tap caarlos0/formulae
brew install github-vacations
```
