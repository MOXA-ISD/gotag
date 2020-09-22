package gotag

import (
	"log"
	"reflect"
)

const (
	ERR_SUCESS        = 0
	ERR_FAILED        = 1
	ERR_INVALID_INPUT = 2
	ERR_NULL_ACCESS   = 3
)

const (
	TAG_VALUE_TYPE_BOOLEAN   = 0
	TAG_VALUE_TYPE_INT8      = 1
	TAG_VALUE_TYPE_INT16     = 2
	TAG_VALUE_TYPE_INT32     = 3
	TAG_VALUE_TYPE_INT64     = 4
	TAG_VALUE_TYPE_UINT8     = 5
	TAG_VALUE_TYPE_UINT16    = 6
	TAG_VALUE_TYPE_UINT32    = 7
	TAG_VALUE_TYPE_UINT64    = 8
	TAG_VALUE_TYPE_FLOAT     = 9
	TAG_VALUE_TYPE_DOUBLE    = 10
	TAG_VALUE_TYPE_STRING    = 11
	TAG_VALUE_TYPE_BYTEARRAY = 12
	TAG_VALUE_TYPE_RAW       = 13
	TAG_VALUE_TYPE_INT       = 90
	TAG_VALUE_TYPE_UINT      = 91
)

type Value struct {
	bl bool
	i  int64
	u  uint64
	f  float32
	d  float64
	s  string
	b  []byte
	rp []byte
}

func (m *Value) GetFloat() float32 {
	if m != nil {
		return m.f
	}
	return 0
}

func (m *Value) SetFloat(f float32) {
	if m != nil {
		m.f = f
	}
}

func (m *Value) GetDouble() float64 {
	if m != nil {
		return m.d
	}
	return 0
}

func (m *Value) SetDouble(d float64) {
	if m != nil {
		m.d = d
	}
}

func (m *Value) GetInt() int64 {
	if m != nil {
		return m.i
	}
	return 0
}

func (m *Value) SetInt(i int64) {
	if m != nil {
		m.i = i
	}
}

func (m *Value) GetUint() uint64 {
	if m != nil {
		return m.u
	}
	return 0
}

func (m *Value) SetUint(u uint64) {
	if m != nil {
		m.u = u
	}
}

func (m *Value) GetStr() string {
	if m != nil {
		return m.s
	}
	return ""
}

func (m *Value) SetStr(s string) {
	if m != nil {
		m.s = s
	}
}

func (m *Value) GetBytes() []byte {
	if m != nil {
		return m.b
	}
	return nil
}

func (m *Value) SetBytes(b []byte) {
	if m != nil {
		m.b = b
	}
}

func (m *Value) GetRaw() []byte {
	if m != nil {
		return m.rp
	}
	return nil
}

func (m *Value) SetRaw(rp []byte) {
	if m != nil {
		m.rp = rp
	}
}

func NewValue(value interface{}) *Value {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Bool:
		return &Value{bl: value.(bool)}
	case reflect.Int:
		return &Value{i: int64(value.(int))}
	case reflect.Int8:
		return &Value{i: int64(value.(int8))}
	case reflect.Int16:
		return &Value{i: int64(value.(int16))}
	case reflect.Int32:
		return &Value{i: int64(value.(int32))}
	case reflect.Int64:
		return &Value{i: value.(int64)}
	case reflect.Uint:
		return &Value{u: uint64(value.(uint))}
	case reflect.Uint8:
		return &Value{u: uint64(value.(uint8))}
	case reflect.Uint16:
		return &Value{u: uint64(value.(uint16))}
	case reflect.Uint32:
		return &Value{u: uint64(value.(uint32))}
	case reflect.Uint64:
		return &Value{u: value.(uint64)}
	case reflect.Float32:
		return &Value{f: value.(float32)}
	case reflect.Float64:
		return &Value{d: value.(float64)}
	case reflect.String:
		return &Value{s: value.(string)}
	case reflect.Array, reflect.Slice:
		return &Value{b: value.([]byte), rp: value.([]byte)}
	default:
		log.Printf("kind: %v\n", reflect.TypeOf(value).Kind())
	}
	return &Value{}
}

type Tag struct {
	SourceName string
	TagName    string
	Val        *Value
	ValType    int32
	Ts         uint64
	Unit       string
}
