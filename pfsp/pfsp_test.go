package pfsp

import (
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	files, err := ioutil.ReadDir("instances")
	if err != nil {
		t.Errorf("could not read dir instance: %v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	for i := range files {
		file := files[i]
		go func() {
			defer wg.Done()
			jobs := 0
			machines := 0
			instanceNumber := 0
			fmt.Sscanf(file.Name(), "tai%d_%d_%d", &jobs, &machines, &instanceNumber)
			_, err := new(jobs, machines, instanceNumber)
			if err != nil {
				t.Errorf("error with jobs %d machines %d instanceNumber %d: %v", jobs, machines, instanceNumber, err)
			}
		}()
	}
	wg.Wait()
}
