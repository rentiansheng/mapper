package mapper

import (
	"context"
	"fmt"
	"github.com/rentiansheng/mapper/mtype"
	"math"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/24
    @desc:

***************************/

var (
	defaultCopyValueHanler = newCopyValue()
)

func newCopyValue() *defaultCopyValue {
	dcv := &defaultCopyValue{
		structCache: newStructCache(),
		register:    newRegister(),
	}
	dcv.buildDefaultRegistry()
	return dcv
}

type defaultCopyValue struct {
	structCache *structCache

	register *register
}

func (dcv *defaultCopyValue) BooleanCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.IsValid() || !dst.CanSet() || dst.Kind() != reflect.Bool && dst.Kind() != reflect.Interface {
		return CopyValueError{Name: "BooleanCopyValue", Kinds: []reflect.Kind{reflect.Bool}, Received: dst}
	}

	src = skipPtrElem(src)
	if src.Kind() != dst.Kind() {
		return fmt.Errorf("cannot copy %v into a boolean", src.Type())
	}

	dst.SetBool(src.Bool())
	return nil
}

func (dcv *defaultCopyValue) IntCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CopyValueError{
			Name:     "IntCopyValue",
			Kinds:    []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int},
			Received: dst,
		}
	}

	src = skipPtrElem(src)
	var i64 int64
	switch src.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64 = src.Int()
	case reflect.Float32, reflect.Float64:
		f64 := src.Float()
		if f64 < float64(math.MinInt64) || f64 > float64(math.MaxInt64) {
			return fmt.Errorf("%g overflows int64", f64)
		}
		i64 = int64(f64)
	default:
		return fmt.Errorf("cannot copy %v into an integer type", src.Type())
	}
	switch dst.Kind() {
	case reflect.Int8:
		if i64 < math.MinInt8 || i64 > math.MaxInt8 {
			return fmt.Errorf("%d overflows int8", i64)
		}
	case reflect.Int16:
		if i64 < math.MinInt16 || i64 > math.MaxInt16 {
			return fmt.Errorf("%d overflows int16", i64)
		}
	case reflect.Int32:
		if i64 < math.MinInt32 || i64 > math.MaxInt32 {
			return fmt.Errorf("%d overflows int32", i64)
		}
	case reflect.Int64:
	case reflect.Int:
		if int64(int(i64)) != i64 { // Can we fit this inside of an int
			return fmt.Errorf("%d overflows int", i64)
		}
	default:
		return CopyValueError{
			Name:     "IntCopyValue",
			Kinds:    []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int},
			Received: dst,
		}
	}

	dst.SetInt(i64)
	return nil
}

func (dcv *defaultCopyValue) UintCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CopyValueError{
			Name:     "IntCopyValue",
			Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
			Received: dst,
		}
	}

	src = skipPtrElem(src)
	var i64 uint64
	switch src.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i64 = src.Uint()
	case reflect.Float32, reflect.Float64:
		f64 := src.Float()
		if f64 > float64(math.MaxUint64) {
			return fmt.Errorf("%g overflows uint64", f64)
		}
		i64 = uint64(f64)
	default:
		return fmt.Errorf("cannot copy %v into an unsigned integer type", src.Type())
	}
	switch dst.Kind() {
	case reflect.Uint8:
		if i64 > math.MaxUint8 {
			return fmt.Errorf("%d overflows uint8", i64)
		}
	case reflect.Int16:
		if i64 > math.MaxUint16 {
			return fmt.Errorf("%d overflows uint16", i64)
		}
	case reflect.Uint32:
		if i64 > math.MaxUint32 {
			return fmt.Errorf("%d overflows uint32", i64)
		}
	case reflect.Uint64:
	case reflect.Uint:
		if uint64(uint(i64)) != i64 { // Can we fit this inside of an int
			return fmt.Errorf("%d overflows uint", i64)
		}

	default:
		return CopyValueError{
			Name:     "IntCopyValue",
			Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
			Received: dst,
		}
	}

	dst.SetUint(i64)
	return nil
}

