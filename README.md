# httpmethodoverride #

`httpmethodoverride` is a Go package that provides an `http.Handle`
wrapper that allows clients to override the provided HTTP method with
the value within a `_method` query parameter.

Some situations in which you might want this:

  * Your client stack places silly restrictions on your use, like
    preventing a body in `DELETE` requests.

  * You are in a browser and get redirected from a `GET` to a `POST`
    but you still want to hit a `POST` API.

If you use this with a router that does method-based routing, as long
as you wrap the toplevel router/muxer in
`httpmethodoverride.Handler()` things will work exactly as if the
actual request were made with the desired HTTP method.

## API ##

```go
mux := http.NewServeMux()
mux.Handle("/", http.HandlerFunc(func w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(r.Method))
})

http.ListenAndServe(":8000", httpmethodoverride.Handler(mux))
```

That's it.  But it's also
[on godoc](http://godoc.org/github.com/joeshaw/httpmethodoverride).

```
$ curl http://localhost:8080/

GET
```

```
$ curl http://localhost:8080/?_method=POST

POST
```
