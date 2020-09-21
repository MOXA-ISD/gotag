package gotag

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lmx-dx
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <dx.h>


extern void dxSubCallback(char *topic, void *value, dx_tag_value_type value_type, uint32_t value_size, uint64_t timestamp, void *user_data);
void *alloc_data_memory(dx_tag_value_type value_type, uint32_t value_size);

void protocol_event_callback(dx_t *dx, dx_write_info *write_info, void *dx_user_data)
{
	int ret;
	uint64_t ts;
	void *getVal;
	char *topic;
	dx_tag_t *cur_tag;
	uint32_t value_size;
	dx_tag_value_type value_type;

	DX_EACH_TAG_HEADER(dx, write_info, cur_tag)

	ret = dx_tag_get_value_info(cur_tag, &value_type, &value_size);
    if (!ret) {
		getVal = alloc_data_memory(value_type, value_size);
		topic = (char *)(dx_tag_get_topic(cur_tag));
		ret = dx_tag_read_with_timestamp(cur_tag, getVal, value_size, &ts);
		dxSubCallback(topic, getVal, value_type, value_size, ts, dx_user_data);
	}
	free(getVal); getVal = NULL;

    DX_EACH_TAG_FOOTER()
}

void *alloc_data_memory(dx_tag_value_type value_type, uint32_t value_size) {
	switch(value_type) {
	case DX_TAG_VALUE_TYPE_BOOLEAN:
		return malloc(sizeof(bool));
	case DX_TAG_VALUE_TYPE_INT8:
	case DX_TAG_VALUE_TYPE_UINT8:
		return malloc(sizeof(uint8_t));
	case DX_TAG_VALUE_TYPE_INT16:
	case DX_TAG_VALUE_TYPE_UINT16:
		return malloc(sizeof(uint16_t));
	case DX_TAG_VALUE_TYPE_INT32:
	case DX_TAG_VALUE_TYPE_UINT32:
		return malloc(sizeof(uint32_t));
	case DX_TAG_VALUE_TYPE_INT64:
	case DX_TAG_VALUE_TYPE_UINT64:
		return malloc(sizeof(uint64_t));
	case DX_TAG_VALUE_TYPE_FLOAT:
		return malloc(sizeof(float));
	case DX_TAG_VALUE_TYPE_DOUBLE:
		return malloc(sizeof(double));
	case DX_TAG_VALUE_TYPE_STRING:
		return malloc(value_size * sizeof(char));
	case DX_TAG_VALUE_TYPE_BYTEARRAY:
		return malloc(value_size * sizeof(char));
	case DX_TAG_VALUE_TYPE_RAW:
		return malloc(value_size * sizeof(char));
	default:
		return NULL;
	}
}

void to_bool_value(int64_t *i, void *val) {
	bool b;
	memcpy(&b, val, sizeof(bool));
	*i = b ? 1 : 0;
}

void to_int_value(int64_t *i, void *val, size_t size) {
	int8_t i8;
	int16_t i16;
	int32_t i32;
	switch (size) {
	case 1:
		memcpy(&i8, val, size);
		*i = (int64_t)i8;
		break;
	case 2:
		memcpy(&i16, val, size);
		*i = (int64_t)i16;
		break;
	case 4:
		memcpy(&i32, val, size);
		*i = (int64_t)i32;
		break;
	case 8:
		memcpy(i, val, size);
		break;
	}
}
void to_uint_value(uint64_t *u, void *val, size_t size) {
	memcpy(u, val, size);
}
void to_float_value(float *f, void *val) {
	memcpy(f, val, sizeof(float));
}
void to_double_value(double *d, void *val) {
	memcpy(d, val, sizeof(double));
}

char *to_str_value(char **str) {
	size_t sl;
	char *s = NULL;
	sl = strlen(*str) + 1;
	s = calloc(sl, sizeof(char));
	strcpy(s, *str);
	return s;
}

uint8_t *to_bytearray_value(uint8_t **b, int len) {
	uint8_t *bytes = malloc(len * sizeof(uint8_t));
	memcpy(bytes, *b, len * sizeof(uint8_t));
	return bytes;
}

uint8_t *to_raw_value(uint8_t **r, int len) {
	uint8_t *rp = malloc(len * sizeof(uint8_t));
	memcpy(rp, *r, len * sizeof(uint8_t));
	return rp;
}

void get_str_value(char **str, void *val, uint32_t len) {
	*str = realloc(*str, (len + 1) * sizeof(char));
	strcpy(*str, val);
}

int get_bytearray_value(uint8_t **b, void *val, uint32_t len) {
	*b = realloc(*b, len * sizeof(uint8_t));
	memcpy(*b, val, len * sizeof(uint8_t));
	return (int)len;
}

void free_alloc(void **val) {
	if (*val) {
		free(*val);
	}
}

*/
import "C"
import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/mattn/go-pointer"
	"github.com/teris-io/shortid"
)

func EncodeTopic(module, source, tag string) string {
	var topic strings.Builder
	topic.WriteString(module)
	topic.WriteString("/")
	topic.WriteString(source)
	topic.WriteString("/")
	topic.WriteString(tag)
	return topic.String()
}

func DecodeTopic(topic string) (string, string, string) {
	tokens := strings.Split(topic, "/")
	return tokens[0], tokens[1], tokens[2]
}

