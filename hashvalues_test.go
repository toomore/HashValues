package hashvalues

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

var (
	hashedKey []byte
	hashkey   = []byte("Toomore.net")
	hashfunc  = sha256.New
	messageO  = []byte("age=30&name=Toomore")
	messageX  = []byte("name=Toomore&age=30")
)

func init() {
	var h = New(hashkey, hashfunc)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	hashedKey, _ = h.Encode()
}

func ExampleNew() {
	var h = New(hashkey, hashfunc)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	key, msg := h.Encode()
	fmt.Printf("Key:[%s] Msg:[%s]", key, msg)
	// output:
	// Key:[aTMzslluGEzE-uNMoLtBC2vN6aDYGc8fIXJFi_oXPG4=] Msg:[YWdlPTMwJm5hbWU9VG9vbW9yZQ==]
}

func TestHashValues_Encode(t *testing.T) {
	var h = New(hashkey, hashfunc)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	key, msg := h.Encode()
	t.Logf("[%s] [%s]", key, msg)
}

func TestHashValues_Decode(t *testing.T) {
	var h = New(hashkey, hashfunc)
	if err := h.Decode(hashedKey, Base64Encode(messageO)); err == nil {
		t.Logf("O: %+v", h.Values)
	} else {
		t.Log(err)
	}

	h = New(hashkey, hashfunc)
	h.Decode(hashedKey, Base64Encode(messageX))
	t.Logf("X: %+v", h.Values)
}

func BenchmarkHashValues_Encode(b *testing.B) {
	var h = New(hashkey, hashfunc)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	for i := 0; i < b.N; i++ {
		h.Encode()
	}
}

func BenchmarkHashValues_Decode(b *testing.B) {
	var h = New(hashkey, hashfunc)
	var messageObyte = Base64Encode(messageO)
	for i := 0; i < b.N; i++ {
		h.Decode(hashedKey, messageObyte)
	}
}
