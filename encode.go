package gotag

import (
    "strings"

    "github.com/golang/protobuf/proto"
    "github.com/CPtung/gotag/protobuf"
    logger "github.com/sirupsen/logrus"
)

func EncodeTopic(source, tag string) string {
    var topic strings.Builder
    topic.WriteString("/equs/")
    topic.WriteString(source)
    topic.WriteString("/tags/")
    topic.WriteString(tag)
    return topic.String()
}

func getEncodeValue(val *Value, valType int32) *mxtag_pb.Value {
    v := &mxtag_pb.Value{}
    switch (valType) {
        case TAG_VALUE_TYPE_BOOLEAN:
            v.IntValue = &val.i
            break
        case TAG_VALUE_TYPE_INT, TAG_VALUE_TYPE_INT8,
                 TAG_VALUE_TYPE_INT16, TAG_VALUE_TYPE_INT32:
            v.IntValue = &val.i
            break
        case TAG_VALUE_TYPE_UINT, TAG_VALUE_TYPE_UINT8,
                 TAG_VALUE_TYPE_UINT16, TAG_VALUE_TYPE_UINT32:
            v.UintValue = &val.u
            break
        case TAG_VALUE_TYPE_FLOAT:
            v.FloatValue = &val.f
            break
        case TAG_VALUE_TYPE_DOUBLE:
            v.DoubleValue = &val.d
            break
        case TAG_VALUE_TYPE_STRING:
            v.StrValue = &val.s
            break
        case TAG_VALUE_TYPE_BYTEARRAY:
            v.BytesValue = val.b
            break
    }
    return v
}

func EncodePayload(source string, tag string, value *Value, valtype int32, at uint64, unit string) []byte {
    p := &mxtag_pb.Tag{
            Equipment:	proto.String(source),
            Tag:		proto.String(tag),
            AtMs:		proto.Uint64(at),
            Value:		getEncodeValue(value, valtype),
            ValueType:	proto.Int32(valtype),
            Unit:		proto.String(unit),
        }
    data, err := proto.Marshal(p)
    if err != nil {
        logger.Error("Marshal tag protobuf got error (%v)\n", err)
        return nil
    }
    return data
}
