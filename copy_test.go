package mapper

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/25
    @desc:

***************************/

var (
	ctx         = context.TODO()
	defaultCopy = newCopyValue()
)

func TestBooleanCopyValue(t *testing.T) {

	var result bool
	srcValue := true
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.BooleanCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")

}

func TestIntCopyValue(t *testing.T) {
	var result int
	srcValue := 10
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.IntCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestInt64CopyValue(t *testing.T) {
	var result int64
	srcValue := 10
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.IntCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, int64(srcValue), result, "result not equal src value")
}

func TestUintCopyValue(t *testing.T) {
	var result uint
	srcValue := uint(10)
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.UintCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestUint64CopyValue(t *testing.T) {
	var result uint64
	srcValue := uint(10)
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.UintCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, uint64(srcValue), result, "result not equal src value")
}

func TestFloatCopyValue(t *testing.T) {
	var result float64
	srcValue := float64(10)
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.FloatCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestStringCopyValue(t *testing.T) {
	var result string
	srcValue := "test string copy value"
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StringCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestStringByteSliceCopyValue(t *testing.T) {
	var result string
	srcValue := []byte("test string copy value")
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StringCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, string(srcValue), result, "result not equal src value")
}

func TestStringByteSliceNilCopyValue(t *testing.T) {
	var result string
	var srcValue []byte = nil
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StringCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, string(srcValue), result, "result not equal src value")
}

func TestStrSliceCopyValue(t *testing.T) {
	var result []string
	srcValue := []string{"str1", "str2", "str3"}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.SliceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestNotZeroStrSliceCopyValue(t *testing.T) {
	result := make([]string, 0)
	srcValue := []string{"str1", "str2", "str3"}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.SliceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestIntSliceCopyValue(t *testing.T) {
	var result []int
	srcValue := []int{1, 2, 3, 4, 5, 6, 7}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.SliceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestNotZeroIntSliceCopyValue(t *testing.T) {
	result := make([]int, 0)
	srcValue := []int{1, 2, 3, 4, 5, 6, 7}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.SliceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestPtrIntSliceCopyValue(t *testing.T) {
	result := []*int{}
	srcValue := []int{1, 2, 3, 4, 5, 6}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	if err := defaultCopy.SliceCopyValue(ctx, src, dst.Elem()); err != nil {
		// TODO: wait implement
		require.Equal(t, "cannot copy slice from []int into []*int", err.Error())
	}
	// TODO: wait implement
	// require.Equal(t, srcValue, *result, "result not equal src value")

}

func TestPtrIntNilSliceCopyValue(t *testing.T) {
	var result []*int = nil
	srcValue := []int{1, 2, 3, 4, 5, 6}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	if err := defaultCopy.SliceCopyValue(ctx, src, dst.Elem()); err != nil {
		// TODO: wait implement
		require.Equal(t, "cannot copy slice from []int into []*int", err.Error())
	}
	// TODO: wait implement
	// require.Equal(t, srcValue, result, "result not equal src value")

}

func TestPtrStrCopyValue(t *testing.T) {
	var rawResult string
	result := &rawResult
	srcValue := "test ptr copy value"
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.PtrCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, *result, "result not equal src value")

}

func TestPtrStrNIlCopyValue(t *testing.T) {
	var result *string = nil
	srcValue := "test ptr copy value"
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.PtrCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, *result, "result not equal src value")

}

func TestPtrIntCopyValue(t *testing.T) {
	var rawResult int
	result := &rawResult
	srcValue := 9999
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.PtrCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, *result, "result not equal src value")

}

func TestPtrIntNolCopyValue(t *testing.T) {
	var result *int = nil
	srcValue := 9999
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.PtrCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, *result, "result not equal src value")

}

