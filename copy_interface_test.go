package mapper

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

/***************************
    @author: tiansheng.ren
    @date: 7/25/23
    @desc:

***************************/

func TestCopyInterfaceUseRawTypeIssue21(t *testing.T) {
	getCopyResult := func() interface{} {
		return make(map[string]int, 0)
	}

	result := getCopyResult()
	srcValue := map[string]interface{}{
		"int":   int(1),
		"int8":  int8(1),
		"int64": int64(1),
		"float": float64(1),
	}
	dstKv := map[string]int{
		"int":   1,
		"int8":  1,
		"int64": 1,
		"float": 1,
	}
	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)

	err := defaultCopy.InterfaceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err, "TestCopyInterfaceUseRawTypeIssue21")
	require.Equal(t, dstKv, result, "TestCopyInterfaceUseRawTypeIssue21")
}

func TestCopyInterfaceUseRawTypeBuildInIssue21(t *testing.T) {
	getCopyResult := func() interface{} {
		var i int64
		return i
	}

	result := getCopyResult()
	srcValue := 1.0

	src, dst := reflect.ValueOf(srcValue), reflect.ValueOf(&result)

	err := defaultCopy.InterfaceCopyValue(ctx, src, dst.Elem())
	require.NoError(t, err, "TestCopyInterfaceUseRawTypeBuildInIssue21")
	require.Equal(t, int64(1), result, "TestCopyInterfaceUseRawTypeBuildInIssue21")
}
