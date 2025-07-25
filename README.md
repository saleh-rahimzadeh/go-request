# go-request

The **go-request** is a web request provider with retry for Golang projects.

---

Install:

```sh
go get -u github.com/saleh-rahimzadeh/go-request
```

Import:

```go
import (
  "github.com/saleh-rahimzadeh/go-request"
)
```

---

## Usage

Demand:

```go
d := request.BuildDemand(http.MethodPost, "https://<URL>", "/<PATH>")

d = d.ContentType(request.HTTP_JSON)
d = d.Header("<KEY>", "<VALUE>")
d = d.Authorization("<VALUE>")
d = d.AuthorizationBearer("<VALUE>")
d = d.Parameter(map[string]string{
  "<KEY>": "<VALUE>",
})

if d.Error != nil {
  if errors.Is(d.Error, request.ErrDemandContentTypeEmpty) {
  }
}
```

Request:

```go
r := request.New(
  time.Minute,     /* Connection timeout */
  []time.Duration{ /* Retries */
    time.Second * 1,
    time.Second * 2,
    time.Second * 3,
  },
)
```

Send:

```go
result, properties, success := r.Send(d)
result, properties, success := r.SendJson(d, map[string]string{ "<KEY>": "<VALUE>" })
result, properties, success := r.SendForm(d, map[string]string{ "<KEY>": "<VALUE>" })

if !success {
  log.Fatalf("failed")
}

log.Println(result.IsOK)       // is status ok, indeed does response got http.StatusOK
log.Println(result.StatusCode) // http status code
log.Println(result.Body)       // represents the response body
log.Println(result.BodyObject) // represents the response body marshaled as map[string]any

log.Println(properties.Elapsed)      // time spend to getting last response
log.Println(properties.TotalElapsed) // total time spend to getting responses
log.Println(properties.Retries)      // number of retries performed
log.Println(properties.Errors)       // an array of all errors that occurred during the request
```

---

### Log web server

Run log web server:

```sh
go run testdata/webserver/main.go
```
