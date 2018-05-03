package pfsp

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/chneau/limiter"
)

func TestNew(t *testing.T) {
	files, err := ioutil.ReadDir("../../instances")
	if err != nil {
		t.Errorf("could not read dir instance: %v", err)
	}
	limit := limiter.New(32)
	for i := range files {
		f := files[i]
		limit.Execute(func() {
			jobs := 0
			machines := 0
			instanceNumber := 0
			fmt.Sscanf(f.Name(), "tai%d_%d_%d", &jobs, &machines, &instanceNumber)
			_, err := new(jobs, machines, instanceNumber)
			if err != nil {
				t.Errorf("error with jobs %d machines %d instanceNumber %d: %v", jobs, machines, instanceNumber, err)
			}
		})
	}
	limit.Wait()
}
