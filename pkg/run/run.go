package run

import (
	"encoding/json"
	"log"
	"os/exec"
)

type Output struct {
	Schema Schema `json:""`
	Job    string `json:""`
}

type Schema struct {
	Results map[string]float64 `json:"results"`
	Labels map[string]string `json:"labels"`
}

func (o *Output) RunJob(p *Params) {
	if p.UseWg {
		defer p.Wg.Done()
	}
	o.RunExec(p.Path)
}

func (o *Output) RunExec(path *string) {

	out, err := exec.Command(*path).Output()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, &o.Schema)
	if err != nil {
		log.Fatal(err)
	}

}
