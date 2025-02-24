# go-zhconv


このリポジトリは、Go言語での文字変換をサポートするライブラリです。全角と半角、カナの相互変換を行います。

## インストール

```sh
go get github.com/suwakei/go-zhconv


## Usage
### convert from FullWidth to HalfWidth

package main

import (
    "fmt"
    "github.com/suwakei/go-zhconv"
)

func main() {
    result := zhconv.Z2h("Ｈｅｌｌｏ， ｗｏｒｌｄ！")
    fmt.Println(result) // "Hello, world!"
}

### convert from HalfWidth to FullWidth
package main

import (
    "fmt"
    "github.com/suwakei/go-zhconv"
)

func main() {
    result := zhconv.H2z("Hello, world!")
    fmt.Println(result) // "Ｈｅｌｌｏ， ｗｏｒｌｄ！"
}