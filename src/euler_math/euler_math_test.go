package euler_math

import "testing"

type NumBool struct {
	number int
	val    bool
}

var prime_pairs []NumBool = []NumBool{
	{3, true},
	{499, true},
	{3613, true},
	{7459, true},
	{2, true},
	{100, false},
	{120, false},
	{500, false},
	{3614, false},
	{7460, false},
	{9, false},
	{63, false},
	{123123123123, false},
	{2971215073, true},
	{19577194573, true},
	{25209506681, true},
	{63018038201, true},
	{13091204281, true},
	{6620830889, true},
	{80630964769, true},
	{6692367337, true},
	{6692337123, false},
}


func TestIsPrime(t *testing.T) {
	for _, num_val := range prime_pairs {
		var is_prime bool = IsPrime(num_val.number)

		if is_prime != num_val.val {
			t.Errorf("Primality error with %d", num_val.number)
		}
	}
}
