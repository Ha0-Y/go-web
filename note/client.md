# Client

## Request

```go
package main

import (
	`fmt`
	`io`
	`log`
	`net/http`
)

func testGet() {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Error get url: %v", err)
		return
	}
	defer r.Body.Close()

	b, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", string(b))
}

func main() {
	testGet()
}
```

Response 是一个结构体
```go
type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	Header Header
	Body io.ReadCloser
	ContentLength int64
	TransferEncoding []string
	Close bool
	Uncompressed bool
	Trailer Header
	Request *Request
	TLS *tls.ConnectionState
}
```

### URL

一个URL的正常形式应该是
```txt
<schema>://<user>:<password>@<host>:<port>/<path>:<params>?<query>#<frag>
```

因此设计成一个结构体，但是在请求时 Fragment 是会被浏览器去除
```go
// net/url
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port (see Hostname and Port methods)
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```

### Body

请求和响应，都可能存在一个Body
```go
type ReadCloser interface {
	Reader
	Closer
}
```

### Header

```go
// type Header map[string][]string
func testHeader() {
    r, err := http.Get("http://localhost:8080")
    if err != nil {
        log.Fatalf("Error get url: %v", err)
        return
    }
    defer r.Body.Close()
    fmt.Println(r.Header)
    fmt.Println(r.Header.Get("User-Agent"))
}
```