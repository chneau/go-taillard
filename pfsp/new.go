package pfsp

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"strconv"
)

//go:embed instances
var instances embed.FS

func new(jobs, machines, instanceNumber int) (*Instance, error) {
	fileName := fmt.Sprintf("instances/tai%d_%d_%d.fsp", jobs, machines, instanceNumber)
	f, err := instances.Open(fileName)
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
		switch found {
		case 0:
			instance.Jobs = int(number)
		case 1:
			instance.Machines = int(number)
		case 2:
			instance.Seed = int(number)
		case 3:
			instance.UpperBound = float64(number)
		case 4:
			instance.LowerBound = float64(number)
		default:
			data = append(data, float64(number))
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
