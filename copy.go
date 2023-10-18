package mapper

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/rentiansheng/mapper/mtype"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/24
    @desc:

***************************/

var (
	defaultCopyValueHandler = newCopyValue()
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
	validate Validator
}

func (dcv *defaultCopyValue) BooleanCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "BooleanCopyValue"}
	}
	if !dst.IsValid() || dst.Kind() != reflect.Bool && dst.Kind() != reflect.Interface {
		return LookupCopyValueError{Name: "BooleanCopyValue", Kinds: []reflect.Kind{reflect.Bool}, Received: dst}
	}

	src = skipElem(src)
	if src.Kind() == reflect.Invalid {
		return nil
	}
	if src.Kind() != reflect.Bool {
		return CopyValueError{Name: "BooleanCopyValue", Kinds: []reflect.Kind{reflect.Bool}, Received: src}
	}

	dst.SetBool(src.Bool())
	return nil
}

func (dcv *defaultCopyValue) IntCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "IntCopyValue"}
	}

	src = skipElem(src)
	var i64 int64
	switch src.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64 = src.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ui64 := src.Uint()
		if ui64 > math.MaxInt64 {
			return fmt.Errorf("%v overflows int64", ui64)
		}
		i64 = int64(ui64)
	case reflect.Float32, reflect.Float64:
		f64 := src.Float()
		if f64 < float64(math.MinInt64) || f64 > float64(math.MaxInt64) {
			return fmt.Errorf("%g overflows int64", f64)
		}
		i64 = int64(f64)
	default:
		if src.Kind() == reflect.Invalid {
			return nil
		}
		switch src.Type() {
		case mtype.JSONNumber:
			jsonNumber := src.Interface().(json.Number)
			var err error
			i64, err = jsonNumber.Int64()
			if err != nil {
				return err
			}

		default:
			return CopyValueError{
				Name: "IntCopyValue",
				// no all kind. allow reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,reflect.Float32, reflect.Float64 no overflows value
				Kinds:    []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64},
				Received: src,
			}
		}
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
		return LookupCopyValueError{
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
		return CanSetError{Name: "UintCopyValue"}
	}

	src = skipElem(src)
	var i64 uint64
	switch src.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i64 = src.Uint()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tmpI64 := src.Int()
		if tmpI64 < 0 {
			return fmt.Errorf("%d overflows uint64", tmpI64)
		}
		i64 = uint64(tmpI64)
	case reflect.Float32, reflect.Float64:
		f64 := src.Float()
		if f64 > float64(math.MaxUint64) {
			return fmt.Errorf("%g overflows uint64", f64)
		}
		i64 = uint64(f64)
	default:
		if src.Kind() == reflect.Invalid {
			return nil
		}
		switch src.Type() {
		case mtype.JSONNumber:
			jsonNumber := src.Interface().(json.Number)
			var err error
			i64, err = strconv.ParseUint(string(jsonNumber), 10, 64)
			if err != nil {
				return err
			}

		default:
			return CopyValueError{
				Name: "UintCopyValue",
				// no all kind. allow reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,reflect.Float32, reflect.Float64 no overflows value
				Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
				Received: src,
			}
		}
	}
	switch dst.Kind() {
	case reflect.Uint8:
		if i64 > math.MaxUint8 {
			return fmt.Errorf("%d overflows uint8", i64)
		}
	case reflect.Uint16:
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
		return LookupCopyValueError{
			Name:     "UintCopyValue",
			Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
			Received: dst,
		}
	}

	dst.SetUint(i64)
	return nil
}

func (dcv *defaultCopyValue) FloatCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "FloatCopyValue"}
	}

	src = skipElem(src)
	var f float64
	switch src.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f = float64(src.Int())
	case reflect.Float32, reflect.Float64:
		f = src.Float()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f = float64(src.Uint())
	default:
		if src.Kind() == reflect.Invalid {
			return nil
		}
		switch src.Type() {
		case mtype.JSONNumber:
			jsonNumber := src.Interface().(json.Number)
			var err error
			f, err = jsonNumber.Float64()
			if err != nil {
				return err
			}

		default:
			return CopyValueError{
				Name: "FloatCopyValue",
				Kinds: []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
					reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64},
				Received: src,
			}
		}
	}

	switch dst.Kind() {
	case reflect.Float32:
		if float64(float32(32)) != f {
			return fmt.Errorf("%g overflows float32", f)
		}
	case reflect.Float64:
	default:
		return LookupCopyValueError{Name: "FloatCopyValue", Kinds: []reflect.Kind{reflect.Float32, reflect.Float64}, Received: dst}
	}

	dst.SetFloat(f)
	return nil
}

func (dcv *defaultCopyValue) StringCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name: "StringCopyValue"}
	}

	src = skipElem(src)
	var str string
	switch src.Kind() {
	case reflect.String:
		str = src.String()
	default:
		if src.Kind() == reflect.Invalid {
			return nil
		}
		// allow []byte to string
		if src.Type() == mtype.ByteSliceType {
			if !src.IsZero() {
				str = string(src.Interface().([]byte))
			}
		} else {
			return CopyValueError{
				Name:     "StringCopyValue",
				Types:    []reflect.Type{mtype.ByteSliceType},
				Kinds:    []reflect.Kind{reflect.String},
				Received: src,
			}
		}
	}
	if dst.Kind() != reflect.String {
		return LookupCopyValueError{Name: "StringCopyValue", Kinds: []reflect.Kind{reflect.String}, Received: dst}
	}

	dst.SetString(str)
	return nil
}

/*
func (dcv *defaultCopyValue) TimeCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() {
		return CanSetError{Name:"TimeCopyValue"}
	}
	if   dst.Type() != mtype.TimeType {
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
	if !dst.CanSet() {
		return CanSetError{Name: "PtrCopyValue"}
	}
	if dst.Kind() != reflect.Ptr {
		return LookupCopyValueError{Name: "PtrCopyValue", Kinds: []reflect.Kind{reflect.Ptr}, Received: dst}
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

func skipElem(elem reflect.Value) reflect.Value {
	for ; elem.Kind() == reflect.Ptr || elem.Kind() == reflect.Interface; elem = elem.Elem() {
	}

	return elem
}
