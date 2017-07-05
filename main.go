package main

import (
	"fmt"
	"os"

	"github.com/soundcloud/ectoken"
)

func main() {
	secret := os.Args[1]
	expire := os.Args[2]
	path := os.Args[3]
	token, _ := ectoken.Generate(secret, "ec_expire="+expire+"&ec_url_allow="+path)
	fmt.Println(token)
}
