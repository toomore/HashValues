package hashvalues

import (
	"hash"
	"net/url"
)

type HashValues struct {
	Values   *url.Values
	hashfunc hash.Hash
	hashkey  []byte
}

func NewHashValues(hashkey []byte, hashfunc hash.Hash) *HashValues {
	return &HashValues{
		Values:   &url.Values{},
		hashfunc: hashfunc,
		hashkey:  hashkey,
	}
}
