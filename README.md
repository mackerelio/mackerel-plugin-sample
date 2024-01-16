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

### Release by GitHub Actions

1. Edit CHANGELOG.md, git commit, git push
2. `git tag vx.y.z`
3. `git push --tags`
4. Wait to build at https://github.com/mackerelio/mackerel-plugin-sample/actions/workflows/goreleaser.yml
5. See https://github.com/mackerelio/mackerel-plugin-sample/releases

### Release by manually

1. Install goreleaser https://goreleaser.com/install/
2. Edit CHANGELOG.md, git commit, git push
3. `git tag vx.y.z`
4. `export GITHUB_TOKEN="<YOUR GITHUB TOKEN>"`
5. `goreleaser release --clean`
6. See https://github.com/mackerelio/mackerel-plugin-sample/releases
