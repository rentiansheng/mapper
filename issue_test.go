package mapper

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

/***************************
    @author: tiansheng.ren
    @date: 5/15/23
    @desc:

***************************/

func TestIssue20(t *testing.T) {
	// err: InterfaceCopyValue can only copy valid and settable map[], but got map[string]interface {}
	testJSON := `{"name":"issue id 20", "_dc_fixed":null}`
	info := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(testJSON), &info)
	require.NoError(t, err)
	fixedMeta := struct {
		FixedInfo map[string]interface{} `json:"_dc_fixed"`
	}{
		FixedInfo: map[string]interface{}{},
	}
	err = Mapper(ctx, info, &fixedMeta)
	require.NoError(t, err)
}

func TestIssue25(t *testing.T) {

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
	err := Mapper(ctx, src, &dst)
	require.NoError(t, err)

	require.Equal(t, src.Name, dst["name"])
	require.Equal(t, src.Age, dst["age"])
	require.Equal(t, reflect.Struct, reflect.ValueOf(dst["sub"]).Kind())
	require.Equal(t, src.Sub.Sub, dst["sub"].(copySubS).Sub)

}