func TestIntInterfaceCopyValue(t *testing.T) {
	var result interface{}
	srcValue := 999
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.InterfaceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestStrInterfaceCopyValue(t *testing.T) {
	var result interface{}
	srcValue := "test str interface copy value"
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.InterfaceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestMapInterfaceCopyValue(t *testing.T) {
	var result interface{}
	srcValue := map[string]string{"test": "test str interface copy value"}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.InterfaceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestIntIntMapCopyValue(t *testing.T) {
	var result map[int]int
	srcValue := map[int]int{1: 1, 2: 2}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.MapCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
	srcValue[1] = 11
	require.Equal(t, 1, result[1], "result not equal src value")

}

func TestStrStrMapCopyValue(t *testing.T) {
	var result map[string]string
	srcValue := map[string]string{"1": "1", "2": "2"}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.MapCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestStrInterfaceMapCopyValue(t *testing.T) {
	var result map[string]interface{}
	srcValue := map[string]interface{}{"str": "str", "int": 2, "uint": int64(100), "ints": []int{1, 2, 3}, "strs": []string{"str", "str"}}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.MapCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue, result, "result not equal src value")
}

func TestMapStructMapCopyValue(t *testing.T) {
	defaultCopy := newCopyValue()
	defaultCopy.structCache.copyPrivate = true
	type testS struct {
		str  string
		int  int
		ints []int
	}
	stringTests := func(t testS) string {
		return fmt.Sprintf("str: %s, int: %d, ints: %v", t.str, t.int, t.ints)
	}
	var result map[string]testS
	srcValue := map[string]testS{
		"tt": testS{
			str:  "str",
			int:  1,
			ints: []int{1, 2, 3},
		},
	}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.MapCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, stringTests(srcValue["tt"]), stringTests(result["tt"]), "result not equal src value")

}

func TestStructCopyValue(t *testing.T) {
	type SrcStruct struct {
		Int  int `json:"int_copy"`
		Strs []string

		testPrivate string
	}

	type DstStruct struct {
		IntCopy     int `json:"int_copy"`
		Strs        []string
		testPrivate string
	}

	srcValue := SrcStruct{Int: 1, Strs: []string{"str1", "str2"}, testPrivate: "test private"}
	result := DstStruct{}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StructCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue.Int, result.IntCopy, "result not equal src value")
	require.Equal(t, srcValue.Strs, result.Strs, "result not equal src value")
	require.Equal(t, "", result.testPrivate, "result not equal src value")

}

func TestPrivateFieldStructCopyValue(t *testing.T) {
	defaultCopy = newCopyValue()
	defaultCopy.structCache.copyPrivate = true
	type SrcStruct struct {
		Int  int `json:"int_copy"`
		Strs []string

		testPrivate string
	}

	type DstStruct struct {
		IntCopy     int `json:"int_copy"`
		Strs        []string
		testPrivate string
	}

	srcValue := SrcStruct{Int: 1, Strs: []string{"str1", "str2"}, testPrivate: "test private"}
	result := DstStruct{}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StructCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue.Int, result.IntCopy, "result not equal src value")
	require.Equal(t, srcValue.Strs, result.Strs, "result not equal src value")
	require.Equal(t, srcValue.testPrivate, result.testPrivate, "result not equal src value")

}

func TestInlineStructCopyValue(t *testing.T) {
	type inlineStruct struct {
		A string
		b string
	}
	type srcStruct struct {
		inlineStruct
	}
	type dstStruct struct {
		inlineStruct
	}
	srcValue := srcStruct{inlineStruct{
		A: "inline a",
		b: "inline b private field",
	}}
	result := dstStruct{}

	defaultCopy := newCopyValue()
	defaultCopy.structCache.copyPrivate = true

	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StructCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue.A, result.A, "result not equal src value")
	require.Equal(t, srcValue.b, result.b, "result not equal src value")
}

func TestMapToInlineStructCopyValue(t *testing.T) {
	type inlineStruct struct {
		A string
		b string
	}
	type srcStruct struct {
		inlineStruct
	}
	type dstStruct struct {
		inlineStruct
	}
	srcValue := map[string]string{
		"A": "inline a",
		"b": "inline b private field",
	}

	result := dstStruct{}

	defaultCopy := newCopyValue()
	defaultCopy.structCache.copyPrivate = true

	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.StructCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue["A"], result.A, "result not equal src value")
	require.Equal(t, srcValue["b"], result.b, "result not equal src value")
}

func TestStructToMapMapCopyValue(t *testing.T) {
	type inlineStruct struct {
		A string
		b string
	}

	type dstStruct struct {
		inlineStruct
	}
	srcValue := dstStruct{
		inlineStruct{
			A: "inline a",
			b: "inline b private field",
		},
	}

	result := make(map[string]interface{}, 0)

	defaultCopy := newCopyValue()
	defaultCopy.structCache.copyPrivate = true

	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)
	err := defaultCopy.MapCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue.A, result["A"], "result not equal src value")
	require.Equal(t, srcValue.b, result["b"], "result not equal src value")

	resultStrStr := make(map[string]string, 0)
	src, dst = reflect.ValueOf(srcValue), reflect.ValueOf(&resultStrStr)
	err = defaultCopy.MapCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err)
	require.Equal(t, srcValue.A, resultStrStr["A"], "result not equal src value")
	require.Equal(t, srcValue.b, resultStrStr["b"], "result not equal src value")
}
