package gotag

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lmx-dx
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <libmx-dx/dx_api.h>

extern void dxSubCallback(DX_TAG_OBJ*, uint16_t, void*);

void dx_tag_proxy_callback(DX_TAG_OBJ* dx_data_obj, uint16_t obj_cnt, void* user_data) {
	dxSubCallback(dx_data_obj, obj_cnt, user_data);
}

void to_int_value(int64_t *i, DX_TAG_VALUE *val, uint16_t _size, int to_dx) {
	if (to_dx) {
		switch (_size) {
		case DX_TAG_VALUE_TYPE_INT8:
			val->i8 = (int8_t)*i;
			break;
		case DX_TAG_VALUE_TYPE_INT16:
			val->i16 = (int16_t)*i;
			break;
		case DX_TAG_VALUE_TYPE_INT32:
			val->i32 = (int32_t)*i;
			break;
		case DX_TAG_VALUE_TYPE_INT64:
			val->i64 = (int64_t)*i;
			break;
		default:
			val->i = *i;
		}
	}
	else {
		switch (_size) {
		case DX_TAG_VALUE_TYPE_INT8:
			*i = (int64_t)val->i8;
			break;
		case DX_TAG_VALUE_TYPE_INT16:
			*i = (int64_t)val->i16;
			break;
		case DX_TAG_VALUE_TYPE_INT32:
			*i = (int64_t)val->i32;
			break;
		case DX_TAG_VALUE_TYPE_INT64:
			*i = (int64_t)val->i64;
			break;
		default:
			*i = val->i;
		}
	}
}

void to_uint_value(uint64_t *u, DX_TAG_VALUE *val, uint16_t _size, int to_dx) {
	if (to_dx) {
		switch (_size) {
		case DX_TAG_VALUE_TYPE_UINT8:
			val->u8 = (uint8_t)*u;
			break;
		case DX_TAG_VALUE_TYPE_UINT16:
			val->u16 = (uint16_t)*u;
			break;
		case DX_TAG_VALUE_TYPE_UINT32:
			val->u32 = (uint32_t)*u;
			break;
		case DX_TAG_VALUE_TYPE_UINT64:
			val->u64 = (uint64_t)*u;
			break;
		default:
			val->u = *u;
		}
	}
	else {
		switch (_size) {
		case DX_TAG_VALUE_TYPE_UINT8:
			*u = (uint64_t)val->i8;
			break;
		case DX_TAG_VALUE_TYPE_UINT16:
			*u = (uint64_t)val->u16;
			break;
		case DX_TAG_VALUE_TYPE_UINT32:
			*u = (uint64_t)val->u32;
			break;
		case DX_TAG_VALUE_TYPE_UINT64:
			*u = (uint64_t)val->u64;
			break;
		default:
			*u = val->u;
		}
	}
}

void to_float_value(float *f, DX_TAG_VALUE *val, int to_dx) {
	if (to_dx)
		val->f = *f;
	else
		*f = val->f;
}

void to_double_value(double *d, DX_TAG_VALUE *val, int to_dx) {
	if (to_dx)
		val->d = *d;
	else
		*d = val->d;
}

void to_str_value(char **str, DX_TAG_VALUE *val, int to_dx) {
	if (to_dx) {
		val->sl = strlen(*str) + 1;
		val->s = calloc(val->sl, sizeof(char));
		strcpy(val->s, *str);
	}
	else {
		if (val->sl > 0 && val->s != NULL) {
			*str = realloc(*str, (val->sl + 1) * sizeof(char));
			strcpy(*str, val->s);
		}
	}
}

int to_bytearray_value(uint8_t **b, int len, DX_TAG_VALUE *val, int to_dx) {
	if (to_dx) {
		if (*b != NULL && len > 0) {
			val->b = malloc(len * sizeof(uint8_t));
			memcpy(val->b, *b, len * sizeof(uint8_t));
			val->bl = len;
		}
		return len;
	}
	else {
		*b = realloc(*b, val->bl * sizeof(uint8_t));
		memcpy(*b, val->b, val->bl * sizeof(uint8_t));
		return val->bl;
	}
}

int to_raw_value(uint8_t **r, int len, DX_TAG_VALUE *val, int to_dx) {
	if (to_dx) {
		if (*r != NULL && len > 0) {
			val->rp = malloc(len * sizeof(uint8_t));
			memcpy(val->rp, *r, len * sizeof(uint8_t));
			val->rl = len;
		}
		return len;
	}
	else {
		*r = realloc(*r, val->rl * sizeof(uint8_t));
		memcpy(*r, val->rp, val->rl * sizeof(uint8_t));
		return val->rl;
	}
}

void free_alloc(DX_TAG_VALUE *val, uint16_t val_type) {
	if (val_type == DX_TAG_VALUE_TYPE_BYTEARRAY && val->b) {
		free(val->b);
	}
	else if (val_type == DX_TAG_VALUE_TYPE_RAW && val->rp) {
		free(val->rp);
	}
	else if (val_type == DX_TAG_VALUE_TYPE_STRING && val->s) {
	       free(val->s);
	}
}

