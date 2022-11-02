package main

import (
	"context"
	"fmt"

	"github.com/rentiansheng/mapper"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/2
    @desc:

***************************/

type testCopyStructInline struct {
	A string `validate:"required,max=5"`
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
	errMsg := "Key: 'testCopyStructDst.testCopyStructInline.A' Error:Field validation for 'A' failed on the 'max' tag"
	dst := testCopyStructDst{}
	m := mapper.NewMapper(mapper.OptionValidateStruct().CopyPrivate())
	err := m.Mapper(context.TODO(), src, &dst)

	fmt.Println("error：`"+errMsg+"`     ", errMsg == err.Error())

}
