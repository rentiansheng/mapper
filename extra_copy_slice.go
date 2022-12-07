package mapper

import (
	"context"
	"fmt"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/9
    @desc:

***************************/

var (
	sliceMapStrAnyType = reflect.TypeOf([]map[string]interface{}{})
	sliceStructType    = reflect.TypeOf([]struct{}{})
)

func (dcv *defaultCopyValue) ToSliceCopyValue(ctx context.Context, fieldName, src, dst reflect.Value) error {
	if src.Kind() != reflect.Slice {
		return fmt.Errorf("ToSliceCopyValue source data type(%s) not slice", src.Kind())
	}
	fieldName = skipElem(fieldName)
	switch src.Type().Elem().Kind() {
	case reflect.Map:
		return dcv.SliceMapToSliceCopyValue(ctx, fieldName, src, dst)
	case reflect.Struct:
		return dcv.SliceStructToSliceCopyValue(ctx, fieldName, src, dst)
	default:
		return fmt.Errorf("ToSliceCopyValue source data type(%s) not support. ssource data type must be slice map or slice struct", src.Kind())
	}

	return nil
}

func (dcv *defaultCopyValue) SliceMapToSliceCopyValue(ctx context.Context, fieldName, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "SliceMapToSliceCopyValue"}
	}
	if dst.Kind() != reflect.Slice {
		return CopyValueError{
			Name:     "SliceMapStrToSliceCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Slice},
			Received: dst,
		}
	}
	src = skipElem(src)
	if (src.Kind() != reflect.Array && src.Kind() != reflect.Slice) ||
		src.Type().Elem().Kind() != reflect.Map || src.Type().Elem().Key().Kind() != fieldName.Kind() {
		return CopyValueError{
			Name:     "SliceMapStrToSliceCopyValue",
			Types:    []reflect.Type{sliceMapStrAnyType},
			Kinds:    nil,
			Received: dst,
		}
	}
	typ := dst.Type().Elem()
	items := make([]reflect.Value, 0, src.Len())
	fn, err := dcv.lookupCopyValue(reflect.New(typ).Elem())
	if err != nil {
		return err
	}
	fieldNameI := fieldName.Interface()
	for idx := 0; idx < src.Len(); idx++ {
		srcItem := src.Index(idx)
		srcMapKeys := srcItem.MapKeys()
		// skip not key elem
		for _, srcMapKey := range srcMapKeys {
			if srcMapKey.Interface() == fieldNameI {
				srcItemValue := srcItem.MapIndex(fieldName)
				dstItem := reflect.New(typ).Elem()
				if err := fn(ctx, srcItemValue, dstItem); err != nil {
					return err
				}
				items = append(items, dstItem)
				break
			}
		}

	}
	dst.SetLen(0)
	dst.Set(reflect.Append(dst, items...))
	return nil
}

func (dcv *defaultCopyValue) SliceStructToSliceCopyValue(ctx context.Context, fieldName, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "SliceStructToSliceCopyValue"}
	}
	if dst.Kind() != reflect.Slice {
		return CopyValueError{
			Name:     "SliceMapStrToSliceCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Slice},
			Received: dst,
		}
	}
	strFieldName := ""
	if fieldName.Kind() != reflect.String {
		return fmt.Errorf("SliceStructToSliceCopyValue field name(type:%v) not strig type", fieldName.Kind())
	}
	strFieldName = fieldName.String()

	src = skipElem(src)
	if (src.Kind() != reflect.Array && src.Kind() != reflect.Slice) ||
		src.Type().Elem().Kind() != reflect.Struct {
		return CopyValueError{
			Name:     "SliceMapStrToSliceCopyValue",
			Types:    []reflect.Type{sliceStructType},
			Kinds:    nil,
			Received: dst,
		}
	}

	srcSD, err := dcv.describeStruct(ctx, src.Type().Elem())
	if err != nil {
		return err
	}
	structFieldDesc, ok := srcSD.fm[strFieldName]
	if !ok {
		return fmt.Errorf("SliceStructToSliceCopyValue field name(type:%v) not found in struct(%s)", fieldName.Kind(), src.Type())
	}

	typ := dst.Type().Elem()
	fn, err := dcv.lookupCopyValue(reflect.New(typ).Elem())
	if err != nil {
		return err
	}
	items := make([]reflect.Value, 0, src.Len())
	for idx := 0; idx < src.Len(); idx++ {
		srcItem := src.Index(idx)
		srvItemValue := srcItem.FieldByName(structFieldDesc.fieldName)
		dstItem := reflect.New(typ).Elem()
		if err := fn(ctx, srvItemValue, dstItem); err != nil {
			return err
		}
		items = append(items, dstItem)
	}
	dst.SetLen(0)
	dst.Set(reflect.Append(dst, items...))
	return nil
}

