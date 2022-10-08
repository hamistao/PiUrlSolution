package pi

import (
	"math"
	"math/big"
)

func mul_mod(a, b, m int64) int64 {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	bigM := big.NewInt(m)
	res := big.NewInt(0)
	res.Mul(bigA, bigB)
	res.Mod(res, bigM)
	return res.Int64()
}

/* return the inverse of x mod y */
func inv_mod(x, y int64) int64 {
	var q, t int64

	u := x
	v := y
	c := int64(1)
	a := int64(0)

	for ok := true; ok; ok = (u != 0) {
		q = v / u

		t = c
		c = a - q*c
		a = t

		t = u
		u = v - q*u
		v = t
	}

	a = a % y
	if a < 0 {
		a = y + a
	}

	return int64(a)
}

/* return (a^b) mod m */
func pow_mod(a, b, m int64) int64 {

	r := int64(1)
	aa := a
	for true {
		if b%2 != 0 {
			r = mul_mod(r, aa, m)
		}
		b = b >> 1
		if b == 0 {
			break
		}
		aa = mul_mod(aa, aa, m)
	}
	return r
}

/* return true if n is prime */
func is_prime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

/* return the prime number immediatly after n */
func next_prime(n int64) int64 {
	for ok := true; ok; ok = !is_prime(n) {
		n++
	}
	return n
}

func Plouffe(n int) int {
	var av, num, den, kq, kq2, t, s int64

	N := int64(float64(n+20) * math.Log(10) / math.Log(2))
	var sum float64 = 0

	for a := int64(3); a <= (2 * N); a = next_prime(a) {
		vmax := int(math.Log(float64(2*N)) / math.Log(float64(a)))
		av = 1

		for i := 0; i < vmax; i++ {
			av = av * a
		}

		s = 0
		num = 1
		den = 1
		v := 0
		kq = 1
		kq2 = 1

		for k := int64(1); k <= N; k++ {
			t = k
			if kq >= a {
				for ok := true; ok; ok = ((t % a) == 0) {
					t = t / a
					v--
				}
				kq = 0
			}
			kq++
			num = mul_mod(num, t, av)

			t = (2*k - 1)
			if kq2 >= a {
				if kq2 == a {
					for ok := true; ok; ok = ((t % a) == 0) {
						t = t / a
						v++
					}
				}
				kq2 -= a
			}
			den = mul_mod(den, t, av)
			kq2 += 2

			if v > 0 {
				t = inv_mod(den, av)
				t = mul_mod(t, num, av)
				t = mul_mod(t, k, av)
				for i := v; i < vmax; i++ {
					t = mul_mod(t, a, av)
				}
				s += t
				if s >= av {
					s -= av
				}
			}

		}

		t = pow_mod(10, int64(n-1), av)
		s = mul_mod(s, t, av)
		sum = math.Mod(sum+float64(s)/float64(av), 1.0)
	}
	return int(sum * 1e9)
}
