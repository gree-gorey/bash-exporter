package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"io/ioutil"

	"github.com/gree-gorey/bash-exporter/pkg/run"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	verbMetrics *prometheus.GaugeVec
)

func main() {
	addr := flag.String("web.listen-address", ":9300", "Address on which to expose metrics")
	interval := flag.Int("interval", 300, "Interval for metrics collection in seconds")
	path := flag.String("path", "/scripts", "path to directory with bash scripts")
	prefix := flag.String("prefix", "bash", "Prefix for metrics")
	debug := flag.Bool("debug", false, "Debug log level")
	flag.Parse()

	verbMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s", *prefix),
			Help: "bash exporter metrics",
		},
		[]string{"verb", "job"},
	)
	prometheus.MustRegister(verbMetrics)

	files, err := ioutil.ReadDir(*path)
  if err != nil {
      log.Fatal(err)
  }

	var names []string
  for _, f := range files {
		if (f.Name()[0:1] != ".") {
			names = append(names, f.Name())
		}
  }

	http.Handle("/metrics", prometheus.Handler())
	go Run(int(*interval), *path, names, *debug)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func Run(interval int, path string, names []string, debug bool) {
	for {
		var wg sync.WaitGroup
		oArr := []*run.Output{}
		wg.Add(len(names))
		for _, name := range names {
			o := run.Output{}
			o.Job = strings.Split(name, ".")[0]
			oArr = append(oArr, &o)
			thisPath := path + "/" + name
			p := run.Params{UseWg: true, Wg: &wg, Path: &thisPath}
			go o.RunJob(&p)
		}
		wg.Wait()
		// if debug == true {
		// 	ser, err := json.Marshal(o)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// 	log.Println(string(ser))
		// }
		verbMetrics.Reset()
		for _, o := range oArr {
			for metric, value := range o.Result {
				verbMetrics.With(prometheus.Labels{"verb": metric, "job": o.Job}).Set(float64(value))
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
