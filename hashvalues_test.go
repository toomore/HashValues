package hashvalues

import (
	"crypto/md5"
	"testing"
)

const messageO = "age=30&name=Toomore"
const messageX = "name=Toomore&age=30"
const hashedKey = "6167653d3330266e616d653d546f6f6d6f7265086c98c4331bffed564a18fbec2b9f96"

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
		t.Logf("%+v", h.Values)
	} else {
		t.Log(err)
	}

	h.Decode([]byte(hashedKey), messageX)
	t.Logf("%+v", h.Values)
}
