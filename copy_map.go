package mapper

import (
	"context"
	"fmt"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/25
    @desc:

***************************/

func (dcv *defaultCopyValue) MapCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "MapCopyValue"}
	}
	if dst.Kind() != reflect.Map {
		return LookupCopyValueError{Name: "MapCopyValue", Kinds: []reflect.Kind{reflect.Map}, Received: dst}
	}
	src = skipElem(src)
	switch src.Kind() {
	case reflect.Map:
	case reflect.Struct:
		return dcv.StructToMapCopyValue(ctx, src, dst)
	default:
		if src.Kind() == reflect.Invalid || src.IsNil() {
			// dst is map, src is nil, just return
			return nil
		}
		return CopyValueError{
			Name:     "MapCopyValue",
			Kinds:    []reflect.Kind{reflect.Map},
			Received: src,
		}

	}

	if src.Type().Key() != dst.Type().Key() {
		return fmt.Errorf("cannot copy map[%v]%v into an map[%v][%v] type",
			src.Type().Key(), src.Type().Elem(), dst.Type().Key(), dst.Type().Elem())
	}
	if dst.IsNil() {
		dst.Set(reflect.MakeMap(dst.Type()))
	}

	keyType := dst.Type().Key()
	valueType := dst.Type().Elem()
	fnKey, err := dcv.lookupCopyValue(reflect.New(keyType).Elem())
	if err != nil {
		return err
	}
	fnValue, err := dcv.lookupCopyValue(reflect.New(valueType).Elem())
	if err != nil {
		return err
	}

	iter := src.MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		newKey := reflect.New(keyType).Elem()
		newValue := reflect.New(valueType).Elem()
		if err := fnKey(ctx, key, newKey); err != nil {
			return err
		}

		if err := fnValue(ctx, value, newValue); err != nil {
			return err
		}
		dst.SetMapIndex(newKey, newValue)
	}
	return nil
}
