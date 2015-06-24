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

func (h *HashValues) Set(key, value string) {
	h.Values.Set(key, value)
}

func (h *HashValues) Add(key, value string) {
	h.Values.Add(key, value)
}

func (h *HashValues) Del(key string) {
	h.Values.Del(key)
}

func (h *HashValues) Get(key string) string {
	return h.Values.Get(key)
}
