package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	t "github.com/MOXA-ISD/gotag"
)

func Exit() chan os.Signal {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	return quit
}

func Handler(module, source, tag string, val *t.Value, valtype uint16, ts uint64) {
	log.Printf("Source: %v,", source)
	log.Printf("Tag: %v,", tag)
	switch valtype {
	case t.TAG_VALUE_TYPE_BOOLEAN:
		log.Printf("ValueType: %v,", "boolean")
		log.Printf("Value: %v,", val.GetInt())
		break
	case t.TAG_VALUE_TYPE_INT64, t.TAG_VALUE_TYPE_INT8,
		t.TAG_VALUE_TYPE_INT16, t.TAG_VALUE_TYPE_INT32:
		log.Printf("ValueType: %v,", "integer")
		log.Printf("Value: %v,", val.GetInt())
		break
	case t.TAG_VALUE_TYPE_UINT64, t.TAG_VALUE_TYPE_UINT8,
		t.TAG_VALUE_TYPE_UINT16, t.TAG_VALUE_TYPE_UINT32:
		log.Printf("ValueType: %v,", "unsigned int")
		log.Printf("Value: %v,", val.GetUint())
		break
	case t.TAG_VALUE_TYPE_FLOAT:
		log.Printf("ValueType: %v,", "float32")
		log.Printf("Value: %v,", val.GetFloat())
		break
	case t.TAG_VALUE_TYPE_DOUBLE:
		log.Printf("ValueType: %v,", "float64")
		log.Printf("Value: %v,", val.GetDouble())
		break
	case t.TAG_VALUE_TYPE_STRING:
		log.Printf("ValueType: %v,", "string")
		log.Printf("Value: %v,", val.GetStr())
		break
	case t.TAG_VALUE_TYPE_BYTEARRAY:
		log.Printf("ValueType: %v,", "bytearray")
		log.Printf("Value: %v,", string(val.GetBytes()))
		break
	}
	log.Printf("At: %v,", ts)
}

func main() {
	_tag, err := t.NewClient()
	if err != nil {
		log.Println(err)
		return
	}
	defer _tag.Delete()

	time.Sleep(time.Second)

	_tag.SubscribeCallback(Handler)
	for i := 0; i < 6; i++ {
		topic := fmt.Sprintf("test%d", i)
		_tag.Subscribe("moxa", "gotag", topic)
	}

	// double value
	value := t.NewValue(1.414)
	_tag.Publish("moxa", "gotag", "test0", value, t.TAG_VALUE_TYPE_DOUBLE, uint64(t.GetTimestamp()))
	// int value
	var iTest int64 = -12345
	value = t.NewValue(iTest)
	_tag.Publish("moxa", "gotag", "test1", value, t.TAG_VALUE_TYPE_INT64, uint64(t.GetTimestamp()))
	// uint value
	var uTest uint64 = 12345
	value = t.NewValue(uTest)
	_tag.Publish("moxa", "gotag", "test2", value, t.TAG_VALUE_TYPE_UINT64, uint64(t.GetTimestamp()))
	// float value
	var fTest float32 = 1.1444
	value = t.NewValue(fTest)
	_tag.Publish("moxa", "gotag", "test3", value, t.TAG_VALUE_TYPE_FLOAT, uint64(t.GetTimestamp()))
	// bytearray value
	var bTest []byte = []byte("Thingspro")
	va := t.NewValue(bTest)
	_tag.Publish("moxa", "gotag", "test4", va, t.TAG_VALUE_TYPE_BYTEARRAY, uint64(t.GetTimestamp()))
	// string value
	var strTest string = "thingspro-gotag-test"
	value = t.NewValue(strTest)
	_tag.Publish("moxa", "gotag", "test5", value, t.TAG_VALUE_TYPE_STRING, uint64(t.GetTimestamp()))
	<-Exit()
}
