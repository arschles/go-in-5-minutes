# testsrv

[![GoDoc](https://godoc.org/github.com/arschles/testsrv?status.svg)](https://godoc.org/github.com/arschles/testsrv)

`testsrv` is a library for running real HTTP servers in the same process as
[Go](http://golang.org) tests, and inspecting the requests that the servers received.

# Sample Usage

```go
myHandler := func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello Gophers!"))
}
srv := testsrv.StartServer(http.HandlerFunc(myHandler))
defer srv.Close()
resp, err := http.Get(srv.URLStr())
//do something with resp and err

// get the last request that the server received
recv := srv.AcceptN(1, 1 * time.Second)
```

# Possible Uses
Since `StartServer` takes in any `http.Handler` it's fairly flexible. Possible applications:

- Testing your own handlers. For example, in situations where [`httptest.ResponseRecorder`](http://godoc.org/net/http/httptest#ResponseRecorder) doesn't meet your needs
- Testing your code that makes its own HTTP requests (for example, an external API call)
