# Gotag [![CircleCI](https://circleci.com/gh/CPtung/gotag/tree/master.svg?style=shield)](https://circleci.com/gh/CPtung/gotag/tree/master)

GoTag is a Go pakcage for [ThingsPro](https://www.moxa.com/en/products/industrial-computing/system-software/thingspro-2). It integrates the mqtt client and the protobuffer which makes data exchanging become easily and narrow down the transmission bandwidth.

Installation
------------

Once you have installed Go, run these commands to install the `gotag`:

```bash
    go get github.com/CPtung/gotag
```

Build a gotag client
--------------

```go
import (
    gotag github.com/CPtung/gotag
)

func main() {   
    client, err := gotag.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Delete()

    //...
}

```

Run data pub/sub
---------------

As mentioned, gotag use a mqtt client to do the transmission which means it has to pub/sub topics to send/receive data.

### Publish data
```go
func PUB(client *gotag.Tagf) {
    value := gotag.NewValue(1.414)
    client.Publish(
        "gotag",
        "test",
        value,
        t.TAG_VALUE_TYPE_DOUBLE,
        1546920188000,
        "unit")
}

```

### Subscribe data
```go
func SUB(client *gotag.Tagf) {
    client.SubscribeCallback(Handler)
    client.Subscribe("gotag", "test")
}

```

### SubscribeCallback
Gotag needs to register a callback function for subscribed topics.
```go
func Handler(source string, tag string, val *t.Value, valtype int32, ts uint64, unit string) {
    fmt.Printf("Source: %v,", source)
    fmt.Printf("Tag: %v,", tag)
    fmt.Printf("Value: %v,", val.GetDouble())
    fmt.Printf("ValueType: %v,", valtype)
    fmt.Printf("At: %v,", ts)
    fmt.Printf("Unit: %v\n", unit)
}
```
