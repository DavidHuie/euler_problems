package euler_math

import "testing"

type NumBool struct {
	number int
	val    bool
}

var prime_pairs []NumBool = []NumBool{
	NumBool{3, true},
	NumBool{499, true},
	NumBool{3613, true},
	NumBool{7459, true},
	NumBool{2, true},
	NumBool{100, false},
	NumBool{120, false},
	NumBool{500, false},
	NumBool{3614, false},
	NumBool{7460, false},
}


func TestIsPrime(t *testing.T) {
	for _, num_val := range prime_pairs {
		var is_prime bool = IsPrime(num_val.number)

		if is_prime != num_val.val {
			t.Errorf("Primality error with %d", num_val.number)
		}
	}
}
