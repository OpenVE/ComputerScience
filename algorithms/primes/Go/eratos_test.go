// Daniel R. (sadasant.com)
// 05/10/2012
//
// License:
//   GNU GPL 3.0
//
// How to run:
//     go test algorithms/primes
//

package primes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEratos(t *testing.T) {
	primes := Eratos(10)
	expected := []int64{2, 3, 5, 7}
	if !reflect.DeepEqual(primes, expected) {
		t.Errorf("Eratos(10) should be: %v, but it was: %v", expected, primes)
	}
}

func BenchmarkEratos(b *testing.B) {
	var i int64
	for ; i < int64(b.N); i++ {
		Eratos(i)
	}
}

func ExampleEratos() {
	primes := Eratos(10)
	fmt.Printf("%v", primes)
	// Output:
	// [2 3 5 7]
}
