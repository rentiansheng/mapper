package mapper

import (
	"strings"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/1
    @desc:

***************************/

type tagHandleS struct {
	tagName string
	newFn   NewTagHandleFn
}

var tagHandleFnList = []tagHandleS{
	{"copy", newCopyTag},
	{"json", newJSONTag},
	{"gorm", newGormTag},
}

type NewTagHandleFn func(string) TagHande

type TagHande interface {
	Name() string
}

type structTagInline struct {
	name string
}

func newJSONTag(tag string) TagHande {
	sti := structTagInline{}
	tags := strings.Split(tag, ",")
	for idx := range tags {
		tags[idx] = strings.TrimSpace(tags[idx])
	}
	if tags[0] != "" && tags[0] != excludeTagValue {
		sti.name = strings.TrimSpace(tags[0])
	}
	// json tag  copy value
	for _, item := range tags[0:] {
		if itemVals := strings.Split(item, "="); strings.TrimSpace(itemVals[0]) == copyValueTagPrefix {
			if len(itemVals) > 1 && itemVals[1] != "" {
				sti.name = itemVals[1]
			}
		}
	}

	return jsonTag{structTagInline: sti}
}

type jsonTag struct {
	structTagInline
}

var _ TagHande = (*jsonTag)(nil)

func newCopyTag(tag string) TagHande {
	ct := copyTag{}
	tags := strings.Split(tag, ",")
	ct.name = tags[0]

	return ct
}

type copyTag struct {
	structTagInline
	to string
}

func (j copyTag) To() string {
	return j.to
}

var _ TagHande = (*copyTag)(nil)

func newGormTag(tag string) TagHande {
	// primary_key;column:id;type:int(10);not null
	sti := structTagInline{}
	tags := strings.Split(tag, ";")
	for _, tag := range tags {
		if strings.HasPrefix(strings.ToLower(tag), "column:") {
			tagItems := strings.Split(tag, ":")
			if len(tagItems) > 1 {
				sti.name = tagItems[1]
				break
			}
		}
	}

	return jsonTag{structTagInline: sti}
}

type gormTag struct {
	structTagInline
}

var _ TagHande = (*gormTag)(nil)

func (j structTagInline) Name() string {
	return j.name
}
