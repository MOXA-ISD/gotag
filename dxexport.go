package gotag
/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lmx-dx
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <libmx-dx/dx_api.h>
*/
import "C"
import "unsafe"

import (
	"errors"
	"github.com/mattn/go-pointer"
)

//export dxSubCallback
func dxSubCallback(dx_tag_obj *C.DX_TAG_OBJ, obj_cnt C.uint16_t, user_data unsafe.Pointer) {
	var (
		val Value = Value{}
		err error
		module, src, tag string
		i, numsObj uint16
		ptr uintptr
		obj *C.DX_TAG_OBJ
	)
	dx := pointer.Restore(user_data).(*DataExchange)
	if dx.ontag != nil && obj_cnt >= 0 {
		numsObj = uint16(obj_cnt)
		ptr = uintptr(unsafe.Pointer(dx_tag_obj))
		for i = 0; i < numsObj; i++ {
			obj = (*C.DX_TAG_OBJ)(unsafe.Pointer(ptr))
			DecodeDxValue(&val, &obj.val, uint16(obj.val_type))
			if module, src, tag, err = DecodeTopic(C.GoString(obj.tag)); err != nil {
				continue
			}
			dx.ontag(module, src, tag, &val, uint16(obj.val_type), uint64(obj.timestamp))
			ptr = (uintptr)(unsafe.Pointer(ptr)) + (uintptr)(C.sizeof_DX_TAG_OBJ * C.int(i))
		}
	}
}

func(d *DataExchange)Publish(topic string, valType uint16, val *Value, ts uint64) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}
	// combile module, source, tag names to the publish topic
	cstrTopic := C.CString(topic)
	defer C.free(unsafe.Pointer(cstrTopic))
	// convert go value to dx value
	dxVal := C.DX_TAG_VALUE{}
	defer FreeAlloc(&dxVal, valType)

	EncodeDxValue(val, &dxVal, valType)
	C.dx_tag_pub(
		 d.c,
		 cstrTopic,
		 C.uint16_t(C.strlen(cstrTopic)),
		 C.DX_TAG_VALUE_TYPE(valType),
		 dxVal,
		 C.uint64_t(ts))
	return nil
}

func(d *DataExchange)Subscribe(topic string) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}
	// combile module, source, tag names to the publish topic
	cstrTopic := C.CString(topic)
	defer C.free(unsafe.Pointer(cstrTopic))
	// subscribe topic
	C.dx_tag_sub(
		d.c,
		cstrTopic,
		C.uint16_t(C.strlen(cstrTopic)),
		nil)
	// add subscribed topic
	for i := range d.topics {
		if d.topics[i] == topic {
			return nil
		}
	}
	d.topics = append(d.topics, topic)
	return nil
}

func(d *DataExchange)UnSubscribe(topic string) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}
	// remove unsubscribed topic
	for i := range d.topics {
		if d.topics[i] == topic {
			d.topics = append(d.topics[:i], d.topics[i+1:]...)
			break
		}
	}
	return nil
}

func(d *DataExchange)SubscribeCallback(hnd OnTagCallback) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}
	d.ontag = hnd
	return nil
}

func(d *DataExchange)Close() error {
	if d.c != nil {
	C.dx_tag_destroy(d.c)
	}
	return nil
}

type DataExchange struct {
	c	  *C.DX_TAG_CLIENT
	topics  []string
	ontag   OnTagCallback
}
