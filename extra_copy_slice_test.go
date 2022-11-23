package mapper

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"sort"
	"testing"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/11
    @desc:

***************************/

func TestSliceMapToSliceCopyValue(t *testing.T) {
	items := []map[int]int{{1: 1}, {1: 2}, {1: 3}, {1: 4}, {1: 5}, {1: 6}}

	dst := make([]int, 0)
	err := defaultCopy.SliceMapToSliceCopyValue(ctx, reflect.ValueOf(int(1)), reflect.ValueOf(items), reflect.ValueOf(&dst).Elem())
	require.NoError(t, err, "TestSliceMapToSliceCopyValue")
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, dst, "TestSliceMapToSliceCopyValue")
}

func TestSkipNotFoundElemSliceMapToSliceCopyValue(t *testing.T) {
	items := []map[string]int{{"field": 1}, {"field": 2}, {"field": 3}, {"field": 4}, {"field1": 5}, {"field1": 6}}

	dst := make([]int, 0)
	err := defaultCopy.SliceMapToSliceCopyValue(ctx, reflect.ValueOf("field"), reflect.ValueOf(items), reflect.ValueOf(&dst).Elem())
	require.NoError(t, err, "TestSliceMapToSliceCopyValue")
	require.Equal(t, []int{1, 2, 3, 4}, dst, "TestSliceMapToSliceCopyValue")
}

func TestSliceStructToSliceCopyValue(t *testing.T) {
	items := []struct {
		ID int `json:"id"`
	}{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}}
	dst := make([]int, 0)
	err := defaultCopy.SliceStructToSliceCopyValue(ctx, reflect.ValueOf("id"), reflect.ValueOf(items), reflect.ValueOf(&dst).Elem())
	require.NoError(t, err, "TestSliceStructToSliceCopyValue")
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, dst, "TestSliceStructToSliceCopyValue")
}

func TestExtraMapIntKeyToSliceCopyValue(t *testing.T) {
	src := map[int]int{1: 1, 2: 2, 3: 3}
	keyDst, valueDst := make([]int, 0), make([]int, 0)
	err := defaultCopy.MapKeyToSliceCopyValue(ctx, reflect.ValueOf(src), reflect.ValueOf(&keyDst).Elem())
	require.NoError(t, err, "TestMapKeyCopyValue")
	sort.Ints(keyDst)
	require.Equal(t, []int{1, 2, 3}, keyDst, "TestMapKeyCopyValue")

	err = defaultCopy.MapValueToSliceCopyValue(ctx, reflect.ValueOf(src), reflect.ValueOf(&valueDst).Elem())
	require.NoError(t, err, "TestMapKeyCopyValue")
	sort.Ints(valueDst)
	require.Equal(t, []int{1, 2, 3}, valueDst, "TestMapKeyCopyValue")

}

func TestExtraMapStrKeyToSliceCopyValue(t *testing.T) {
	src := map[string]int{"str": 1}
	keyDst, valueDst := make([]string, 0), make([]int, 0)
	err := defaultCopy.MapKeyToSliceCopyValue(ctx, reflect.ValueOf(src), reflect.ValueOf(&keyDst).Elem())
	require.NoError(t, err, "TestMapKeyCopyValue")
	require.Equal(t, []string{"str"}, keyDst, "TestMapKeyCopyValue")

	err = defaultCopy.MapValueToSliceCopyValue(ctx, reflect.ValueOf(src), reflect.ValueOf(&valueDst).Elem())
	require.NoError(t, err, "TestMapKeyCopyValue")
	require.Equal(t, []int{1}, valueDst, "TestMapKeyCopyValue")
}
