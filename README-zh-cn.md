## Mapper
golang 数据深拷贝的类库，支持数据自动映射。 map to struct, struct to map, struct to struct.


### Install 
```go
go get -u github.com/rentiansheng/mapper
```

### Getting Started

```go
package main

import (
	"context"
	"fmt"

	"github.com/rentiansheng/mapper"
)

type testCopyStructInline struct {
	A string
	b string
}

type testCopyStructSrc struct {
	Int           int
	AliasCopy     string `json:"alias_copy"`
	privateString string
	Strings       []string
	testCopyStructInline
}

type testCopyStructDst struct {
	Int           int
	Copy          string `json:"alias_copy"`
	privateString string
	Strings       []string
	testCopyStructInline
}

func main() {
	src := testCopyStructSrc{
		Int:           100,
		AliasCopy:     "alias copy",
		privateString: "private",
		Strings:       []string{"item", "item"},
		testCopyStructInline: testCopyStructInline{
			A: "inline a",
			b: "inline b private",
		},
	}
	dst := testCopyStructDst{}
	err := mapper.AllMapper(context.TODO(), src, &dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("struct.Int field                    ", src.Int == dst.Int)
	fmt.Println("struct.AliasCopy field              ", src.AliasCopy == dst.Copy)
	fmt.Println("struct.privateString field          ", src.privateString == dst.privateString)
	fmt.Println("struct.Strings field                ", len(src.Strings) == len(dst.Strings))
	fmt.Println("struct.testCopyStructInline.A field ", len(src.A) == len(dst.A))
	fmt.Println("struct.testCopyStructInline.B field ", len(src.b) == len(dst.b))

	// test deep copy
	src.Strings[0] = "change item"
	fmt.Println("struct.Strings deep copy test       ", src.Strings[0] != dst.Strings[0])

}

```


### Features

- 支持struct私有字段自动映射
- 支持slice 自动映射
- 支持按照字段名自动映射
- 支持按照tag 自动映射
- 支持struct 到map 自动映射
- 支持map 到 struct 自动映射
- 支持[]byte to string 
- 数据类型自动识别
- 支持 数据 to interface 自动映射
- 实现[]*Type to []Type
- 实现[]Type to []*Type 
- 支持struct tag 别名拷贝，`json:"aa,copy=bb"`

