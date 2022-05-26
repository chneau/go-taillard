package pfsp

// Instance is an instance of the permutation flowshop scheduling problem.
type Instance struct {
	Evaluater  Evaluater
	Jobs       int
	Machines   int
	Seed       int
	UpperBound float64
	LowerBound float64
	Instance   int
	Matrix     [][]float64
}

// Evaluate returns the fitness of the given permutation.
func (i *Instance) Evaluate(permuration []int) (float64, error) {
	return i.Evaluater.Evaluate(i, permuration)
}
