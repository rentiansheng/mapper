package mapper

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestValidateStructMapper(t *testing.T) {
	type inlineStruct struct {
		A string `validate:"required,max=5"`
		b string `validate:"required"`
	}

	// tst copy []struct to []*struct
	type dstStruct struct {
		inlineStruct
	}
	src :=
		&dstStruct{
			inlineStruct{
				A: "inline a",
				b: "inline b private field",
			},
		}
	m := NewMapper(OptionValidateStruct().CopyPrivate())

	result := dstStruct{}

	err := m.Mapper(ctx, src, &result)
	require.Equal(t, "Key: 'dstStruct.inlineStruct.A' Error:Field validation for 'A' failed on the 'max' tag", err.Error())
}

func TestTagStructMapper(t *testing.T) {
	type Copy struct {
		F string `json:"json_f"  gorm:"column:gorm_f" copy:"copy_f"`
	}
	type CopyJSON struct {
		F string `json:"json_f,copy=json_copy_f"  gorm:"column:gorm_f"`
	}
	type JSON struct {
		F string `json:"json_f"  gorm:"column:gorm_f"`
	}
	type GORM struct {
		F string `gorm:"column:gorm_f"`
	}
	type RawField struct {
		F string
	}
	src := map[string]string{
		"copy_f":      "copy_f",
		"json_copy_f": "json_copy_f",
		"json_f":      "json_f",
		"gorm_f":      "gorm_f",
		"F":           "F",
	}

	resultCopy := Copy{}
	err := Mapper(ctx, src, &resultCopy)
	require.NoError(t, err)
	require.Equal(t, "copy_f", resultCopy.F)

	resultJSONCopy := CopyJSON{}
	err = Mapper(ctx, src, &resultJSONCopy)
	require.NoError(t, err)
	require.Equal(t, "json_copy_f", resultJSONCopy.F)

	resultJSON := JSON{}
	err = Mapper(ctx, src, &resultJSON)
	require.NoError(t, err)
	require.Equal(t, "json_f", resultJSON.F)

	resultGORM := GORM{}
	err = Mapper(ctx, src, &resultGORM)
	require.NoError(t, err)
	require.Equal(t, "gorm_f", resultGORM.F)

	resultRawField := RawField{}
	err = Mapper(ctx, src, &resultRawField)
	require.NoError(t, err)
	require.Equal(t, "F", resultRawField.F)

}

func TestExtraObjectField(t *testing.T) {
	src := []map[string]interface{}{
		{"field": 1, "field2": "field2"},
		{"field": "str", "field2": "field2"},
	}
	dstField := make([]interface{}, 0)
	err := ObjectsField(ctx, "field", src, &dstField)
	require.NoError(t, err, "TestExtraObjectField field")
	require.Equal(t, []interface{}{1, "str"}, dstField, "TestExtraObjectField field")

	dstField2 := make([]interface{}, 0)
	err = ObjectsField(ctx, "field2", src, &dstField2)
	require.NoError(t, err, "TestExtraObjectField field")
	require.Equal(t, []interface{}{"field2", "field2"}, dstField2, "TestExtraObjectField field2")

}

func TestExtraObjectKeyValue(t *testing.T) {
	src := map[int]int{1: 1, 2: 2}
	keyDst, valueDst := make([]int, 0), make([]int, 0)
	err := MapKeys(ctx, src, &keyDst)
	require.NoError(t, err, "TestExtraObjectKeyValue field")
	sort.Ints(keyDst)
	require.Equal(t, []int{1, 2}, keyDst, "TestExtraObjectKeyValue key")

	err = MapValues(ctx, src, &valueDst)
	require.NoError(t, err, "TestExtraObjectKeyValue field")
	sort.Ints(valueDst)
	require.Equal(t, []int{1, 2}, valueDst, "TestExtraObjectKeyValue value")

}

func TestChunk(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for size := 1; size < 20; size++ {
		dst := make([][]int64, 0)
		err := Chunk(ctx, src, &dst, size)
		require.NoError(t, err, "TestChunk field")
		l := (len(src) + size - 1) / size
		require.Equal(t, l, len(dst), "size %v", size)
		lastItemLen := len(src) % size
		for idx := 0; idx < l; idx++ {
			if idx == l-1 && lastItemLen > 0 {
				require.Equal(t, lastItemLen, len(dst[idx]))
			} else {
				require.Equal(t, size, len(dst[idx]))

			}
		}
	}

	// test not can set
	dst := make([][]int64, 0)
	err := Chunk(ctx, src, dst, 1)
	require.Error(t, err, "TestChunk not can set field")

	// test not can set
	err = Chunk(ctx, src, &dst, 0)
	require.Error(t, err, "TestChunk size = 0field")
}
