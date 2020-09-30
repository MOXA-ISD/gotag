package gotag_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	gotag "github.com/MOXA-ISD/gotag"
)

var (
	moduleName string
	sourceName string
	tagName    string
	timestamp  uint64
	dataType   uint16
	retValue   *gotag.Value
	subscriber *gotag.Tagf
	publisher  *gotag.Tagf
)

func Handler(module string, source string, tag string, val *gotag.Value, valtype uint16, ts uint64) {
	moduleName = module
	sourceName = source
	tagName = tag
	dataType = valtype
	timestamp = ts
	retValue = val
}

func Test_GoTag_Create(t *testing.T) {
	var (
		err    error
		module string = "moxa-dx"
		source string = "gotag"
	)

	subscriber, err = gotag.NewClient()
	assert.Equal(t, err, nil)

	subscriber.SubscribeCallback(Handler)
	subscriber.Subscribe(module, source, "boolean")
	subscriber.Subscribe(module, source, "int8")
	subscriber.Subscribe(module, source, "int16")
	subscriber.Subscribe(module, source, "int32")
	subscriber.Subscribe(module, source, "int64")
	subscriber.Subscribe(module, source, "uint16")
	subscriber.Subscribe(module, source, "uint64")
	subscriber.Subscribe(module, source, "float")
	subscriber.Subscribe(module, source, "double")
	subscriber.Subscribe(module, source, "string")
	subscriber.Subscribe(module, source, "bytes")

	publisher, err = gotag.NewClient()
	assert.Equal(t, err, nil)

	time.Sleep(time.Second)
}

func Test_GoTag_Publish_Boolean(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "boolean"
		ts     uint64 = uint64(gotag.GetTimestamp())
		dType  uint16 = gotag.TAG_VALUE_TYPE_BOOLEAN
	)

	var blTest bool = true
	value := gotag.NewValue(blTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, int64(1), retValue.GetInt())
}

func Test_GoTag_Publish_Int8(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "int8"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_INT8
	)

	var iTest int8 = -1
	value := gotag.NewValue(iTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, int8(value.GetInt()), int8(retValue.GetInt()))
}

func Test_GoTag_Publish_Int16(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "int16"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_INT16
	)

	var iTest int16 = -1
	value := gotag.NewValue(iTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, int16(value.GetInt()), int16(retValue.GetInt()))
}

func Test_GoTag_Publish_Int32(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "int32"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_INT32
	)

	var iTest int32 = -1
	value := gotag.NewValue(iTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, int32(value.GetInt()), int32(retValue.GetInt()))
}

func Test_GoTag_Publish_Int64(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "int64"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_INT64
	)

	var iTest int64 = -1234567890
	value := gotag.NewValue(iTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetInt(), retValue.GetInt())
}

func Test_GoTag_Publish_Uint16(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "uint16"
		ts     uint64 = uint64(gotag.GetTimestamp())
		dType  uint16 = gotag.TAG_VALUE_TYPE_UINT16
	)

	var uTest uint16 = 65535
	value := gotag.NewValue(uTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetUint(), retValue.GetUint())
}

func Test_GoTag_Publish_Uint64(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "uint64"
		ts     uint64 = uint64(gotag.GetTimestamp())
		dType  uint16 = gotag.TAG_VALUE_TYPE_UINT64
	)

	var uTest uint64 = 18446744073709551615
	value := gotag.NewValue(uTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetUint(), retValue.GetUint())
}

func Test_GoTag_Publish_Float(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "float"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_FLOAT
	)

	var fTest float32 = 1.0999999
	value := gotag.NewValue(fTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, source, sourceName)
	assert.Equal(t, tag, tagName)
	assert.Equal(t, ts, timestamp)
	assert.Equal(t, dType, dataType)
	assert.Equal(t, value.GetFloat(), retValue.GetFloat())
}

func Test_GoTag_Publish_Bytes(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "bytes"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_BYTEARRAY
	)

	var bTest []byte = []byte("Thingspro")
	value := gotag.NewValue(bTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, source, sourceName)
	assert.Equal(t, tag, tagName)
	assert.Equal(t, ts, timestamp)
	assert.Equal(t, dType, dataType)
	assert.Equal(t, value.GetBytes(), retValue.GetBytes())
}

func Test_GoTag_Publish_String(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "string"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_STRING
	)

	var strTest string = "thingspro-gotag-test"
	value := gotag.NewValue(strTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, source, sourceName)
	assert.Equal(t, tag, tagName)
	assert.Equal(t, ts, timestamp)
	assert.Equal(t, dType, dataType)
	assert.Equal(t, value.GetStr(), retValue.GetStr())
}

func Test_GoTag_Publish_Double(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag    string = "double"
		ts     uint64 = uint64(uint64(gotag.GetTimestamp()))
		dType  uint16 = gotag.TAG_VALUE_TYPE_DOUBLE
	)

	var dTest float64 = 123.099999999999999
	value := gotag.NewValue(dTest)
	publisher.Publish(module, source, tag, value, dType, ts)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetDouble(), retValue.GetDouble())
}

func Test_GoTag_Destroy(t *testing.T) {
	subscriber.Delete()
	publisher.Delete()
}
