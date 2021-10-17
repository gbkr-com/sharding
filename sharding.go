package sharding

import (
	"hash/maphash"
)

// Scheme is a sharding scheme for a specific number of shards, N. A scheme
// returns the index, [0, N), to a shard given a byte slice (see WithBytes) or a
// raw hash value (see WithHash). The implementation uses bit masking, so the
// number of shards must be a power of two.
//
type Scheme struct {
	hash *maphash.Hash
	mask uint64
}

// New returns a sharding scheme for the given number of shards. That number
// must be a power of two.
//
func New(n uint64) (*Scheme, error) {
	if n < 2 {
		return nil, ErrBadSchemeParameter
	}
	if (n & (n - 1)) != 0 {
		return nil, ErrBadSchemeParameter
	}
	return &Scheme{hash: new(maphash.Hash), mask: n - 1}, nil
}

// WithHash returns the shard index for the given hash key.
//
func (s *Scheme) WithHash(key uint64) int {
	if s.hash == nil {
		panic(BadScheme)
	}
	return int(key & s.mask)
}

// WithBytes returns the shard index for the given bytes.
//
func (s *Scheme) WithBytes(key []byte) (i int) {
	if s.hash == nil {
		panic(BadScheme)
	}
	s.hash.Write(key)
	i = s.WithHash(s.hash.Sum64())
	s.hash.Reset()
	return
}
