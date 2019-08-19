package gotag_test

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"

    gotag "github.com/MOXA-ISD/gotag"
)

var (
    sourceName  string
    tagName	string
    dataUnit    string
    timestamp   uint64
    dataType    int32
    retValue    *gotag.Value
)

func Handler(source string, tag string, val *gotag.Value, valtype int32, ts uint64, unit string) {
    sourceName = source
    tagName = tag
    dataType = valtype
    dataUnit = unit
    timestamp = ts
    retValue = val
}

func Test_GoTag_Create(t *testing.T) {
    _tag, err := gotag.NewClient("")
    assert.Equal(t, err, nil)
    defer _tag.Delete()
}

/* BOOLEAN */
func Test_GoTag_Publish_Boolean(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_BOOLEAN
	dUnit string = ""
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var blTest bool = false
    value := gotag.NewValue(blTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetInt(), retValue.GetInt())
}

/* INT */
func Test_GoTag_Publish_Int(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_INT
	dUnit string = ""
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var iTest int64 = -1234567890
    value := gotag.NewValue(iTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetInt(), retValue.GetInt())
}

/* UINT */
func Test_GoTag_Publish_Uint(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_UINT
	dUnit string = "v"
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var uTest uint64 = 1234567890
    value := gotag.NewValue(uTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetUint(), retValue.GetUint())
}


/* Float */
func Test_GoTag_Publish_Float(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_FLOAT
	dUnit string = "v"
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var fTest float32 = 1.1444
    value := gotag.NewValue(fTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetFloat(), retValue.GetFloat())
}

/* String */
func Test_GoTag_Publish_String(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_STRING
	dUnit string = "v"
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var strTest string = "thingspro-gotag-test"
    value := gotag.NewValue(strTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetStr(), retValue.GetStr())
}

/* Bytes */
func Test_GoTag_Publish_Bytes(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_BYTEARRAY
	dUnit string = "v"
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var bTest []byte = []byte("Thingspro")
    value := gotag.NewValue(bTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetBytes(), retValue.GetBytes())
}

/* Double */
func Test_GoTag_Publish_Double(t *testing.T) {
    var (
        source string = "gotag"
        tag string = "test"
	ts uint64 = 1563760922000
	dType int32 = gotag.TAG_VALUE_TYPE_DOUBLE
	dUnit string = "v"
    )

    _tag, err := gotag.NewClient("")
    if err != nil {
       assert.Fail(t, "Testing Publish Failed")
    }
    defer _tag.Delete()

    _tag.SubscribeCallback(Handler)
    _tag.Subscribe("gotag", "test")

    var dTest float64 = 123.033321
    value := gotag.NewValue(dTest)
    _tag.Publish(source, tag, value, dType, ts, dUnit)
    time.Sleep(1 * time.Second)

    assert.Equal(t, sourceName, source)
    assert.Equal(t, tagName, tag)
    assert.Equal(t, timestamp, ts)
    assert.Equal(t, dataType, dType)
    assert.Equal(t, dataUnit, dUnit)
    assert.Equal(t, value.GetDouble(), retValue.GetDouble())
}
