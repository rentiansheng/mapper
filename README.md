## Mapper

golang  deep copy library，automatic data mapping。 map to struct, struct to map, struct to struct.
[中文文档](/README-zh-cn.md)

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
	A string `validate:"required,max=5"
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
    //
	m := mapper.NewMapper(OptionValidateStruct().CopyPrivate())

	result := dstStruct{}

	err := m.Mapper(ctx, src, &result)
    fmt.Println(err)
	//
}
```


### Features

- struct private field automatic mapping
- slice automatic mapping
- automatic mapping by field name
- automatic mapping by field tag
- struct to map automatic mapping
- map to struct automatic mapping
- []byte to string automatic mapping
- data type automatic mapping 
-  any data type to interface data type
- []*Type to []Type automatic mapping
- []Type to []*Type  automatic mapping
- copy use struct tag alias name，`json:"aa,copy=bb"`
- validate data by struct tag role [rule detail go-playground/validator](https://github.com/go-playground/validator#baked-in-validations)