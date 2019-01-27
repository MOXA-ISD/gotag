package gotag

import (
    "errors"

    zmq "github.com/pebbe/zmq4"
)

type ZmqClient struct {
    publisher   *zmq.Socket
    subscriber  *zmq.Socket
}

type TpZmq struct {
    MsgQueueBase
    stop    bool
    c       *ZmqClient
    ontag   OnTagCallback
}

func(self *TpZmq)Publish(topic string, payload []byte) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    if _, err := self.c.publisher.Send(topic, zmq.SNDMORE); err != nil {
        return err
    }
    if _, err := self.c.publisher.SendBytes(payload, 0); err != nil {
        return err
    }
    return nil
}

func(self *TpZmq)Subscribe(topic string) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    if self.ontag == nil {
        return errors.New("needs to assign a handler before subscribe topics")
    }
    if err := self.c.subscriber.SetSubscribe(topic); err != nil {
        return err
    }
    return nil
}

func(self *TpZmq)Unsubscribe(topic string) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    if err := self.c.subscriber.SetUnsubscribe(topic); err != nil {
        return err
    }
    return nil
}

func(self *TpZmq)SubscribeCallback(hnd OnTagCallback) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    self.ontag = hnd
    return nil
}

func(self *TpZmq)Close() error {
    self.stop = true
    if self.c != nil {
        if self.c.subscriber != nil {
            self.c.subscriber.Close()
        }
        self.c.subscriber = nil
        if self.c.publisher != nil {
            self.c.publisher.Close()
        }
        self.c.publisher = nil
        self.c = nil
    }
    return nil
}

func(self *TpZmq)onMessageReceived() {
    for !self.stop {
        //  Read topic with address
        topic, terr := self.c.subscriber.Recv(0)
        if terr != nil {
            return
        }
        //  Read message contents
        payload, perr := self.c.subscriber.RecvBytes(0)
        if perr != nil {
            return
        }
        //  Decode topic to get source and tag name
        srcName, tagName := DecodeTopic(topic)
        if srcName == "" || tagName == "" {
            return
        }
        //  Decode payload to tag
        tag := &Tag{}
        err := DecodePayload(payload, tag)
        if err != nil {
            return
        }
        if self.ontag != nil {
            self.ontag(tag.sourceName, tag.tagName, tag.val, tag.valType, tag.ts, tag.unit)
        }
        tag = nil
    }
}

func NewZmq() (*TpZmq, error) {
    var err error
    t := &TpZmq{
        c: &ZmqClient{
            subscriber: nil,
            publisher:  nil,
        },
    }
    if t.c.subscriber, err = zmq.NewSocket(zmq.SUB); err != nil {
        t.Close()
        return nil, err
    }
	if err = t.c.subscriber.Connect("ipc:///run/mxtagf/xpub.ipc"); err != nil {
        t.Close()
        return nil, err
    }

    if t.c.publisher, err = zmq.NewSocket(zmq.PUB); err != nil {
        t.Close()
        return nil, err
    }
    if err = t.c.publisher.Connect("ipc:///run/mxtagf/xsub.ipc"); err != nil {
        t.Close()
        return nil, err
    }
    // start recv message thread
    go t.onMessageReceived()
    return t, nil
}
