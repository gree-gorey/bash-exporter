package run

import (
	"encoding/json"
	"log"
	"os/exec"
)

type Output struct {
	Result map[string]float64 `json:""`
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

	err = json.Unmarshal(out, &o.Result)
	if err != nil {
		log.Fatal(err)
	}

}
