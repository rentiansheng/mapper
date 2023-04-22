package mapper

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

/*
**************************

	@author: tiansheng.ren
	@date: 2022/10/25
	@desc:

**************************
*/
var (
	tagName            = "json"
	excludeTagValue    = "-"
	copyValueTagPrefix = "copy"
)

type fieldDescription struct {
	fieldName string
	name      string
	idx       int
	inline    bool
	omitEmpty bool
	private   bool
	typ       reflect.Type
}

type structDescription struct {
	fm map[string]*fieldDescription
}

type structCache struct {
	cache map[reflect.Type]*structDescription
	sync.RWMutex
	copyPrivate    bool
	validateStruct bool
}

func newStructCache() *structCache {
	return &structCache{
		cache:       make(map[reflect.Type]*structDescription, 0),
		copyPrivate: false,
	}
}

func (dcv *defaultCopyValue) StructCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "StructCopyValue"}
	}
	if dst.Kind() != reflect.Struct {
		return CopyValueError{Name: "copyStruct.StructCopyValue", Kinds: []reflect.Kind{reflect.Struct}, Received: dst}
	}

	src = skipElem(src)
	switch src.Kind() {
	case reflect.Struct:
	case reflect.Map:
		// copy map to struct
		return dcv.MapToStructCopyValue(ctx, src, dst)
	default:
		return CopyValueError{
			Name:     "copyStruct.StructCopyValue",
			Types:    nil,
			Kinds:    []reflect.Kind{reflect.Map, reflect.Struct},
			Received: dst,
		}
	}
	if dst.IsZero() {
		dst.Set(reflect.New(dst.Type()).Elem())
	}
	srcSD, err := dcv.describeStruct(ctx, src.Type())
	if err != nil {
		return err
	}
	dstSD, err := dcv.describeStruct(ctx, dst.Type())
	if err != nil {
		return err
	}
	for name, descField := range dstSD.fm {
		srcDescField, ok := srcSD.fm[name]
		if !ok {
			// not found, continue
			continue
		}

		srcItem, dstItem := src.FieldByName(srcDescField.fieldName), dst.FieldByName(descField.fieldName)
		if descField.private {
			if !dcv.structCache.copyPrivate {
				continue
			}
			dstItem = reflect.NewAt(dstItem.Type(), unsafe.Pointer(dstItem.UnsafeAddr())).Elem()
		}
		if dstItem.Kind() == reflect.Interface && dstItem.IsNil() {
			dstItem.Set(reflect.New(descField.typ).Elem())
		}
		fn, err := dcv.lookupCopyValue(dstItem)
		if err != nil {
			return err
		}
		if err := fn(ctx, srcItem, dstItem); err != nil {
			return err
		}

	}
	if dcv.validate != nil {
		if err := dcv.validate.Struct(dst.Interface()); err != nil {
			return err
		}

	}

	return nil

}

func (dcv *defaultCopyValue) MapToStructCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "MapToStructCopyValue"}
	}
	if dst.Kind() != reflect.Struct {
		return CopyValueError{Name: "copyStruct.MapToStructCopyValue", Kinds: []reflect.Kind{reflect.Struct}, Received: dst}
	}
	if src.Kind() != reflect.Map || src.Type().Key().Kind() != reflect.String {
		return CopyValueError{Name: "copyStruct.MapToStructCopyValue", Kinds: []reflect.Kind{reflect.Map}, Received: dst}
	}

	src = skipElem(src)
	if dst.IsZero() {
		dst.Set(reflect.New(dst.Type()).Elem())
	}
	dstSD, err := dcv.describeStruct(ctx, dst.Type())
	if err != nil {
		return err
	}
	iter := src.MapRange()
	for iter.Next() {
		key := iter.Key().String()
		value := iter.Value()
		descField, ok := dstSD.fm[key]
		if !ok {
			continue
		}

		dstVal := dst.FieldByName(descField.fieldName)
		if descField.private {
			if !dcv.structCache.copyPrivate {
				continue
			}
			dstVal = reflect.NewAt(dstVal.Type(), unsafe.Pointer(dstVal.UnsafeAddr())).Elem()
		}
		fn, err := dcv.lookupCopyValue(dstVal)
		if err != nil {
			return err
		}
		if err := fn(ctx, value, dstVal); err != nil {
			return err
		}

	}
	return nil

}

