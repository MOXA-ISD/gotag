package gotag

import (
    "time"
    "sync"
    "fmt"
    "errors"

    zmq "github.com/pebbe/zmq4"
)

const SIG_MXTAGF_STOP = "mxtagf/signal/stop/"

type ZmqClient struct {
    publisher   *zmq.Socket
    subscriber  *zmq.Socket
}

type TpZmq struct {
    MsgQueueBase
    c       *ZmqClient
    isRun   bool
    sigStop string
    ontag   OnTagCallback
    wg      *sync.WaitGroup
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
    self.isRun = false
    if self.c != nil {
        self.wg.Add(1)
        if _, err := self.c.publisher.Send(self.sigStop, 0); err != nil {
            return err
        }
        self.wg.Wait()
    }
    return nil
}

func(self *TpZmq)onMessageReceived() {
    defer self.c.subscriber.Close()
    defer self.c.publisher.Close()
    defer self.wg.Done()
    for self.isRun {
        //  Read topic with address
        topic, terr := self.c.subscriber.Recv(0)
        if terr != nil || topic == self.sigStop {
            continue
        }
        //  Read message contents
        payload, perr := self.c.subscriber.RecvBytes(0)
        if perr != nil {
            continue
        }
        //  Decode topic to get source and tag name
        srcName, tagName := DecodeTopic(topic)
        if srcName == "" || tagName == "" {
            continue
        }
        //  Decode payload to tag
        tag := &Tag{}
        err := DecodePayload(payload, tag)
        if err == nil && self.ontag != nil {
            self.ontag(tag.sourceName, tag.tagName, tag.val, tag.valType, tag.ts, tag.unit)
        }
        tag = nil
    }
}

func NewZmq() (*TpZmq, error) {
    var err error
    t := &TpZmq{
        isRun: true,
        wg: &sync.WaitGroup{},
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
    time.Sleep(time.Millisecond)

    // hook stop signal
    t.sigStop = fmt.Sprintf("%v%v", SIG_MXTAGF_STOP, genId(8))
    if err := t.c.subscriber.SetSubscribe(t.sigStop); err != nil {
        return nil, err
    }
    return t, nil
}
