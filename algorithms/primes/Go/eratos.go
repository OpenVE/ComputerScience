// Daniel R. (sadasant.com)
// 23/09/2012
//
// License:
//   GNU GPL 3.0
//

package primes

// Sieve of Eratosthenes
//
// Runs almost instantly even over 5000000.
//
func Eratos(n int64) []int64 {
	bools := make([]bool, n+1)
	primes := make([]int64, n+1)
	c := 0
	for i := int64(2); i < n; i++ {
		if !bools[i] {
			bools[i] = true
			primes[c] = i
			c++
			for j := i + i; j <= n; j += i {
				bools[j] = true
			}
		}
	}
	return primes[:c]
}
