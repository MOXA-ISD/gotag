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
    log.Printf("Value: %v,", val.GetDouble())
    log.Printf("ValueType: %v,", valtype)
    log.Printf("At: %v,", ts)
    log.Printf("Unit: %v\n", unit)
}

func main() {

    tag := t.NewClient()
    if tag == nil {
        log.Println("No client")
        return
    }

    tag.SubscribeCallback(Handler)
    tag.Subscribe("+", "+")

    value := t.NewValue(1.414)
    tag.Publish("electricity", "voltage", value, t.TAG_VALUE_TYPE_DOUBLE, 1546920188000, "v")
    <-Exit()
    tag.Delete()
}