func (dcv *defaultCopyValue) FloatCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CopyValueError{Name: "FloatCopyValue", Kinds: []reflect.Kind{reflect.Float32, reflect.Float64}, Received: dst}
	}

	src = skipPtrElem(src)
	var f float64
	switch src.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f = float64(src.Int())
	case reflect.Float32, reflect.Float64:
		f = src.Float()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f = float64(src.Uint())
	default:
		return CopyValueError{
			Name: "FloatCopyValue",
			Kinds: []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64},
			Received: dst,
		}
	}

	switch dst.Kind() {
	case reflect.Float32:
		if float64(float32(32)) != f {
			return fmt.Errorf("%g overflows float32", f)
		}
	case reflect.Float64:
	default:
		return CopyValueError{Name: "FloatCopyValue", Kinds: []reflect.Kind{reflect.Float32, reflect.Float64}, Received: dst}
	}

	dst.SetFloat(f)
	return nil
}

func (dcv *defaultCopyValue) StringCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CopyValueError{Name: "StringCopyValue", Kinds: []reflect.Kind{reflect.String}, Received: dst}
	}

	src = skipPtrElem(src)
	var str string
	switch src.Kind() {
	case reflect.String:
		str = src.String()
	default:
		if src.Type() == mtype.ByteSliceType {
			if !src.IsZero() {
				str = string(src.Interface().([]byte))
			}
		} else {
			return fmt.Errorf("cannot copy %v into a string type", src.Type())
		}
	}

	dst.SetString(str)
	return nil
}

func (dcv *defaultCopyValue) SliceCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() || dst.Kind() != reflect.Slice {
		return CopyValueError{Name: "SliceCopyValue", Kinds: []reflect.Kind{reflect.Slice}, Received: dst}
	}

	src = skipPtrElem(src)
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

/*
func (dcv *defaultCopyValue) TimeCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() || dst.Type() != mtype.TimeType {
		return CopyValueError{
			Name:     "InterfaceCopyValue",
			Types:    []reflect.Type{dst.Type()},
			Received: dst,
		}
	}

	if src.Type() == mtype.TimeType {
		reflect.Copy(dst, src)
		return nil
	} else if src.Kind() == reflect.Int || src.Kind() == reflect.Int32 || src.Kind() == reflect.Int64 {
		//  TODO: 如何区分毫秒和秒, 是否可以采用int32 最大值
		sec, nsec := src.Int(), src.Int()
		if sec > math.MaxInt32 {
			sec = sec / 1000
			nsec = nsec % 1000
		}
		time.Unix(sec, nsec)
		return nil
	} else if src.Kind() == reflect.Uint || src.Kind() == reflect.Uint32 || src.Kind() == reflect.Uint64 {
		sec, nsec := src.Uint(), src.Uint()
		if sec > math.MaxInt32 {
			sec = sec / 1000
			nsec = nsec % 1000
		}
		time.Unix(int64(sec), int64(nsec))
		return nil
	} else {

		return CopyValueError{
			Name:     "CopyTimeValue",
			Types:    []reflect.Type{dst.Type()},
			Received: dst,
		}
	}
}*/

func (dcv *defaultCopyValue) PtrCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() || dst.Kind() != reflect.Ptr {
		return CopyValueError{Name: "PtrCopyValue", Kinds: []reflect.Kind{reflect.Ptr}, Received: dst}
	}

	if src.IsZero() {
		dst.Set(reflect.Zero(dst.Type()))
		return nil
	}

	if dst.IsNil() {
		dst.Set(reflect.New(dst.Type().Elem()))
	}
	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}
	fn, err := dcv.lookupCopyValue(dst.Elem())
	if err != nil {
		return err
	}
	return fn(ctx, src, dst.Elem())
}

func skipPtrElem(elem reflect.Value) reflect.Value {
	for elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	return elem
}
