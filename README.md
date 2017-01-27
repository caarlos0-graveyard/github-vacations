# github-vacations [![Build Status](https://travis-ci.org/caarlos0/github-vacations.svg?branch=master)](https://travis-ci.org/caarlos0/github-vacations) [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser) [![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)

Automagically ignore all notifications related to work when you are on vacations

Just put the binary somewhere, export a `GITHUB_TOKEN` environment variable,
and put it in your crontab:

```crontab
* * * * * /path/to/github-vacations -t My-Github-Token -o SomeOrg > /dev/null 2>&1
```

Enjoy your vacations! 🏖

## Install

```console
brew install caarlos0/tap/github-vacations
```
