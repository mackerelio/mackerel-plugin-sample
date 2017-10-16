# mackerel-plugin-sample

Sample plugin for mackerel.io agent.  This repository releases an artifact to Github Releases, which satisfy the format for mkr plugin installer.

## Synopsis

```shell
mackerel-plugin-sample [-metric-key-prefix=<prefix>]
```

## Example of mackerel-agent.conf

```
[plugin.metrics.sample]
command = "/path/to/mackerel-plugin-sample"
```

## How to release

[goxc](https://github.com/laher/goxc) and [ghr](https://github.com/tcnksm/ghr) are used to release.

### Release by TravisCI

1. Edit CHANGELOG.md, git commit, git push
2. `git tag vx.y.z`
3. git push --tags
4. Wait to build at https://travis-ci.org/mackerelio/mackerel-plugin-sample
5. See https://github.com/mackerelio/mackerel-plugin-sample/releases

Don't forget setting GITHUB_TOKEN as environment variables in TravisCI.  If you don't know how, see https://docs.travis-ci.com/user/environment-variables/#Defining-Variables-in-Repository-Settings .

### Release by manually

1. Install goxc and ghr by `make setup`
2. Edit CHANGELOG.md, git commit, git push
3. `git tag vx.y.z`
4. GITHUB_TOKEN=... script/release
5. See https://github.com/mackerelio/mackerel-plugin-sample/releases
