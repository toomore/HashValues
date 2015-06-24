package hashvalues

import (
	"crypto/md5"
	"testing"
)

const messageO = "age=30&name=Toomore"
const messageX = "name=Toomore&age=30"
const hashedKey = "4e0b9ce9751ddfdd32d82fb233144467"

var hashkey = []byte("Toomore.net")

func TestHashValues_Encode(t *testing.T) {
	var h = New(hashkey, md5.New)
	h.Set("name", "Toomore")
	h.Set("age", "30")
	key, msg := h.Encode()
	t.Logf("[%x] [%s]", key, msg)
}

func TestHashValues_Decode(t *testing.T) {
	var h = New(hashkey, md5.New)
	if err := h.Decode([]byte(hashedKey), messageO); err == nil {
		t.Logf("O: %+v", h.Values)
	} else {
		t.Log(err)
	}

	h = New(hashkey, md5.New)
	h.Decode([]byte(hashedKey), messageX)
	t.Logf("X: %+v", h.Values)
}
