package mapper

import (
	"context"
	"reflect"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/25
    @desc:

***************************/

func (dcv *defaultCopyValue) InterfaceCopyValue(ctx context.Context, src, dst reflect.Value) error {
	if !dst.CanSet() && dst.Kind() != reflect.Interface {
		return CopyValueError{Name: "InterfaceCopyValue", Kinds: []reflect.Kind{reflect.Interface}, Received: dst}
	}
	//
	for src.Kind() == reflect.Interface {
		src = src.Elem()
	}
	tmpVal := reflect.New(src.Type()).Elem()
	fn, err := dcv.lookupCopyValue(src)
	if err != nil {
		return err
	}
	if err := fn(ctx, src, tmpVal); err != nil {
		return err
	}

	dst.Set(tmpVal)
	return nil
}
