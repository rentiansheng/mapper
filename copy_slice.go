package mapper

import (
	"context"
	"fmt"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/2
    @desc:

***************************/

func (dcv *defaultCopyValue) SliceCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() || dst.Kind() != reflect.Slice {
		return CopyValueError{Name: "SliceCopyValue", Kinds: []reflect.Kind{reflect.Slice}, Received: dst}
	}

	src = skipElem(src)
	switch src.Kind() {
	case reflect.Slice:
		dstKind := dst.Type().Elem().Kind()
		dstElem := src.Type().Elem()
		if (dstElem.Kind() != dstKind) &&
			(dstElem.Kind() == reflect.Ptr && dstElem.Elem().Kind() != dstKind) {
			// not support copy []*int to []int
			// not support copy []int to []*int
			return fmt.Errorf("cannot copy slice from %s into %s", src.Type(), dst.Type())
		}
	default:
		return fmt.Errorf("cannot copy %v into a slice", src.Type())
	}

	if dst.IsZero() {
		dst.Set(reflect.MakeSlice(dst.Type(), 0, src.Len()))
	}
	typ := dst.Type().Elem()
	items := make([]reflect.Value, 0, src.Len())
	for i := 0; i < src.Len(); i++ {
		itemDst := reflect.New(typ).Elem()
		fn, err := dcv.lookupCopyValue(itemDst)
		if err != nil {
			return err
		}
		if err := fn(ctx, src.Index(i), itemDst); err != nil {
			return err
		}
		items = append(items, itemDst)
	}
	dst.SetLen(0)
	dst.Set(reflect.Append(dst, items...))
	return nil
}

func (dcv *defaultCopyValue) ArrayCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() || dst.Kind() != reflect.Array {
		return CopyValueError{Name: "SliceCopyValue", Kinds: []reflect.Kind{reflect.Slice}, Received: dst}
	}

	if src.Len() > dst.Len() {
		return fmt.Errorf("more elements returned in array than can fit inside %s", dst.Type())
	}

	switch src.Kind() {
	case reflect.Array:
		if src.Elem().Kind() != dst.Elem().Kind() {
			return fmt.Errorf("cannot decode array into %s", src.Type())
		}
	default:
		return fmt.Errorf("cannot copy %v into a array", src.Type())
	}

	typ := dst.Elem().Type()
	// TODO: string to bytes
	for i := 0; i < src.Len(); i++ {
		fn, err := dcv.lookupCopyValue(dst.Elem())
		if err != nil {
			return err
		}
		itemDst := reflect.New(typ)
		if err := fn(ctx, src.Index(i), itemDst); err != nil {
			return err
		}
		dst.Index(i).Set(itemDst)

	}
	return nil
}
