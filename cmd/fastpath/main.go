package main

import (
	"fmt"

	"github.com/renanbastos93/fastpath"
)

func main() {
	p := fastpath.New("api/:param/abc/*")
	fmt.Println(p.Match("api/entity/abc/wildcard"))
}
