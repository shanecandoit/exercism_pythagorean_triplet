package pythagorean

import "fmt"

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

	for a := min; a < max; a++ {
		for b := a + 1; b < max; b++ {
			for c := b + 1; c < max; c++ {
				//fmt.Println("a", a, "b", b, "c", c)

				if a*a+b*b == c*c {
					fmt.Println("valid a", a, "b", b, "c", c, "=", a*a, b*b, c*c)
					results = append(results, Triplet{a, b, c})
				}

			}
		}
	}

	fmt.Println("Range", min, max, "#", len(results))
	return results
}

func Sum(p int) []Triplet {
	return nil
}
