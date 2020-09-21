package gotag

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lmx-dx
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <dx.h>
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/mattn/go-pointer"
)

//export dxSubCallback
func dxSubCallback(topic *C.char, value unsafe.Pointer, valueType C.dx_tag_value_type, valueSize C.uint32_t, timestamp C.uint64_t, userData unsafe.Pointer) {
	var (
		val              Value
		module, src, tag string
	)
	dx := pointer.Restore(userData).(*DataExchange)
	if dx.ontag != nil {
		val = Value{}
		DecodeDxValue(&val, value, valueSize, uint16(valueType))
		module, src, tag = DecodeTopic(C.GoString(topic))
		dx.ontag(module, src, tag, &val, uint16(valueType), uint64(timestamp))
	}
}

func (d *DataExchange) Publish(topic string, valType uint16, val *Value, ts uint64) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}
	// combile module, source, tag names to the publish topic
	cstrTopic := C.CString(topic)
	defer C.free(unsafe.Pointer(cstrTopic))
	// convert go value to dx value
	dxVal, size := EncodeDxValue(val, valType)

	C.dx_tag_instant_write(
		d.c,
		cstrTopic,
		dxVal,
		C.dx_tag_value_type(valType),
		C.uint32_t(size),
		C.uint64_t(ts))
	return nil
}

func (d *DataExchange) Subscribe(topic string) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}

	if _, ok := d.topics[topic]; ok {
		return nil
	}

	// combile module, source, tag names to the publish topic
	cstrTopic := C.CString(topic)
	defer C.free(unsafe.Pointer(cstrTopic))

	// subscribe topic
	tag := C.dx_tag_consumer_new(
		d.c,
		cstrTopic)

	// add subscribed topic
	d.topics[topic] = tag
	return nil
}

func (d *DataExchange) SubscribeCallback(hnd OnTagCallback) error {
	if d.c == nil {
		return errors.New("tag client not found")
	}
	d.ontag = hnd
	return nil
}

func (d *DataExchange) Close() error {
	// free all tags
	for _, tag := range d.topics {
		C.dx_tag_destroy(d.c, tag)
	}

	if d.c != nil {
		C.dx_destroy(d.c)
	}
	return nil
}

type DataExchange struct {
	c      *C.dx_t
	name   string
	topics map[string]*C.dx_tag_t
	ontag  OnTagCallback
}
