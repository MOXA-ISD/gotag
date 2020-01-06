package main

import (
    "os"
    "log"
    "syscall"
    "os/signal"

    t "github.com/MOXA-ISD/gotag"
)

func Exit() chan os.Signal {
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    return quit
}

func Handler(source string, tag string, val *t.Value, valtype int32, ts uint64, unit string) {
    log.Printf("Source: %v,", source)
    log.Printf("Tag: %v,", tag)
    switch (valtype) {
        case t.TAG_VALUE_TYPE_BOOLEAN:
            log.Printf("ValueType: %v,", "boolean")
            log.Printf("Value: %v,", val.GetInt())
            break
        case t.TAG_VALUE_TYPE_INT, t.TAG_VALUE_TYPE_INT8,
                 t.TAG_VALUE_TYPE_INT16, t.TAG_VALUE_TYPE_INT32:
            log.Printf("ValueType: %v,", "integer")
            log.Printf("Value: %v,", val.GetInt())
            break
        case t.TAG_VALUE_TYPE_UINT, t.TAG_VALUE_TYPE_UINT8,
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
    log.Printf("Unit: %v\n", unit)
}

func main() {
    _tag, err := t.NewClient()
    if err != nil {
        log.Println(err)
        return
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")
/*
    value := t.NewValue(1.414)
    _tag.Publish("gotag", "test", value, t.TAG_VALUE_TYPE_DOUBLE, 1546920188000, "v")
    var iTest int64 = -12345
    value = t.NewValue(iTest)
    _tag.Publish("gotag", "test", value, t.TAG_VALUE_TYPE_INT, 1546920188000, "v")
    var uTest uint64 = 12345
    value = t.NewValue(uTest)
    _tag.Publish("gotag", "test", value, t.TAG_VALUE_TYPE_UINT, 1546920188000, "v")
    var fTest float32 = 1.1444
    value = t.NewValue(fTest)
    _tag.Publish("gotag", "test", value, t.TAG_VALUE_TYPE_FLOAT, 1546920188000, "v")
    var strTest string = "thingspro-gotag-test"
    value = t.NewValue(strTest)
    _tag.Publish("gotag", "test", value, t.TAG_VALUE_TYPE_STRING, 1546920188000, "v")
    var bTest []byte = []byte("Thingspro")
    value = t.NewValue(bTest)
    _tag.Publish("gotag", "test", value, t.TAG_VALUE_TYPE_BYTEARRAY, 1546920188000, "v")*/

	log.Println("get tag value")
	t := _tag.Get("modbus", "ioLogik", "di0")
	log.Printf("value: %v\n", t)

    <-Exit()
}
