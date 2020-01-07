# Gotag [![Build Status](http://icsdrone.moxa.online/api/badges/MOXA-ISD/gotag/status.svg?ref=refs/heads/chore/deb)](http://icsdrone.moxa.online/MOXA-ISD/gotag)

GoTag is a Go pakcage for [ThingsPro](https://www.moxa.com/en/products/industrial-computing/system-software/thingspro-2). It integrates mqtt client and protobuffer which makes data exchanging become easily and narrow down the transmission bandwidth.

Installation
------------

Since `Gotag` is a golang version wrapper of libmx-dx, you will also need to install the necessary dynamic libraries from the following download links:

#### libparson1
[[armhf](https://moxaics.s3-ap-northeast-1.amazonaws.com/debian/all/libparson1_1.1.0-1_armhf.deb)] [[amd64](https://moxaics.s3-ap-northeast-1.amazonaws.com/debian/all/libparson1_1.1.0-1_amd64.deb)]

#### libmosquitto1
[[armhf](https://moxaics.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mosquitto/feat/support-unixsocket/16/libmosquitto1_1.6.8-1%2Bun1_armhf.deb)] [[amd64](https://moxaics.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mosquitto/feat/support-unixsocket/16/libmosquitto1_1.6.8-1%2Bun1_amd64.deb)]

#### libmx-dx1
[[armhf](https://moxaics.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/edge-dx-engine/refactor/dx-unix/72/build-armhf/libmx-dx1_0.12.2-1_armhf.deb)] [[amd64](https://moxaics.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/edge-dx-engine/refactor/dx-unix/72/build-amd64/libmx-dx1_0.12.2-1_amd64.deb)]


### Apt Install
```bash
apt-get update
apt-get install -y -f ./*.deb
```

### Docker Image
If you are a docker user, downsizing the image could be the one of tough issues.
We provide the following sample for you to only copy the necessary libraries.
```yaml
COPY --from=build-env \
        /usr/lib/arm-linux-gnueabihf/libmx* \
        /usr/lib/arm-linux-gnueabihf/libparson.so* \
        /usr/lib/arm-linux-gnueabihf/libssl.so* \
        /usr/lib/arm-linux-gnueabihf/libcrypto.so* \
        /usr/lib/arm-linux-gnueabihf/libprotobuf-c* \
        /usr/lib/arm-linux-gnueabihf/libcares.so* \
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
    fmt.Printf("Source: %v,", source)
    fmt.Printf("Tag: %v,", tag)
    fmt.Printf("Value: %v,", val.GetDouble())
    fmt.Printf("ValueType: %v,", valtype)
    fmt.Printf("At: %v,", ts)
}
```
