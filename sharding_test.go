package sharding

import (
	"testing"
)

func TestAllowedSchemes(t *testing.T) {
	cases := []struct {
		shards uint64
		ok     bool
	}{
		{0, false},
		{2, true},
		{4, true},
		{6, false},
	}
	for i, c := range cases {
		_, err := New(c.shards)
		if err == nil {
			if !c.ok {
				t.Error(i)
			}
		} else {
			if c.ok {
				t.Error(i)
			}
		}
	}
}

func TestHash(t *testing.T) {
	//
	// Use four shards, which ought to map to the lower two bits of each
	// hash value.
	//
	s, _ := New(4)
	//
	// Expected shards for a 'hash' equal to the index of the slice.
	//
	expected := []int{0, 1, 2, 3, 0, 1}
	for i, e := range expected {
		h := s.WithHash(uint64(i))
		if h != e {
			t.Errorf("expected %d got %d", e, h)
		}
	}
}

func TestBytes(t *testing.T) {
	scheme, _ := New(4)
	cases := []struct {
		key   string
		index int
	}{
		{"abc", -1},
		{"def", -1},
		{"ghi", -1},
		{"jkl", -1},
		{"mno", -1},
		{"pqr", -1},
		{"stu", -1},
		{"vwx", -1},
		{"yz ", -1},
	}
	for i, c := range cases {
		cases[i].index = scheme.WithBytes([]byte(c.key))
	}
	//
	// Now repeat and assert the same index is calculated.
	//
	for _, c := range cases {
		i := scheme.WithBytes([]byte(c.key))
		if i != c.index {
			t.Errorf("want %d got %d\n", c.index, i)
		}
	}

}
