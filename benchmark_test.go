package cosmo

import (
	"reflect"
	"testing"
)

func BenchmarkE(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.E(z)
	}
}

func BenchmarkEinv(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.Einv(z)
	}
}

func BenchmarkComovingDistance(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.ComovingTransverseDistance(z)
	}
}

func BenchmarkComovingTransverseDistance(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.ComovingTransverseDistance(z)
	}
}

func BenchmarkLuminosityDistance(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.LuminosityDistance(z)
	}
}

// benchmarkNdistance is a helper function to be called by specific benchmarks
func benchmarkNdistance(n int, distFunc string, b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
//	var funcToTest func(float64) float64
	funcToTest := reflect.ValueOf(&cos).MethodByName(distFunc)
	var z float64
	z_max := 1.0
	step := z_max / float64(n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			z = 0.001 + step*float64(j)
			funcToTest.Call([]reflect.Value{reflect.ValueOf(z)})
		}
	}
}

func BenchmarkNComovingDistance(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	const n = 10000
	var z float64
	z_max := 1.0
	step := z_max / n
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			z = 0.001 + step*float64(j)
			cos.ComovingDistance(z)
		}
	}
}

func BenchmarkNLuminosityDistance(b *testing.B) {
	cos := Cosmology{Om0: 0.27, Ol0: 0.73, Ok0: 0., H0: 70, w0: -1.0, Tcmb0: 0.}
	const n = 10000
	var z float64
	z_max := 1.0
	step := z_max / n
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			z = 0.001 + step*float64(j)
			cos.LuminosityDistance(z)
		}
	}
}