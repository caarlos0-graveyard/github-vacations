# github-vacations

[![Release](https://img.shields.io/github/release/caarlos0/github-vacations.svg?style=flat-square)](https://github.com/caarlos0/github-vacations/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/caarlos0/github-vacations.svg?style=flat-square)](https://travis-ci.org/caarlos0/github-vacations)
[![Coverage Status](https://img.shields.io/coveralls/caarlos0/github-vacations/master.svg?style=flat-square)](https://coveralls.io/github/caarlos0/github-vacations?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/caarlos0/github-vacations?style=flat-square)](https://goreportcard.com/report/github.com/caarlos0/github-vacations)
[![Godoc](https://godoc.org/github.com/caarlos0/github-vacations?status.svg&style=flat-square)](http://godoc.org/github.com/caarlos0/github-vacations)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)


Automagically ignore all notifications related to work when you are on vacations

Just put the binary somewhere, export a `GITHUB_TOKEN` environment variable,
and put it in your crontab:

```crontab
* * * * * /path/to/github-vacations -t My-Github-Token -o SomeOrg > /dev/null 2>&1
```

Your notifications will be stored on `%HOME/.vacations.db`. You can read them
when you get back by using `github-vacations read`.

![screenshot](screen.png)

Enjoy your vacations! 🏖

## Install

```console
brew install caarlos0/tap/github-vacations
```
