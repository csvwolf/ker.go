package main

import (
	"fmt"
	"github.com/csvwolf/ker.go/api"
	"os"
)

func main() {
	ker := &api.Ker{
		SecretKey: os.Getenv("KER_SECRET_KEY"),
		Email:     os.Getenv("KER_EMAIL"),
	}
	res, err := ker.GetDomainList()
	fmt.Printf("%+v, %+v\n", res, err)
}
