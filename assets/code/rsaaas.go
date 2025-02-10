package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const e = 65537

func generatePrime(bits int) *big.Int {
	// We want p = e*k + 1 to be prime
	// This ensures e divides p-1
	for {
		// Generate k with appropriate size
		// Create upper bound for k: 2^(bits-17)
		bound := new(big.Int).Lsh(big.NewInt(1), uint(bits-17))
		k, err := rand.Int(rand.Reader, bound)
		if err != nil {
			panic(err)
		}

		// Calculate p = e*k + 1
		p := new(big.Int).Mul(big.NewInt(e), k)
		p.Add(p, big.NewInt(1))

		// Check if p is in valid range (63-64 bits) and prime
		if p.BitLen() == 64 && p.ProbablyPrime(20) {
			return p
		}
	}
}

func main() {
	fmt.Println("Finding suitable primes...")

	// Generate two different primes
	var p, q *big.Int
	for {
		p = generatePrime(64)
		q = generatePrime(64)

		if p.Cmp(q) != 0 {
			break
		}
	}

	fmt.Printf("Found primes:\n")
	fmt.Printf("p = %s\n", p.String())
	fmt.Printf("q = %s\n", q.String())

	// Verify they'll work
	p1 := new(big.Int).Sub(p, big.NewInt(1))
	q1 := new(big.Int).Sub(q, big.NewInt(1))
	phi := new(big.Int).Mul(p1, q1)

	gcdVal := new(big.Int).GCD(nil, nil, big.NewInt(e), phi)
	fmt.Printf("gcd(e, φ(n)) = %s\n", gcdVal.String())
}
