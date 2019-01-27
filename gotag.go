package gotag

import (
    "log"
    "errors"
)

type Tagf struct {
    client          MsgQueueBase
    topics          []string
    nums_of_topics  int32
}

func getHost(host []string) string {
    if len(host) > 0 {
        return host[0]
    }
    return ""
}

func NewClient(host ...string) *Tagf {
    /*config := &MQConfig{
                Host: getHost(host),
                Port: "1883",
                Qos: 0,
                Retained: false,
            }
    */
    c, err := NewZmq() //NewMqtt(config)
    if c == nil  && err != nil {
        log.Printf("NewClient Err (%v)\n", err)
        return nil
    }

    return &Tagf{
        client: c,
        topics: []string{},
        nums_of_topics: 0,
    }
}

func(self *Tagf) Publish(sourceName string, tagName string, val *Value, valType int32, timestamp uint64, unit string) error {
    if !(self != nil && self.client != nil) {
        return errors.New("tag client not found")
    }
    topic := EncodeTopic(sourceName, tagName)
    payload := EncodePayload(sourceName, tagName, val, valType, timestamp, unit)
    if payload == nil {
        return errors.New("Invalid Input")
    }
    return self.client.Publish(topic, payload)
}

func(self *Tagf) Subscribe(sourceName, tagName string) error {
    if !(self != nil && self.client != nil) {
        return errors.New("tag client not found")
    }
    topic := EncodeTopic(sourceName, tagName)
    return self.client.Subscribe(topic)
}

func(self *Tagf) UnSubscribe(sourceName, tagName string) error {
    if !(self != nil && self.client != nil) {
        return errors.New("tag client not found")
    }
    topic := EncodeTopic(sourceName, tagName)
    return self.client.Subscribe(topic)
}

func(self *Tagf) SubscribeCallback(ontag OnTagCallback) error {
    if !(self != nil && self.client != nil) {
        return errors.New("tag client not found")
    }
    return self.client.SubscribeCallback(ontag)
}

func(self *Tagf) Delete() error {
    if !(self != nil && self.client != nil) {
        return errors.New("tag client not found")
    }
    return errors.New("Tag client not found")
}
