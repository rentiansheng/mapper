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
