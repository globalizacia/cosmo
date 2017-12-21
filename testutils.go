package cosmo

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"runtime"
	"testing"
)

const distmodTol = 1e-8 // mag
const distTol = 1e-6    // Mpc
const ageTol = 1e-6     // Gyr

// runTests run cos_func on an array of 'inputs' and compares to 'expected'
func runTests(testFunc func(float64) float64, inputs, expected []float64, tol float64, t *testing.T) {
	// We ask runTest to look one additional stack level down
	// to get the original caller of runTest
	stackLevel := 1
	for i, z := range inputs {
		runTest(testFunc, z, expected[i], tol, t, stackLevel)
	}
}

// runTest runs 'testFunc' on scalar 'input' and compares to 'exp'
func runTest(testFunc func(float64) float64, input float64, exp, tol float64, t *testing.T, stackLevel int) {
	var test_description, test_line string

	pc, file, no, ok := runtime.Caller(stackLevel + 1)
	if ok {
		details := runtime.FuncForPC(pc)
		test_description = details.Name()
		test_line = fmt.Sprintf("%s#%d", file, no)
	} else {
		test_description = "CAN'T DETERMINE TEST NAME"
		test_line = "CAN'T DETERMINE TEST LINE"
	}

	obs := testFunc(input)
	if !floats.EqualWithinAbs(obs, exp, tol) {
		t.Errorf("Failed %s at\n %s\n"+"  Expected %f, return %f",
			test_description, test_line, exp, obs)
	}
}

// makeZeroPairs returns pairs of [0, z] for z in 'inputs'
func makeZeroPairs(inputs []float64) [][2]float64 {
	var newInputs [][2]float64
	var pair [2]float64

	// We ask runTest to look one additional stack level down
	// to get the original caller of runTest
	for _, z := range inputs {
		pair = [2]float64{0, z}
		newInputs = append(newInputs, pair)
	}
	return newInputs
}

// runTestsZ0Z2 run cos_func on a set of inputs and compares to 'expected'
// Creates pairs of (0, z) for z in 'inputs' and passes those to runTestZ1Z2
func runTestsZ0Z2(testFunc func(float64, float64) float64, inputs []float64, expected []float64, tol float64, t *testing.T) {
	newInputs := makeZeroPairs(inputs)

	stackLevel := 1
	for i, zs := range newInputs {
		runTestZ1Z2(testFunc, zs, expected[i], tol, t, stackLevel)
	}
}

// runTest runs 'testFunc' on scalar 'input' and compares to 'exp'
func runTestZ1Z2(testFunc func(float64, float64) float64, input [2]float64, exp, tol float64, t *testing.T, stackLevel int) {
	var test_description, test_line string

	pc, file, no, ok := runtime.Caller(stackLevel + 1)
	if ok {
		details := runtime.FuncForPC(pc)
		test_description = details.Name()
		test_line = fmt.Sprintf("%s#%d", file, no)
	} else {
		test_description = "CAN'T DETERMINE TEST NAME"
		test_line = "CAN'T DETERMINE TEST LINE"
	}

	z1, z2 := input[0], input[1]
	obs := testFunc(z1, z2)
	if !floats.EqualWithinAbs(obs, exp, tol) {
		t.Errorf("Failed %s at\n %s\n"+"  Expected %f, return %f",
			test_description, test_line, exp, obs)
	}
}
