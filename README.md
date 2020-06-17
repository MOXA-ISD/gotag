# Gotag [![Build Status](http://icsdrone.moxa.online/api/badges/MOXA-ISD/gotag/status.svg?ref=refs/heads/develop)](http://icsdrone.moxa.online/MOXA-ISD/gotag)

GoTag is a Go pakcage for [ThingsPro](https://www.moxa.com/en/products/industrial-computing/system-software/thingspro-2). It integrates mqtt client and protobuffer which makes data exchanging become easily and narrow down the transmission bandwidth.

Installation
------------

Since `Gotag` is a golang version wrapper of libmx-dx, you will also need to install the necessary dynamic libraries from the 3rd-party directory:

#### libparson1
#### libmosquitto1
#### libmx-dx1

### Apt Install
```bash
apt-get update
apt-get install -y -f ./*.deb
```

### Docker Image
If you are a docker user, downsizing the image could be the one of tough issues.
We provide the following sample for you to easily copy the necessary libraries in a Dockerfile.
```
RUN mkdir -p /usr/include/libmx-dx
COPY --from=build-env \
		/usr/include/libmx-dx \
		/usr/include/libmx-dx

COPY --from=build-env \
		/usr/include/parson.h \
		/usr/include/mosquitto.h \
		/usr/include/

COPY --from=build-env \
        /usr/lib/arm-linux-gnueabihf/libmx* \
        /usr/lib/arm-linux-gnueabihf/libparson.so* \
        /usr/lib/arm-linux-gnueabihf/libssl.so* \
        /usr/lib/arm-linux-gnueabihf/libcrypto.so* \
        /usr/lib/arm-linux-gnueabihf/libprotobuf-c* \
        /usr/lib/arm-linux-gnueabihf/libmosquitto* \
        /usr/lib/arm-linux-gnueabihf/
```


Once you have finished the prerequisite, run the command to install `Gotag`:

```bash
    go get github.com/MOXA-ISD/gotag
```

Build a gotag client
--------------

```go
import (
    gotag github.com/MOXA-ISD/gotag
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
func Publish(client *gotag.Tagf) {
    value := gotag.NewValue(1.414)
    client.Publish(
        "moduleName"
        "sourceName",
        "tagName",
        value,
        t.TAG_VALUE_TYPE_DOUBLE,
        1546920188000)
}

```

### Subscribe data
```go
func Subscribe(client *gotag.Tagf) {
    client.SubscribeCallback(Handler)
    client.Subscribe("moduleName", "sourceName", "tagName")
}
```

### SubscribeCallback
Gotag needs to register a callback function for subscribed topics.
```go
func Handler(module, source, tag string, val *t.Value, valtype uint16, ts uint64) {
    fmt.Printf("Module: %v,", module)
    fmt.Printf("Source: %v,", source)
    fmt.Printf("Tag: %v,", tag)
    fmt.Printf("Value: %v,", val.GetDouble())
    fmt.Printf("ValueType: %v,", valtype)
    fmt.Printf("At: %v,", ts)
}
```
