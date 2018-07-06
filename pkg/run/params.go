package run

import "sync"

// Type Params stores parameters.
type Params struct {
	Path  *string
	UseWg bool
	Wg    *sync.WaitGroup
}
