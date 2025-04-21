package mapper

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

/***************************
    @author: tiansheng.ren
    @date: 2025/4/21
    @desc:

***************************/

func getCopyStruct() *defaultCopyValue {
	dcv := &defaultCopyValue{
		structCache: newStructCache(),
		register:    newRegister(),
	}
	return dcv
}

func TestStructCopy1(t *testing.T) {

	type copySubS struct {
		Sub string `json:"sub"`
	}
	type copyS struct {
		Name string   `json:"name"`
		Age  int      `json:"age"`
		Sub  copySubS `json:"sub"`
	}

	src := copyS{Name: "test", Age: 10, Sub: copySubS{Sub: "test"}}
	dst := copyS{}
	ctx := context.Background()
	err := Mapper(ctx, src, &dst)
	require.NoError(t, err)
	require.Equal(t, src.Name, dst.Name)
	require.Equal(t, src.Age, dst.Age)
	require.Equal(t, src.Sub.Sub, dst.Sub.Sub)
}

func TestStructCopy(t *testing.T) {

	type copySubS struct {
		Sub string `json:"sub"`
	}
	type copyS struct {
		Name string   `json:"name"`
		Age  int      `json:"age"`
		Sub  copySubS `json:"sub"`
	}

	src := copyS{Name: "test", Age: 10, Sub: copySubS{Sub: "test"}}
	dst := map[string]interface{}{}
	ctx := context.Background()
	err := Mapper(ctx, src, &dst)
	require.NoError(t, err)

}
