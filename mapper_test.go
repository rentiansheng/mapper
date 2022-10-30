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

func TestStructToPtrStructMapper(t *testing.T) {
	type inlineStruct struct {
		A string
		b string
	}

	// tst copy []struct to []*struct
	type dstStruct struct {
		inlineStruct
	}
	src := []dstStruct{
		dstStruct{
			inlineStruct{
				A: "inline a",
				b: "inline b private field",
			},
		},
	}

	result := make([]*dstStruct, 0)

	err := AllMapper(ctx, src, &result)
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, src[0].inlineStruct, result[0].inlineStruct)

}

func TestPtrStructToStructMapper(t *testing.T) {
	type inlineStruct struct {
		A string
		b string
	}

	// tst copy []struct to []*struct
	type dstStruct struct {
		inlineStruct
	}
	src := []*dstStruct{
		&dstStruct{
			inlineStruct{
				A: "inline a",
				b: "inline b private field",
			},
		},
	}

	result := make([]dstStruct, 0)

	err := AllMapper(ctx, src, &result)
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, src[0].inlineStruct, result[0].inlineStruct)

}
