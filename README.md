toomore/hashvalues
===================

Simple way to hash key-value data by using url.Values.

Install
--------

    go get -v github.com/toomore/hashvalues

Example - Encode
-----------------

```go
import (
	"crypto/sha256"
	"fmt"
)

var h = hashvalues.New([]byte("Toomore.net"), sha256.New)
h.Set("name", "Toomore")
h.Set("age", "30")
key, msg := h.Encode()
fmt.Printf("Key:[%s] Msg:[%s]", key, msg)
// output:
// Key:[aTMzslluGEzE-uNMoLtBC2vN6aDYGc8fIXJFi_oXPG4=] Msg:[YWdlPTMwJm5hbWU9VG9vbW9yZQ==]
```

Example - Decode
-----------------

```go
var hashkey = []byte("Toomore.net")
var hashedKey = "aTMzslluGEzE-uNMoLtBC2vN6aDYGc8fIXJFi_oXPG4="
var message  = "YWdlPTMwJm5hbWU9VG9vbW9yZQ=="
var h = hashvalues.New(hashkey, sha256.New)
if err := h.Decode(hashedKey, message); err == nil {
    t.Logf("O: %+v", h.Values)
} else {
    t.Log(err)
}
```
