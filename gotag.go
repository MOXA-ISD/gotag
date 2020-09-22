package gotag

import (
	"errors"
)

type Tagf struct {
	client *DataExchange
}

func NewClient() (*Tagf, error) {
	c := NewDataExchange()
	if c == nil {
		return nil, errors.New("create data exchange client failed")
	}
	return &Tagf{client: c}, nil
}

func (self *Tagf) Publish(modName, srcName, tagName string, val *Value, valType uint16, timestamp uint64) error {
	if !(self != nil && self.client != nil) {
		return errors.New("tag client not found")
	}
	topic := EncodeTopic(modName, srcName, tagName)
	return self.client.Publish(topic, valType, val, timestamp)
}

func (self *Tagf) Subscribe(modName, sourceName, tagName string) error {
	if !(self != nil && self.client != nil) {
		return errors.New("tag client not found")
	}
	topic := EncodeTopic(modName, sourceName, tagName)
	return self.client.Subscribe(topic)
}

func (self *Tagf) SubscribeCallback(ontag OnTagCallback) error {
	if !(self != nil && self.client != nil) {
		return errors.New("tag client not found")
	}
	return self.client.SubscribeCallback(ontag)
}

func (self *Tagf) UnSubscribe(modName, sourceName, tagName string) error {
	if !(self != nil && self.client != nil) {
		return errors.New("tag client not found")
	}
	topic := EncodeTopic(modName, sourceName, tagName)
	return self.client.UnSubscribe(topic)
}

func (self *Tagf) Unmarshal(module, source, tag string) *Tag {
	return nil
}

func (self *Tagf) Marshal(module, source, tag string) *Tag {
	return nil
}

func (self *Tagf) Delete() error {
	if !(self != nil && self.client != nil) {
		return errors.New("tag client not found")
	}
	return self.client.Close()
}
