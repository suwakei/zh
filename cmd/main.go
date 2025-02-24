package main

import (
	"fmt"

	"github.com/suwakei/go-zhconv/zhconv"
)

func main() {
    fmt.Println(zhconv.Z2h("Ｈｅｌｌｏ， ｗｏｒｌｄ！")) // "Hello, world!"
    fmt.Println(zhconv.H2z("Hello, world!")) // "Ｈｅｌｌｏ， ｗｏｒｌｄ！"
}