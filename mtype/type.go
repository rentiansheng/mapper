package mtype

import (
	"encoding/json"
	"reflect"
	"time"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/10/25
    @desc:

***************************/

var (
	ByteSliceType = reflect.TypeOf([]byte{})
	TimeType      = reflect.TypeOf(time.Time{})
	TimePtrType   = reflect.TypeOf(&time.Time{})
)

var Bool = reflect.TypeOf(false)
var Float32 = reflect.TypeOf(float32(0))
var Float64 = reflect.TypeOf(float64(0))
var Int = reflect.TypeOf(int(0))
var Int8 = reflect.TypeOf(int8(0))
var Int16 = reflect.TypeOf(int16(0))
var Int32 = reflect.TypeOf(int32(0))
var Int64 = reflect.TypeOf(int64(0))
var String = reflect.TypeOf("")
var Time = reflect.TypeOf(time.Time{})
var Uint = reflect.TypeOf(uint(0))
var Uint8 = reflect.TypeOf(uint8(0))
var Uint16 = reflect.TypeOf(uint16(0))
var Uint32 = reflect.TypeOf(uint32(0))
var Uint64 = reflect.TypeOf(uint64(0))
var JSONNumber = reflect.TypeOf(json.Number(""))
