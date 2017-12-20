# go-zip-db

Go ZIP lookup database optimized for fast application startup (like FaaS) and ease of use. 

## Usage

Trivial to use:

```go

import (
    "github.com/almartin/go-zip-db"
)

```
Goroutine safe lookup function (reader-only). Returns Geolocation structure and any errors.

```go

gl, err := ZipDb.Lookup(zip)

type GeoLocation struct {
	Latitude  float32
	Longitude float32
	Name      string
	State     string
}
```
