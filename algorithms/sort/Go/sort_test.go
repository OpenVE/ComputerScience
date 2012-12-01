// Daniel R. (sadasant.com)
// 05/10/2012
//
// License:
//   GNU GPL 3.0
//
// How to run:
//     go test algorithms/primes
//

package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestAll1000(t *testing.T) {
	array := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		array[i] = rand.Int()
	}
	i := Insertion(array)
	b := Bubble(array)
	q := Quick(array)
	m := Merge(array)
	if !reflect.DeepEqual(i, b) {
		t.Errorf("Insertion and Bubble didn't match")
	}
	if !reflect.DeepEqual(i, q) {
		t.Errorf("Insertion and Quick didn't match")
	}
	if !reflect.DeepEqual(i, m) {
		t.Errorf("Insertion and Merge didn't match")
	}
	if !reflect.DeepEqual(b, i) {
		t.Errorf("Bubble and Insertion didn't match")
	}
	if !reflect.DeepEqual(b, q) {
		t.Errorf("Bubble and Quick didn't match")
	}
	if !reflect.DeepEqual(b, m) {
		t.Errorf("Bubble and Merge didn't match")
	}
	if !reflect.DeepEqual(q, m) {
		t.Errorf("Quick and Merge didn't match")
	}
}

func BenchmarkInsertion(b *testing.B) {
	b.StopTimer()
	array := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		array[i] = rand.Int()
	}
	b.StartTimer()
	Insertion(array)
}

func BenchmarkBubble(b *testing.B) {
	b.StopTimer()
	array := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		array[i] = rand.Int()
	}
	b.StartTimer()
	Bubble(array)
}

func BenchmarkQuick(b *testing.B) {
	b.StopTimer()
	array := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		array[i] = rand.Int()
	}
	b.StartTimer()
	Quick(array)
}

func BenchmarkMerge(b *testing.B) {
	b.StopTimer()
	array := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		array[i] = rand.Int()
	}
	b.StartTimer()
	Merge(array)
}
