package pythagorean

import "fmt"

// Triplet 3 ints, ideally a<b<c, a**a+b**b == c**c
type Triplet [3]int

// Pythagorean triplet is a set of three natural numbers, {a, b, c},
// for which,
// a**2 + b**2 = c**2
// and such that,
// a < b < c

// GOAL Given an input integer N, find all Pythagorean triplets
// for which a + b + c = N.

// Range return a list of triplets for N, where a+b+c = N
func Range(min, max int) []Triplet {

	results := []Triplet{}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				//fmt.Println("a", a, "b", b, "c", c)

				if a*a+b*b == c*c {
					// fmt.Println("valid a", a, "b", b, "c", c, "=", a*a, b*b, c*c)
					results = append(results, Triplet{a, b, c})
				}

			}
		}
	}

	fmt.Println("Range", min, max, "#", len(results))
	return results
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c == p
func Sum(p int) []Triplet {
	allTriplets := Range(1, p)

	results := []Triplet{}

	// filter on a+b+c==p
	for _, trip := range allTriplets {

		a, b, c := trip[0], trip[1], trip[2]
		if a+b+c == p {
			fmt.Println("triplet", trip)
			results = append(results, trip)
		}
	}

	fmt.Println("Sum", 1, p, "#", len(results))
	// for i, res := range results {
	//	fmt.Println(" ", i, "\t", res)
	// }
	return results
}