func (dcv *defaultCopyValue) StructToMapCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "StructToMapCopyValue"}
	}
	if dst.Kind() != reflect.Map || dst.Type().Key().Kind() != reflect.String {
		return CopyValueError{Name: "copyStruct.CopyStructMapValue", Kinds: []reflect.Kind{reflect.Struct}, Received: dst}
	}

	src = skipElem(src)
	if src.Kind() != reflect.Struct {
		return fmt.Errorf("cannot copy %v into struct type", src.Type())
	}
	if dst.IsZero() {
		dst.Set(reflect.New(dst.Type()).Elem())
	}
	srcSD, err := dcv.describeStruct(ctx, src.Type())
	if err != nil {
		return err
	}
	myIsNil := func(v reflect.Value) bool {

		k := v.Kind()
		switch k {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			return v.IsNil()

		}
		return false
	}
	for name, descField := range srcSD.fm {
		if descField.private && !dcv.structCache.copyPrivate {
			//fix map to struct 拷贝私有对象
			continue
		}
		fieldSrc := src.FieldByName(descField.fieldName)
		// copy struct to map, if the field is empty and tag has omitempty, we skip it.
		if descField.omitEmpty && myIsNil(fieldSrc) {
			continue
		}

		dstVal := reflect.New(dst.Type().Elem()).Elem()
		fn, err := dcv.lookupCopyValue(dstVal)
		if err != nil {
			return nil
		}
		if err := fn(ctx, fieldSrc, dstVal); err != nil {
			return err
		}
		dst.SetMapIndex(reflect.ValueOf(name), dstVal)

	}
	return nil
}

func (dcv *defaultCopyValue) describeStruct(ctx context.Context, t reflect.Type) (*structDescription, error) {
	// We need to analyze the struct, including getting the tags, collecting
	// information about inlining, and create a map of the field name to the field.
	dcv.structCache.RLock()
	ds, exists := dcv.structCache.cache[t]
	dcv.structCache.RUnlock()
	if exists {
		return ds, nil
	}
	numFields := t.NumField()
	sd := &structDescription{
		fm: make(map[string]*fieldDescription, numFields),
	}
	for i := 0; i < numFields; i++ {
		sf := t.Field(i)

		desc := &fieldDescription{
			fieldName: sf.Name,
			name:      sf.Name,
			idx:       i,
			typ:       sf.Type,
		}
		if !sf.IsExported() {
			desc.private = true
		}

		if sf.Anonymous {
			inlineSF, err := dcv.describeStruct(ctx, sf.Type)
			if err != nil {
				return nil, err
			}
			for _, fd := range inlineSF.fm {
				if _, exists := sd.fm[fd.name]; exists {
					return nil, fmt.Errorf("(struct %s) duplicated key %s", t.String(), fd.name)
				}
				sd.fm[fd.name] = fd
			}
			continue
		}
		for _, fn := range tagHandleFnList {
			if tag := fn.newFn(sf.Tag.Get(fn.tagName)); tag.Name() != "" {
				desc.name = tag.Name()
				desc.omitEmpty = tag.OmitEmpty()

				break
			}
		}

		if _, exists := sd.fm[desc.name]; exists {
			return nil, fmt.Errorf("(struct %s) duplicated key %s", t.String(), desc.name)
		}
		sd.fm[desc.name] = desc
	}
	dcv.structCache.Lock()
	dcv.structCache.cache[t] = sd
	dcv.structCache.Unlock()

	return sd, nil
}
