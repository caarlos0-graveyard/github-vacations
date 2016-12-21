# github-vacations
Automagically ignore all notifications related to work when you are on vacations

Just put the binary somewhere, export a `GITHUB_TOKEN` environment variable,
and put it in your crontab:

```crontab
* * * * * GITHUB_TOKEN="xyz" /path/to/github-vacations SomeOrg > /dev/null 2>&1
```

Enjoy your vacations! 🏖
