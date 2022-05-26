package pfsp

import "fmt"

// Evaluater defines different strategy of evaluation.
type Evaluater interface {
	Evaluate(*Instance, []int) (float64, error)
}

// Makespan is a way of getting fitness.
type Makespan struct{}

// Evaluate returns the fitness of the given permutation.
func (*Makespan) Evaluate(instance *Instance, permutation []int) (float64, error) {
	if instance.Jobs != len(permutation) {
		return 0, fmt.Errorf("invalid permutation length, expected %d, got %d", instance.Jobs, len(permutation))
	}
	fitness := float64(0)
	timeTable := []float64{}
	for i := 0; i < instance.Machines; i++ {
		timeTable = append(timeTable, 0)
	}
	for _, individual := range permutation {
		for machine := 0; machine < instance.Machines; machine++ {
			processingTime := instance.Matrix[machine][individual]
			if machine == 0 {
				timeTable[machine] = timeTable[machine] + processingTime
			} else {
				if timeTable[machine-1] < timeTable[machine] {
					timeTable[machine] = timeTable[machine] + processingTime
				} else {
					timeTable[machine] = timeTable[machine-1] + processingTime
				}
			}
		}
	}
	fitness = timeTable[instance.Machines-1]
	return fitness, nil
}

// Totalflowtime is a way of getting fitness.
type Totalflowtime struct{}

// Evaluate returns the fitness of the given permutation.
func (*Totalflowtime) Evaluate(instance *Instance, permutation []int) (float64, error) {
	fitness := float64(0)
	return fitness, nil
}