*/
import "C"
import (
	"strings"
	"time"
	"unsafe"

	"github.com/mattn/go-pointer"
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

func EncodeDxValue(val *Value, v *C.DX_TAG_VALUE, valType uint16) {
	switch valType {
	case C.DX_TAG_VALUE_TYPE_BOOLEAN:
		fallthrough
	case C.DX_TAG_VALUE_TYPE_INT, C.DX_TAG_VALUE_TYPE_INT8, C.DX_TAG_VALUE_TYPE_INT16,
		C.DX_TAG_VALUE_TYPE_INT32, C.DX_TAG_VALUE_TYPE_INT64:
		C.to_int_value((*C.int64_t)(unsafe.Pointer(&val.i)), v, C.uint16_t(valType), 1)
	case C.DX_TAG_VALUE_TYPE_UINT, C.DX_TAG_VALUE_TYPE_UINT8, C.DX_TAG_VALUE_TYPE_UINT16,
		C.DX_TAG_VALUE_TYPE_UINT32, C.DX_TAG_VALUE_TYPE_UINT64:
		C.to_uint_value((*C.uint64_t)(unsafe.Pointer(&val.u)), v, C.uint16_t(valType), 1)
	case C.DX_TAG_VALUE_TYPE_FLOAT:
		C.to_float_value((*C.float)(unsafe.Pointer(&val.f)), v, 1)
	case C.DX_TAG_VALUE_TYPE_DOUBLE:
		C.to_double_value((*C.double)(unsafe.Pointer(&val.d)), v, 1)
	case C.DX_TAG_VALUE_TYPE_STRING:
		cstr := C.CString(val.s)
		defer C.free(unsafe.Pointer(cstr))
		C.to_str_value((**C.char)(unsafe.Pointer(&cstr)), v, 1)
	case C.DX_TAG_VALUE_TYPE_BYTEARRAY:
		ucstr := (*C.uint8_t)(C.CBytes(val.b))
		defer C.free(unsafe.Pointer(ucstr))
		C.to_bytearray_value((**C.uint8_t)(unsafe.Pointer(&ucstr)), C.int(len(val.b)), v, 1)
	case C.DX_TAG_VALUE_TYPE_RAW:
		ucstr := (*C.uint8_t)(C.CBytes(val.rp))
		defer C.free(unsafe.Pointer(ucstr))
		C.to_raw_value((**C.uint8_t)(unsafe.Pointer(&ucstr)), C.int(len(val.rp)), v, 1)
	}
}

func DecodeDxValue(val *Value, v *C.DX_TAG_VALUE, valType uint16) {
	switch valType {
	case C.DX_TAG_VALUE_TYPE_BOOLEAN:
		fallthrough
	case C.DX_TAG_VALUE_TYPE_INT, C.DX_TAG_VALUE_TYPE_INT8, C.DX_TAG_VALUE_TYPE_INT16,
		C.DX_TAG_VALUE_TYPE_INT32, C.DX_TAG_VALUE_TYPE_INT64:
		C.to_int_value((*C.int64_t)(unsafe.Pointer(&val.i)), v, C.uint16_t(valType), 0)
	case C.DX_TAG_VALUE_TYPE_UINT, C.DX_TAG_VALUE_TYPE_UINT8, C.DX_TAG_VALUE_TYPE_UINT16,
		C.DX_TAG_VALUE_TYPE_UINT32, C.DX_TAG_VALUE_TYPE_UINT64:
		C.to_uint_value((*C.uint64_t)(unsafe.Pointer(&val.u)), v, C.uint16_t(valType), 0)
	case C.DX_TAG_VALUE_TYPE_FLOAT:
		C.to_float_value((*C.float)(unsafe.Pointer(&val.f)), v, 0)
	case C.DX_TAG_VALUE_TYPE_DOUBLE:
		C.to_double_value((*C.double)(unsafe.Pointer(&val.d)), v, 0)
	case C.DX_TAG_VALUE_TYPE_STRING:
		cstr := (*C.char)(C.malloc(C.sizeof_char))
		C.to_str_value((**C.char)(unsafe.Pointer(&cstr)), v, 0)
		if cstr != nil {
			val.s = C.GoString(cstr)
			defer C.free(unsafe.Pointer(cstr))
		}
	case C.DX_TAG_VALUE_TYPE_RAW:
		fallthrough
	case C.DX_TAG_VALUE_TYPE_BYTEARRAY:
		ucstr := (*C.uint8_t)(C.malloc(C.sizeof_uint8_t))
		bsize := C.to_bytearray_value((**C.uint8_t)(unsafe.Pointer(&ucstr)), 0, v, 0)
		if ucstr != nil {
			val.b = C.GoBytes(unsafe.Pointer(ucstr), bsize)
			defer C.free(unsafe.Pointer(ucstr))
		}
	}
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func FreeAlloc(val *C.DX_TAG_VALUE, valType uint16) {
	C.free_alloc(val, C.uint16_t(valType))
}

func NewDataExchange() *DataExchange {
	dx := DataExchange{}
	if dx.c = C.dx_tag_client_init(nil,
		(*[0]byte)(unsafe.Pointer(C.dx_tag_proxy_callback))); dx.c == nil {
		return nil
	}
	ptr := pointer.Save(&dx)
	C.dx_tag_set_user_obj(dx.c, ptr)
	return &dx
}
