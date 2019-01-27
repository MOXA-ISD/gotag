package main

import (
    "os"
    "log"
    "syscall"
    "os/signal"

    t "github.com/CPtung/gotag"
)

func Exit() chan os.Signal {
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    return quit
}

func Handler(source string, tag string, val *t.Value, valtype int32, ts uint64, unit string) {
    log.Printf("Source: %v,", source)
    log.Printf("Tag: %v,", tag)
    log.Printf("Value: %v,", val.GetDouble())
    log.Printf("ValueType: %v,", valtype)
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
    _tag.Subscribe("electricity", "voltage")

    value := t.NewValue(1.414)
    _tag.Publish("electricity", "voltage", value, t.TAG_VALUE_TYPE_DOUBLE, 1546920188000, "v")

    <-Exit()
}
