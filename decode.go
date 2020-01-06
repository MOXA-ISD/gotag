package gotag

import (
	"errors"
	"fmt"
	"math"
	"regexp"

	mxtag_pb "github.com/MOXA-ISD/gotag/protobuf"
	"github.com/golang/protobuf/proto"
	logger "github.com/sirupsen/logrus"
)

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}

func DecodeTopic(topic string) (string, string, error) {
	re := regexp.MustCompile("/(\\+|[\\w-]+)")
	matches := re.FindAllStringSubmatch(topic, -1)
	if matches != nil && len(matches) > 3 {
		return matches[1][1], matches[3][1], nil
	}
	return "", "", errors.New("invalid topic format")
}

func getDecodeValue(v *mxtag_pb.Value, t int32) *Value {
	value := &Value{}
	switch t {
	case TAG_VALUE_TYPE_BOOLEAN:
		value.i = v.GetIntValue()
	case TAG_VALUE_TYPE_INT, TAG_VALUE_TYPE_INT8,
		TAG_VALUE_TYPE_INT16, TAG_VALUE_TYPE_INT32:
		value.i = v.GetIntValue()
	case TAG_VALUE_TYPE_UINT, TAG_VALUE_TYPE_UINT8,
		TAG_VALUE_TYPE_UINT16, TAG_VALUE_TYPE_UINT32:
		value.u = v.GetUintValue()
	case TAG_VALUE_TYPE_FLOAT:
		value.f = v.GetFloatValue()
	case TAG_VALUE_TYPE_DOUBLE:
        value.d = Round(v.GetDoubleValue(), 0.000001)
	case TAG_VALUE_TYPE_STRING:
		value.s = v.GetStrValue()
	case TAG_VALUE_TYPE_BYTEARRAY:
		value.b = v.GetBytesValue()
	default:
		logger.Debugf("decode invalid value type: %v", t)
	}
	return value
}

func DecodePayload(payload []byte, tag *Tag) error {
	data := &mxtag_pb.Tag{}
	if err := proto.Unmarshal(payload, data); err != nil {
		return fmt.Errorf("unmarshaling error: %v", err)
	}
	tag.sourceName = data.GetEquipment()
	tag.tagName = data.GetTag()
	tag.val = getDecodeValue(data.GetValue(), data.GetValueType())
	tag.valType = data.GetValueType()
	tag.ts = data.GetAtMs()
	tag.unit = data.GetUnit()
	return nil
}

func ProtobufToTag(protobuf []*mxtag_pb.Tag) []Tag {
	list := make([]Tag, 0, len(protobuf))
	for _, data := range protobuf {
		tag := Tag{}
		tag.sourceName = data.GetEquipment()
		tag.tagName = data.GetTag()
		tag.val = getDecodeValue(data.GetValue(), data.GetValueType())
		tag.valType = data.GetValueType()
		tag.ts = data.GetAtMs()
		tag.unit = data.GetUnit()
		list = append(list, tag)
	}
	return list
}
