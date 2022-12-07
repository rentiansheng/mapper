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
	if !dst.CanSet() {
		return CanSetError{Name: "InterfaceCopyValue"}
	}
	if dst.Kind() != reflect.Interface {
		return CopyValueError{Name: "InterfaceCopyValue", Kinds: []reflect.Kind{reflect.Interface}, Received: dst}
	}

	src = skipElem(src)
	// 特殊处理
	if src.Kind() == reflect.Invalid {
		if dst.Kind() != reflect.Invalid {
			// 元数据没有初始化，情况dst
			dst.Set(reflect.New(dst.Type()).Elem())
		}
		return nil
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