func EncodeDxValue(val *Value, valType uint16) (unsafe.Pointer, uint32) {
	switch valType {
	case C.DX_TAG_VALUE_TYPE_BOOLEAN:
		return (unsafe.Pointer(&val.bl)), 1
	case C.DX_TAG_VALUE_TYPE_INT8:
		return (unsafe.Pointer(&val.i)), 1
	case C.DX_TAG_VALUE_TYPE_INT16:
		return (unsafe.Pointer(&val.i)), 2
	case C.DX_TAG_VALUE_TYPE_INT32:
		return (unsafe.Pointer(&val.i)), 4
	case C.DX_TAG_VALUE_TYPE_INT64:
		return (unsafe.Pointer(&val.i)), 8
	case C.DX_TAG_VALUE_TYPE_UINT8:
		return (unsafe.Pointer(&val.u)), 1
	case C.DX_TAG_VALUE_TYPE_UINT16:
		return (unsafe.Pointer(&val.u)), 2
	case C.DX_TAG_VALUE_TYPE_UINT32:
		return (unsafe.Pointer(&val.u)), 4
	case C.DX_TAG_VALUE_TYPE_UINT64:
		return (unsafe.Pointer(&val.u)), 8
	case C.DX_TAG_VALUE_TYPE_FLOAT:
		return (unsafe.Pointer(&val.f)), 4
	case C.DX_TAG_VALUE_TYPE_DOUBLE:
		return (unsafe.Pointer(&val.d)), 8
	case C.DX_TAG_VALUE_TYPE_STRING:
		cstr := C.CString(val.s)
		defer C.free(unsafe.Pointer(cstr))
		return unsafe.Pointer(C.to_str_value((**C.char)(unsafe.Pointer(&cstr)))), uint32(len(val.s) + 1)
	case C.DX_TAG_VALUE_TYPE_BYTEARRAY:
		ucstr := (*C.uint8_t)(C.CBytes(val.b))
		defer C.free(unsafe.Pointer(ucstr))
		return unsafe.Pointer(C.to_bytearray_value((**C.uint8_t)(unsafe.Pointer(&ucstr)), C.int(len(val.b)))), uint32(len(val.b))
	case C.DX_TAG_VALUE_TYPE_RAW:
		ucstr := (*C.uint8_t)(C.CBytes(val.rp))
		defer C.free(unsafe.Pointer(ucstr))
		return unsafe.Pointer(C.to_raw_value((**C.uint8_t)(unsafe.Pointer(&ucstr)), C.int(len(val.rp)))), uint32(len(val.rp))
	default:
		fmt.Println("default type")
		return nil, 0
	}
}

func DecodeDxValue(val *Value, v unsafe.Pointer, size C.uint32_t, valType uint16) {
	switch valType {
	case C.DX_TAG_VALUE_TYPE_BOOLEAN:
		C.to_bool_value((*C.int64_t)(unsafe.Pointer(&val.i)), v)
	case C.DX_TAG_VALUE_TYPE_INT8:
		C.to_int_value((*C.int64_t)(unsafe.Pointer(&val.i)), v, 1)
	case C.DX_TAG_VALUE_TYPE_INT16:
		C.to_int_value((*C.int64_t)(unsafe.Pointer(&val.i)), v, 2)
	case C.DX_TAG_VALUE_TYPE_INT32:
		C.to_int_value((*C.int64_t)(unsafe.Pointer(&val.i)), v, 4)
	case C.DX_TAG_VALUE_TYPE_INT64:
		C.to_int_value((*C.int64_t)(unsafe.Pointer(&val.i)), v, 8)
	case C.DX_TAG_VALUE_TYPE_UINT8:
		C.to_uint_value((*C.uint64_t)(unsafe.Pointer(&val.u)), v, 1)
	case C.DX_TAG_VALUE_TYPE_UINT16:
		C.to_uint_value((*C.uint64_t)(unsafe.Pointer(&val.u)), v, 2)
	case C.DX_TAG_VALUE_TYPE_UINT32:
		C.to_uint_value((*C.uint64_t)(unsafe.Pointer(&val.u)), v, 4)
	case C.DX_TAG_VALUE_TYPE_UINT64:
		C.to_uint_value((*C.uint64_t)(unsafe.Pointer(&val.u)), v, 8)
	case C.DX_TAG_VALUE_TYPE_FLOAT:
		C.to_float_value((*C.float)(unsafe.Pointer(&val.f)), v)
	case C.DX_TAG_VALUE_TYPE_DOUBLE:
		C.to_double_value((*C.double)(unsafe.Pointer(&val.d)), v)
	case C.DX_TAG_VALUE_TYPE_STRING:
		cstr := (*C.char)(C.malloc(C.sizeof_char))
		C.get_str_value((**C.char)(unsafe.Pointer(&cstr)), v, size)
		if cstr != nil {
			val.s = C.GoString(cstr)
			defer C.free(unsafe.Pointer(cstr))
		}
	case C.DX_TAG_VALUE_TYPE_RAW:
		fallthrough
	case C.DX_TAG_VALUE_TYPE_BYTEARRAY:
		ucstr := (*C.uint8_t)(C.malloc(C.sizeof_uint8_t))
		bsize := C.get_bytearray_value((**C.uint8_t)(unsafe.Pointer(&ucstr)), v, size)
		if ucstr != nil {
			val.b = C.GoBytes(unsafe.Pointer(ucstr), bsize)
			defer C.free(unsafe.Pointer(ucstr))
		}
	}
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func FreeAlloc(val unsafe.Pointer) {
	C.free_alloc((*unsafe.Pointer)(unsafe.Pointer(&val)))
}

func NewDataExchange() *DataExchange {
	id, _ := shortid.Generate()
	dx := DataExchange{
		name:   id,
		topics: make(map[string]*C.dx_tag_t),
	}

	if dx.c = C.dx_new(C.CString(dx.name), (*[0]byte)(unsafe.Pointer(C.protocol_event_callback))); dx.c == nil {
		return nil
	}

	ptr := pointer.Save(&dx)
	C.dx_set_user_data(dx.c, ptr)
	return &dx
}
