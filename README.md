# Prometheus bash exporter

Simple & minimalistic Prometheus exporter for bash scripts.

[![Go Report Card](https://goreportcard.com/badge/github.com/gree-gorey/bash-exporter)](https://goreportcard.com/report/github.com/gree-gorey/bash-exporter)  
[![](https://images.microbadger.com/badges/image/greegorey/bash-exporter.svg)](https://microbadger.com/images/greegorey/bash-exporter "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/greegorey/bash-exporter.svg)](https://microbadger.com/images/greegorey/bash-exporter "Get your own version badge on microbadger.com")

## Installation
Use [Docker image](https://hub.docker.com/r/greegorey/bash-exporter/) or binary file from [releases](https://github.com/gree-gorey/bash-exporter/releases).

## Usage

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

Just point `-path` flag to the directory with your bash scripts. Names of the files (`(.*).sh`) will be used as the `job` label. Bash scripts should return valid json (see [examples](https://github.com/gree-gorey/bash-exporter/tree/master/examples)).

Example output:
```console
# HELP bash bash exporter metrics
# TYPE bash gauge
bash{job="job-1",verb="items"} 21
bash{job="job-2",verb="get"} 0.003
bash{job="job-2",verb="put"} 0.13
bash{job="job-2",verb="time"} 0.5
...
```

## TODO
- [x] Docker image
- [x] Several scripts