func (dcv *defaultCopyValue) MapKeyToSliceCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "MapKeyToSliceCopyValue"}
	}
	if dst.Kind() != reflect.Slice {
		return CopyValueError{
			Name:     "MapKeyToSliceCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Slice},
			Received: dst,
		}
	}

	src = skipElem(src)
	if src.Kind() != reflect.Map {
		return CopyValueError{
			Name:     "MapKeyToSliceCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Map},
			Received: dst,
		}
	}

	typ := dst.Type().Elem()
	decode, err := dcv.lookupCopyValue(reflect.New(typ).Elem())
	if err != nil {
		return nil
	}
	keys := src.MapKeys()
	items := make([]reflect.Value, 0, len(keys))
	for _, key := range keys {
		dstItem := reflect.New(typ).Elem()
		if err := decode(ctx, key, dstItem); err != nil {
			return err
		}
		items = append(items, dstItem)
	}
	dst.SetLen(0)
	dst.Set(reflect.Append(dst, items...))

	return nil
}

func (dcv *defaultCopyValue) MapValueToSliceCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "MapValueToSliceCopyValue"}
	}
	if dst.Kind() != reflect.Slice {
		return CopyValueError{
			Name:     "MapValueToSliceCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Slice},
			Received: dst,
		}
	}

	src = skipElem(src)
	if src.Kind() != reflect.Map {
		return CopyValueError{
			Name:     "MapValueToSliceCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Map},
			Received: dst,
		}
	}

	typ := dst.Type().Elem()
	decode, err := dcv.lookupCopyValue(reflect.New(typ).Elem())
	if err != nil {
		return nil
	}
	items := make([]reflect.Value, 0, src.Len())
	mapIter := src.MapRange()
	for mapIter.Next() {
		dstItem := reflect.New(typ).Elem()
		if err := decode(ctx, mapIter.Value(), dstItem); err != nil {
			return err
		}
		items = append(items, dstItem)
	}
	dst.SetLen(0)
	dst.Set(reflect.Append(dst, items...))

	return nil
}

func (dcv *defaultCopyValue) SliceChunk(ctx context.Context, src, dst reflect.Value, size int) error {
	if size <= 0 {
		return fmt.Errorf("slice chunk parameter size must be > 0")
	}
	if !dst.CanSet() {
		return CanSetError{Name: "SliceChunk"}
	}
	if dst.Kind() != reflect.Slice || dst.Type().Elem().Kind() != reflect.Slice {
		return fmt.Errorf("slice chunk can only copy valid and settable from [][]*")
	}
	if dst.IsNil() {
		dst.Set(reflect.New(dst.Type()).Elem())
	}

	sliceType := dst.Type().Elem()
	dstValue := reflect.New(sliceType).Elem()
	items, err := dcv.sliceCopyValue(ctx, src, dstValue)
	if err != nil {
		return err
	}

	var dstItems []reflect.Value
	for idx := 0; idx < len(items); idx += size {
		end := idx + size
		dstItemPart := reflect.New(sliceType).Elem()
		if end > len(items) {
			dstItemPart.Set(reflect.Append(dstItemPart, items[idx:]...))
		} else {
			dstItemPart.Set(reflect.Append(dstItemPart, items[idx:end]...))
		}
		dstItems = append(dstItems, dstItemPart)
	}
	dst.SetLen(0)
	dst.Set(reflect.Append(dst, dstItems...))

	return nil
}
