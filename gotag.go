package gotag

import (
    "errors"
    logger "github.com/sirupsen/logrus"
)

type Tagf struct {
    client          MsgQueueBase
}

func getHost(host []string) string {
    if len(host) > 0 {
        return host[0]
    }
    return ""
}

func NewClient(host ...string) (*Tagf, error) {
    config := &MQConfig{
                Host: getHost(host),
                Port: "1883",
                Qos: 0,
                Retained: false,
             }

    c, err := NewMqtt(config)
    if err != nil {
        return nil, err
    }

    logger.SetLevel(logger.ErrorLevel)

    return &Tagf{
        client: c,
    }, nil
}

func(self *Tagf) SetLogLevel(level string) error {
    if level == "info" {
        logger.SetLevel(logger.InfoLevel)
    } else if level == "debug" {
        logger.SetLevel(logger.DebugLevel)
    } else if level == "warn" {
        logger.SetLevel(logger.WarnLevel)
    } else if level == "error" {
        logger.SetLevel(logger.ErrorLevel)
    } else {
        return errors.New("Level not defined")
    }
    return nil
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
    return self.client.UnSubscribe(topic)
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
