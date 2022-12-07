package mapper

import (
	"context"
	"fmt"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/24
    @desc:

***************************/

var (
	mapperCV    *defaultCopyValue
	allMapperCV *defaultCopyValue
)

func init() {
	mapperCV = newCopyValue()
	allMapperCV = newCopyValue()
	allMapperCV.structCache.copyPrivate = true
}

func Mapper(ctx context.Context, src, dst interface{}) error {
	return mapperHandler(ctx, mapperCV, src, dst)
}

func AllMapper(ctx context.Context, src, dst interface{}) error {
	return mapperHandler(ctx, allMapperCV, src, dst)
}

func ObjectsField(ctx context.Context, field, src, dst interface{}) error {
	return mapperCV.ToSliceCopyValue(ctx, reflect.ValueOf(field), reflect.ValueOf(src), reflect.ValueOf(dst).Elem())
}

func MapKeys(ctx context.Context, src, dst interface{}) error {
	return mapperCV.MapKeyToSliceCopyValue(ctx, reflect.ValueOf(src), reflect.ValueOf(dst).Elem())
}

func MapValues(ctx context.Context, src, dst interface{}) error {
	return mapperCV.MapValueToSliceCopyValue(ctx, reflect.ValueOf(src), reflect.ValueOf(dst).Elem())

}

func Chunk(ctx context.Context, src, dst interface{}, size int) error {
	dstV := reflect.ValueOf(dst)
	if dstV.Kind() != reflect.Ptr {
		return fmt.Errorf("copy to object must be pointer")
	}

	dstV = dstV.Elem()
	if !dstV.CanSet() {
		return fmt.Errorf("copy to object must be can set")
	}

	srcV := reflect.ValueOf(src)

	return mapperCV.SliceChunk(ctx, srcV, dstV, size)
}

func mapperHandler(ctx context.Context, dcv *defaultCopyValue, src, dst interface{}) error {
	dstV := reflect.ValueOf(dst)
	if dstV.Kind() != reflect.Ptr {
		return fmt.Errorf("copy to object must be pointer")
	}

	srcV := reflect.ValueOf(src)
	if dstV.Kind() == reflect.Invalid {
		if srcV.Kind() == reflect.Invalid {
			return nil
		}
		return fmt.Errorf("copy to object is nil")
	}
	dstV = dstV.Elem()
	if !dstV.CanSet() {
		return CopyValueError{Name: "mapper", Kinds: []reflect.Kind{reflect.Bool}, Received: dstV}
	}

	fn, err := dcv.lookupCopyValue(dstV)
	if err != nil {
		return err
	}
	return fn(ctx, srcV, dstV)
}

type mapper interface {
	Mapper(ctx context.Context, src, dst interface{}) error
}

func NewMapper(options ...Option) mapper {
	option := mergeOption(options...)
	handler := newCopyValue()
	if option.copyPrivate != nil {
		handler.structCache.validateStruct = *option.copyPrivate
	}
	if option.validateStruct != nil && *option.validateStruct == true {
		// Notice: copy tag change, here need change
		handler.validate = newValidateStruct()
	}
	return &mapperInstance{
		cv: handler,
	}
}

type mapperInstance struct {
	cv *defaultCopyValue
}

func (m mapperInstance) Mapper(ctx context.Context, src, dst interface{}) error {
	return mapperHandler(ctx, m.cv, src, dst)
}
