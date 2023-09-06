package mapper

import (
	"fmt"
	"reflect"
	"strings"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/24
    @desc:

***************************/

// CopyValueError  copy element does not match destination element
type CopyValueError struct {
	Name     string
	Types    []reflect.Type
	Kinds    []reflect.Kind
	Received reflect.Value
}

func (vde CopyValueError) Error() string {
	typeKinds := make([]string, 0, len(vde.Types)+len(vde.Kinds))
	for _, t := range vde.Types {
		typeKinds = append(typeKinds, t.String())
	}
	for _, k := range vde.Kinds {
		if k == reflect.Map {
			typeKinds = append(typeKinds, "map[*]*")
			continue
		}
		typeKinds = append(typeKinds, k.String())
	}
	received := vde.Received.Kind().String()
	if vde.Received.IsValid() {
		received = vde.Received.Type().String()
	}
	return fmt.Sprintf("%s can only copy valid and settable %s, but got %s", vde.Name, strings.Join(typeKinds, ", "), received)
}

type CanSetError struct {
	Name string
}

func (cse CanSetError) Error() string {
	return fmt.Sprintf("%s copy to object must be pointers", cse.Name)
}

// LookupCopyValueError The current received element does not match the function
type LookupCopyValueError struct {
	Name     string
	Types    []reflect.Type
	Kinds    []reflect.Kind
	Received reflect.Value
}

func (l LookupCopyValueError) Error() string {
	typeKinds := make([]string, 0, len(l.Types)+len(l.Kinds))
	for _, t := range l.Types {
		typeKinds = append(typeKinds, t.String())
	}
	for _, k := range l.Kinds {
		if k == reflect.Map {
			typeKinds = append(typeKinds, "map[*]*")
			continue
		}
		typeKinds = append(typeKinds, k.String())
	}
	received := l.Received.Kind().String()
	if l.Received.IsValid() {
		received = l.Received.Type().String()
	}
	return fmt.Sprintf("%s lookup copy function error, function can only copy valid and settable %s, but got %s", l.Name, strings.Join(typeKinds, ", "), received)
}
