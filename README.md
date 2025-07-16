# Woodpecker-ascii-junit

A simple Woodpecker CI plugin that prints out JUnit summaries in ASCII:

```
| Passed | Failed | Errored | Skipped | Total |
_______________________________________________
| 20     | 5      | 0       | 1       | 1     | 

Total time: 1.378s
```

If there are failing tests, details of those will be printed as well.

If you are using Drone-CI, consider using [drone-junit](https://github.com/rohit-gohri/drone-junit/) instead 
that has a nice Adaptive Card UI which is currently not supported by Woodpecker-CI.

## Configuration

See `docker-compose.yml` as an example:

- `PLUGIN_PATH` env var or `path` setting in Woodpecker
- Optional: `PLUGIN_LOG_LEVEL` env var or `log-level` (built-in Woodpecker plugin)

Here's an example how to include it in your Woodpecker workflow:

```
  - name: junit-reports
    image: ghcr.io/wgroeneveld/woodpecker-ascii-junit:main
    settings:
      log-level: debug
      path: /tmp/reports/**/*.xml
    when:
      status: [
        'success',
        'failure',
      ]
```
