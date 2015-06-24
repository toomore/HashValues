/* Package hashvalues is a simple wrapper for hmac data.

*/
package hashvalues

import (
	"crypto/hmac"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"hash"
	"net/url"
)

// HashValues struct
type HashValues struct {
	Values   url.Values
	hashfunc func() hash.Hash
	hashkey  []byte
}

// New to new a HashValues.
func New(hashkey []byte, hashfunc func() hash.Hash) *HashValues {
	return &HashValues{
		Values:   url.Values{},
		hashfunc: hashfunc,
		hashkey:  hashkey,
	}
}

// Set to set a key-value.
func (h *HashValues) Set(key, value string) {
	h.Values.Set(key, value)
}

// Add to add a key-value.
func (h *HashValues) Add(key, value string) {
	h.Values.Add(key, value)
}

// Del to del a key.
func (h *HashValues) Del(key string) {
	h.Values.Del(key)
}

// Get to get a value of key.
func (h *HashValues) Get(key string) string {
	return h.Values.Get(key)
}

// Decode to decode a hmac key with message.
func (h *HashValues) Decode(key []byte, message string) error {
	var err error

	if key, err = decode(key); err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(h.createMac([]byte(message)), key) == 1 {
		h.Values, err = url.ParseQuery(message)
	} else {
		err = errors.New("wrong key")
	}
	return err
}

// Encode to encode all data.
func (h *HashValues) Encode() ([]byte, []byte) {
	var value = []byte(h.Values.Encode())
	return encode(h.createMac(value)), encode(value)
}

// createMac to create and sum hash.
func (h HashValues) createMac(message []byte) []byte {
	var hashed = hmac.New(h.hashfunc, h.hashkey)
	hashed.Write(message)
	return hashed.Sum(nil)
}

// encode encodes a value using base64.
func encode(value []byte) []byte {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(value)))
	base64.URLEncoding.Encode(encoded, value)
	return encoded
}

// decode decodes a cookie using base64.
func decode(value []byte) ([]byte, error) {
	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(value)))
	b, err := base64.URLEncoding.Decode(decoded, value)
	if err != nil {
		return nil, err
	}
	return decoded[:b], nil
}
