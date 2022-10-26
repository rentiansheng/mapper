package mapper

import (
	"context"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/24
    @desc:

***************************/

type Option struct {
	copyPrivate bool
}

var (
	mapper    *defaultCopyValue
	allMapper *defaultCopyValue
)

func init() {
	mapper = newCopyValue()
	allMapper = newCopyValue()
	allMapper.structCache.copyPrivate = true
}

func Mapper(src, dst interface{}) error {
	return mapperHandler(mapper, src, dst)
}

func AllMapper(src, dst interface{}) error {
	return mapperHandler(allMapper, src, dst)
}

func mapperHandler(dcv *defaultCopyValue, src, dst interface{}) error {
	dstV := reflect.ValueOf(dst).Elem()
	if !dstV.CanSet() {
		return CopyValueError{Name: "mapper", Kinds: []reflect.Kind{reflect.Bool}, Received: dstV}
	}
	srcV := reflect.ValueOf(src)
	fn, err := dcv.lookupCopyValue(dstV)
	if err != nil {
		return err
	}
	return fn(context.TODO(), srcV, dstV)
}
