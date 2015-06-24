package hashvalues

import (
	"crypto/hmac"
	"errors"
	"hash"
	"net/url"
)

type HashValues struct {
	Values   url.Values
	hashfunc func() hash.Hash
	hashkey  []byte
}

func NewHashValues(hashkey []byte, hashfunc func() hash.Hash) *HashValues {
	return &HashValues{
		Values:   url.Values{},
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

func (h *HashValues) Decode(key []byte, message string) error {
	var err error
	if hmac.Equal(h.hashkey, hmac.New(h.hashfunc, key).Sum([]byte(message))) {
		h.Values, err = url.ParseQuery(message)
	} else {
		err = errors.New("wrong key!")
	}
	return err
}

func (h *HashValues) Encode() ([]byte, string) {
	var value = h.Values.Encode()
	return hmac.New(h.hashfunc, h.hashkey).Sum([]byte(value)), value
}
