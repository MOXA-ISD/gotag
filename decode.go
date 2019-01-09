package gotag

import (
	"fmt"
	"log"
	"regexp"
	"errors"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/CPtung/gotag/protobuf"
)

func Split(str string, symbol string) (string, string, error) {
	s := strings.Split(str, symbol)
	if len(s) < 2 {
		return str, "", errors.New("no matched")
	}
	return s[0], s[1], nil
}

func DecodeTopic(topic string) (string, string) {
    re, _ := regexp.Compile("/equs/(.*)/tags(.*)")
	str := re.ReplaceAllString(topic, "$1$2")
	source, tag, err := Split(str, "/")
	if err != nil {
		return "", ""
	}
	return source, tag
}

func getDecodeValue(v *mxtag_pb.Value, t int32) *Value{
	value := &Value{}
	if v.IntValue != nil && t == TAG_VALUE_TYPE_BOOLEAN {
		value.i = v.GetIntValue()
	} else if v.IntValue != nil && t == TAG_VALUE_TYPE_INT {
		value.i = v.GetIntValue()
	} else if v.UintValue != nil && t == TAG_VALUE_TYPE_UINT {
		value.u = v.GetUintValue()
	} else if v.FloatValue != nil && t == TAG_VALUE_TYPE_FLOAT {
		value.f = v.GetFloatValue()
	} else if v.DoubleValue != nil && t == TAG_VALUE_TYPE_DOUBLE {
		value.d = v.GetDoubleValue()
	} else if v.StrValue != nil && t == TAG_VALUE_TYPE_STRING {
		value.s = v.GetStrValue()
	} else if t == TAG_VALUE_TYPE_BYTEARRAY {
		value.b = v.GetBytesValue()
	} else {
		log.Printf("not support type (%v)\n", t)
		return nil
	}
	return value
}

func DecodePayload(payload []byte, tag *Tag) error {
    data := &mxtag_pb.Tag{}
	if err := proto.Unmarshal(payload, data); err != nil {
		return errors.New(fmt.Sprintf("unmarshaling error: %v", err))
	}
	tag.sourceName = data.GetEquipment()
	tag.tagName	= data.GetTag()
	tag.val = getDecodeValue(data.GetValue(), data.GetValueType())
	tag.valType = data.GetValueType()
	tag.ts = data.GetAtMs()
	tag.unit = data.GetUnit()
	return nil
}
