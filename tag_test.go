package mapper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/2
    @desc:

***************************/

func TestCopyTag(t *testing.T) {
	tag := newCopyTag("id")
	require.Equal(t, "id", tag.Name())
	tag = newCopyTag("")
	require.Equal(t, "", tag.Name())
}

func TestJSONTag(t *testing.T) {
	tag := newJSONTag("id")
	require.Equal(t, "id", tag.Name())
	tag = newJSONTag("id,copy=copy_id")
	require.Equal(t, "copy_id", tag.Name())
	tag = newJSONTag("")
	require.Equal(t, "", tag.Name())
}

func TestGormTag(t *testing.T) {
	tag := newGormTag("primary_key;column:id;type:int(10);not null")
	require.Equal(t, "id", tag.Name())
	tag = newGormTag("column:id;type:int(10);not null")
	require.Equal(t, "id", tag.Name())
	tag = newGormTag("")
	require.Equal(t, "", tag.Name())
}
