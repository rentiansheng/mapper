package mapper

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/24
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

func TestMapper(t *testing.T) {
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
	err := Mapper(context.TODO(), src, &dst)
	require.NoError(t, err)
	require.Equal(t, src.Int, dst.Int)
	require.Equal(t, src.AliasCopy, dst.Copy)
	require.Equal(t, "", dst.privateString)
	require.Equal(t, src.A, dst.A)
	require.Equal(t, "", dst.b)

}

func TestMapperAll(t *testing.T) {
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
	err := AllMapper(context.TODO(), src, &dst)
	require.NoError(t, err)
	require.Equal(t, src.Int, dst.Int)
	require.Equal(t, src.AliasCopy, dst.Copy)
	require.Equal(t, src.privateString, dst.privateString)
	require.Equal(t, src.A, dst.A)
	require.Equal(t, src.b, dst.b)
}
