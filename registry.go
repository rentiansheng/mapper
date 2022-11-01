package mapper

import (
	"context"
	"github.com/rentiansheng/mapper/mtype"
	"reflect"
	"sync"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/25
    @desc:

***************************/

var (
	lock        sync.RWMutex
	typeHandler = make(map[reflect.Type]CopyValueFunc, 0)
	kindHandler = make(map[reflect.Kind]CopyValueFunc, 0)
)

type register struct {
	sync.RWMutex
	typeHandler map[reflect.Type]CopyValueFunc
	kindHandler map[reflect.Kind]CopyValueFunc
}

func newRegister() *register {
	return &register{
		typeHandler: make(map[reflect.Type]CopyValueFunc, 0),
		kindHandler: make(map[reflect.Kind]CopyValueFunc, 0),
	}
}

type CopyValueFunc func(ctx context.Context, src reflect.Value, dst reflect.Value) error

func (dcv *defaultCopyValue) Register(typ reflect.Type, fn CopyValueFunc) {
	dcv.register.Lock()
	defer dcv.register.Unlock()
	dcv.register.typeHandler[typ] = fn
}

func (dcv *defaultCopyValue) RegistryKind(kind reflect.Kind, fn CopyValueFunc) {
	dcv.register.Lock()
	defer dcv.register.Unlock()
	dcv.register.kindHandler[kind] = fn
}

func (dcv *defaultCopyValue) lookupCopyValue(value reflect.Value) (CopyValueFunc, error) {
	if fn, ok := dcv.register.typeHandler[value.Type()]; ok {
		return fn, nil
	}
	if fn, ok := dcv.register.kindHandler[value.Kind()]; ok {
		return fn, nil
	}
	return nil, ErrNoCopyValue{value.Type()}
}

func (dcv *defaultCopyValue) buildDefaultRegistry() {

	dcv.Register(mtype.Bool, dcv.BooleanCopyValue)
	dcv.Register(mtype.Int, dcv.IntCopyValue)
	dcv.Register(mtype.Int8, dcv.IntCopyValue)
	dcv.Register(mtype.Int16, dcv.IntCopyValue)
	dcv.Register(mtype.Int32, dcv.IntCopyValue)
	dcv.Register(mtype.Int64, dcv.IntCopyValue)
	dcv.Register(mtype.Uint, dcv.UintCopyValue)
	dcv.Register(mtype.Uint8, dcv.UintCopyValue)
	dcv.Register(mtype.Uint16, dcv.UintCopyValue)
	dcv.Register(mtype.Uint32, dcv.UintCopyValue)
	dcv.Register(mtype.Uint64, dcv.UintCopyValue)
	dcv.Register(mtype.Float32, dcv.FloatCopyValue)
	dcv.Register(mtype.Float64, dcv.FloatCopyValue)
	dcv.Register(mtype.String, dcv.StringCopyValue)

	dcv.RegistryKind(reflect.Bool, dcv.BooleanCopyValue)
	dcv.RegistryKind(reflect.Int, dcv.IntCopyValue)
	dcv.RegistryKind(reflect.Int8, dcv.IntCopyValue)
	dcv.RegistryKind(reflect.Int16, dcv.IntCopyValue)
	dcv.RegistryKind(reflect.Int32, dcv.IntCopyValue)
	dcv.RegistryKind(reflect.Int64, dcv.IntCopyValue)
	dcv.RegistryKind(reflect.Uint, dcv.UintCopyValue)
	dcv.RegistryKind(reflect.Uint8, dcv.UintCopyValue)
	dcv.RegistryKind(reflect.Uint16, dcv.UintCopyValue)
	dcv.RegistryKind(reflect.Uint32, dcv.UintCopyValue)
	dcv.RegistryKind(reflect.Uint64, dcv.UintCopyValue)
	dcv.RegistryKind(reflect.Float32, dcv.FloatCopyValue)
	dcv.RegistryKind(reflect.Float64, dcv.FloatCopyValue)
	dcv.RegistryKind(reflect.String, dcv.StringCopyValue)
	dcv.RegistryKind(reflect.Array, dcv.ArrayCopyValue)
	dcv.RegistryKind(reflect.Map, dcv.MapCopyValue)
	dcv.RegistryKind(reflect.Slice, dcv.SliceCopyValue)
	dcv.RegistryKind(reflect.Struct, dcv.StructCopyValue)
	dcv.RegistryKind(reflect.Ptr, dcv.PtrCopyValue)
	dcv.RegistryKind(reflect.Interface, dcv.InterfaceCopyValue)
}

var (
	Register     = defaultCopyValueHandler.Register
	RegisterKind = defaultCopyValueHandler.RegistryKind
)

// ErrNoCopyValue is returned when there wasn't an copy available for a type.
type ErrNoCopyValue struct {
	Type reflect.Type
}

func (ene ErrNoCopyValue) Error() string {
	if ene.Type == nil {
		return "no found copy value handler for <nil>"
	}
	return "no found copy value handler for " + ene.Type.String()
}
