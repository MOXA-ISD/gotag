package gotag_test

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"

	gotag "github.com/MOXA-ISD/gotag"
)

var (
	moduleName	string
	sourceName	string
	tagName		string
	timestamp	uint64
	dataType	uint16
	retValue	*gotag.Value
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
	_tag, err := gotag.NewClient()
	assert.Equal(t, err, nil)
	defer _tag.Delete()
}

/* BOOLEAN */
func Test_GoTag_Publish_Boolean(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_BOOLEAN
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var blTest bool = false
	value := gotag.NewValue(blTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetInt(), retValue.GetInt())
}

/* INT */
func Test_GoTag_Publish_Int(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_INT
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var iTest int64 = -1234567890
	value := gotag.NewValue(iTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetInt(), retValue.GetInt())
}

/* UINT */
func Test_GoTag_Publish_Uint(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_UINT
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var uTest uint64 = 1234567890
	value := gotag.NewValue(uTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetUint(), retValue.GetUint())
}


/* Float */
func Test_GoTag_Publish_Float(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_FLOAT
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var fTest float32 = 1.0999999
	value := gotag.NewValue(fTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetFloat(), retValue.GetFloat())
}

/* String */
func Test_GoTag_Publish_String(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_STRING
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var strTest string = "thingspro-gotag-test"
	value := gotag.NewValue(strTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetStr(), retValue.GetStr())
}

/* Bytes */
func Test_GoTag_Publish_Bytes(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_BYTEARRAY
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var bTest []byte = []byte("Thingspro")
	value := gotag.NewValue(bTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetBytes(), retValue.GetBytes())
}

/* Double */
func Test_GoTag_Publish_Double(t *testing.T) {
	var (
		module string = "moxa-dx"
		source string = "gotag"
		tag string = "test"
		ts uint64 = 1563760922000
		dType uint16 = gotag.TAG_VALUE_TYPE_DOUBLE
	)

	_tag, err := gotag.NewClient()
	if err != nil {
	   assert.Fail(t, "Testing Publish Failed")
	}
	defer _tag.Delete()

	_tag.SubscribeCallback(Handler)
	_tag.Subscribe(module, source, tag)

	var dTest float64 = 123.099999999999999
	value := gotag.NewValue(dTest)
	_tag.Publish(module, source, tag, value, dType, ts)
	time.Sleep(1 * time.Second)

	assert.Equal(t, sourceName, source)
	assert.Equal(t, tagName, tag)
	assert.Equal(t, timestamp, ts)
	assert.Equal(t, dataType, dType)
	assert.Equal(t, value.GetDouble(), retValue.GetDouble())
}
