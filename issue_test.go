package mapper

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
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
