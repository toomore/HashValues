toomore/hashvalues
===================

[![GoDoc](https://godoc.org/github.com/toomore/hashvalues?status.svg)](https://godoc.org/github.com/toomore/hashvalues)

Simple way to hash key-value data by using [url.Values](https://golang.org/pkg/net/url/#Values).

Install
--------

    go get -v github.com/toomore/hashvalues

Example - Encode
-----------------

```go
import (
    "crypto/sha256"
    "fmt"

    "github.com/toomore/hashvalues"
)

var h = hashvalues.New([]byte("Toomore.net"), sha256.New)
h.Set("name", "Toomore")
h.Set("age", "30")
key, msg := h.Encode()
fmt.Printf("Key:[%s] Msg:[%s]", key, msg)
```

output:

    Key:[aTMzslluGEzE-uNMoLtBC2vN6aDYGc8fIXJFi_oXPG4=] Msg:[YWdlPTMwJm5hbWU9VG9vbW9yZQ==]

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

LICENSE
--------

The MIT License (MIT)

Copyright © 2015 Toomore Chiang, http://toomore.net/ <toomore0929@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
