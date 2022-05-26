package pfsp

// Evaluater defines different strategy of evaluation.
type Evaluater interface {
	Evaluate(*Instance, []int) (float64, error)
}

// Makespan is a way of getting fitness.
type Makespan struct{}

// Evaluate returns the fitness of the given permutation.
func (*Makespan) Evaluate(instc *Instance, pp []int) (float64, error) {
	fitness := float64(0)
	timeTable := []float64{}
	for i := 0; i < instc.Machines; i++ {
		timeTable = append(timeTable, 0)
	}
	for _, p := range pp {
		for m := 0; m < instc.Machines; m++ {
			processingTime := instc.Matrix[m][p]
			if m == 0 {
				timeTable[m] = timeTable[m] + processingTime
			} else {
				if timeTable[m-1] < timeTable[m] {
					timeTable[m] = timeTable[m] + processingTime
				} else {
					timeTable[m] = timeTable[m-1] + processingTime
				}
			}
		}
	}
	fitness = timeTable[instc.Machines-1]
	return fitness, nil
}

// Totalflowtime is a way of getting fitness.
type Totalflowtime struct{}

// Evaluate returns the fitness of the given permutation.
func (*Totalflowtime) Evaluate(instc *Instance, pp []int) (float64, error) {
	fitness := float64(0)
	timeTable := []float64{}
	for i := 0; i < instc.Machines; i++ {
		timeTable = append(timeTable, 0)
	}
	for _, p := range pp {
		_ = p
	}
	return fitness, nil
}
