# ker.go
Hostker Api SDK 2.0

API List From: <https://doc.hostker.dev/>

Document: <https://pkg.go.dev/github.com/csvwolf/ker.go/api>

Home page: Doing?

## Install
```go
go get github.com/csvwolf/ker
```

##  Usage

```go
package main

import (
	kerApi "github.com/csvwolf/ker/api"
)

func main() {
	ker := kerApi.New(os.Getenv("KER_EMAIL"), os.Getenv("KER_SECRET_KEY"))
	res, err := ker.GetDomainList()
}
```
