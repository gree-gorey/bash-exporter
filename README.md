# bash-exporter

Prometheus exporter for bash scripts

```console
Usage of ./bash-exporter:
  -debug
    	Debug log level
  -interval int
    	Interval for metrics collection in seconds (default 300)
  -path string
    	path to directory with bash scripts (default "/scripts")
  -prefix string
    	Prefix for metrics (default "bash")
  -web.listen-address string
    	Address on which to expose metrics (default ":9300")
```

## TODO
- [x] Docker image
- [x] Several scripts
