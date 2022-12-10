package main

import (
	"errors"
	"fmt"
	"math/big"
)

// egcd computes the variables x,y,m where m = ax + by = m
// Where m = gcd of a and b
// egcd takes in two variables a,b. Where a is the modulus and b is the iterator
func egcd(a, b int) (int, int, int) {
	//Base Case
	//When a = 0, The GCD(0, b) = b. The Largest Number that can divide 0,b evenly is b
	//When (a= 0, m = b) then x = 0, y = 1 (This can be obtained by substituting a,m into the equation ax + by = m)
	if a == 0 {
		return b, 0, 1
	}	
	m, x1, y1 := egcd(b % a, a)
	x := y1 - (b/a)*x1
	y := x1
	// The GCD cannot be negative
	m = abs(m)
	return m, x, y
}

//ModInv Returns the Inverse of a mod b such that a * a^-1 mod b = 1
//The function has two parameters a and b.
func modInv(a, b int) (int, error) {
	m, x, _ := egcd(a, b)
	//If GCD = m is not 1, this implies that a,b are not coprime and hence there is no inverse for a mod b
	if m != 1 {
		return 0, errors.New("this value has no inverse")
	}
	//x is obtained from the extended euclidean function above. The Value of Y from the function is 0
	//The equation can now be written as ax --- 1 (mod b)
	//The inverse of a mod b = x % b
	return mod(x, b), nil
}

//Turns out, the "%" operator in golang does nor return the remainder value I need
// I will the using the golang math package for my mod (%) operations
func mod(a, b int) int {
	z := big.Int{}
	a_p := big.NewInt(int64(a))
	b_p := big.NewInt(int64(b))
	m := big.Int{}
	z.DivMod(a_p, b_p, &m)
	return (int(m.Int64()))
}

// Return the absolute value of the input
func abs (val int) int {
	if val < 1 {
		return val * -1
	}
	return val
}

func main() {
	res, err := modInv(3, 7)
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Println(res)
}
