package pfsp

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	// statik
	_ "github.com/chneau/go-taillard/pkg/statik"
	"github.com/rakyll/statik/fs"
)

func new(jobs, machines, instanceNumber int) (*Instance, error) {
	fsys, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("could not instanciate new fs: %v", err)
	}
	fileName := fmt.Sprintf("/tai%d_%d_%d.fsp", jobs, machines, instanceNumber)
	f, err := fsys.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open instance %s: %v", fileName, err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read instance %s: %v", fileName, err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(bufio.ScanWords)

	instance := &Instance{}
	data := []float64{}
	for found := 0; scanner.Scan(); {
		text := scanner.Text()
		number, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			continue
		}
		n := int(number)
		switch found {
		case 0:
			instance.Jobs = n
		case 1:
			instance.Machines = n
		case 2:
			instance.Seed = n
		case 3:
			instance.UpperBound = n
		case 4:
			instance.LowerBound = n
		default:
			data = append(data, float64(n))
		}
		found++
	}
	for i := 0; i < instance.Machines; i++ {
		instance.Matrix = append(instance.Matrix, data[i*instance.Jobs:(i+1)*instance.Jobs])
	}
	return instance, nil
}

// NewTotalflowtime returns a pointer to an instance.
func NewTotalflowtime(jobs, machines, instanceNumber int) (*Instance, error) {
	instance, err := new(jobs, machines, instanceNumber)
	instance.Evaluater = &Totalflowtime{}
	return instance, err
}

// NewMakespan returns a pointer to an instance.
func NewMakespan(jobs, machines, instanceNumber int) (*Instance, error) {
	instance, err := new(jobs, machines, instanceNumber)
	instance.Evaluater = &Makespan{}
	return instance, err
}
