/*
Read Doc in api dir
*/
package main

import (
	"fmt"
	"github.com/csvwolf/ker.go/api"
	"os"
)

func main() {
	ker := api.New(os.Getenv("KER_EMAIL"), os.Getenv("KER_SECRET_KEY"))
	res, err := ker.GetDomainList()
	fmt.Printf("%+v, %+v\n", res, err)
}
