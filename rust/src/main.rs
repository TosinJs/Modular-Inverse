// xgcd returns three integers m, x, y which satisfy the equation m = ax + by
// Where m = GCD(a,b)
// x, y are the Bezout's coefficients
// Note: GCD = Greatest Common Divisor
// @Params (a,b) integers obtained from a mod b
fn xgcd(a: isize, b: isize) -> (isize, isize, isize) {

    // Base Case
    // When (a = 0) then m = GCD(0, b) = b. The Largest Number that can divide 0 and b evenly is b
    // When (a = 0, m = b) then b = 1. Obtained by substituting a = 0, m = b into the equation ax + by = m
    if a == 0 {
        return (b, 0, 1);
    };

    // Recursive Case
    // On each iteration the value of a = b % a and b = a
    let (m, x1, y1) = xgcd(b % a, a);
    // Substitute a = b % a and b = a in ax + by = m
    // ax + by = (b % a)x1 + ay1 
    // New values for x and y are obtained from the equation above on each iteration
    let x = y1 - (b/a)*x1;
	let y = x1;
    // m.abs() retuns the absolute value of m. This is used because the GCD is always positive
    return (m.abs(), x, y)
}

// mod_inv returns an integer which is the inverse of a mod b where a_inv x a mod b = 1
// @Params(a, b) where a,b are integers obtained from a mod b
// mod_inv calculates the inverse using the extended Euclidean algorithm defined above (xgcd)
// This function returns 0 if a mod b has no inverse
fn mod_inv(a: isize, b: isize) -> isize {

    // Obtain the GCD(a, b), x, y from the xgcd(a,b)
    let (m, x, _y) = xgcd(a,b);
    // If the GCD(a,b) != 1 This implies that a,b are not coprime and hence there is no inverse for a mod b
    if m != 1 {
        return 0
    };

	// x is obtained from the extended euclidean function above. The Value of y from the function is not required
	// The equation can now be written as ax congruent(=) 1 (mod b)
    // (x mod b)a congruent(=) 1 mod b
	// The inverse of a mod b = x mod b = x % b
    // rem.euclid returns a postive value for x % b. This is for scenarios where the x % b is negative
    return x.rem_euclid(b)
}

fn main() {
    let m_i = mod_inv(7,26);
    println!("{m_i}");
}