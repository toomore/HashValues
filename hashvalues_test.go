package hashvalues

import (
	"crypto/hmac"
	"crypto/md5"
	"testing"
)

const messageO = "age=30&name=Toomore"
const messageX = "name=Toomore&age=30"

var (
	hashedKey []byte
	hashkey   = []byte("Toomore.net")
	hashed    = hmac.New(md5.New, hashkey)
)

func init() {
	hashed.Write([]byte(messageO))
	hashedKey = hashed.Sum(nil)
}

func TestHashValues_Encode(t *testing.T) {
	var h = New(hashkey, md5.New)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	key, msg := h.Encode()
	t.Logf("[%x] [%s]", key, msg)
}

func TestHashValues_Decode(t *testing.T) {
	var h = New(hashkey, md5.New)
	if err := h.Decode(hashedKey, messageO); err == nil {
		t.Logf("O: %+v", h.Values)
	} else {
		t.Log(err)
	}

	h = New(hashkey, md5.New)
	h.Decode(hashedKey, messageX)
	t.Logf("X: %+v", h.Values)
}

func BenchmarkHashValues_Encode(b *testing.B) {
	var h = New(hashkey, md5.New)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	for i := 0; i < b.N; i++ {
		h.Encode()
	}
}

func BenchmarkHashValues_Decode(b *testing.B) {
	var h = New(hashkey, md5.New)
	for i := 0; i < b.N; i++ {
		h.Decode(hashedKey, messageO)
	}
}
