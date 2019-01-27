package gotag

import (
    "log"
    "sync"
    "time"
    "errors"
    "math/rand"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

type TpMqtt struct {
    MsgQueueBase
    c		mqtt.Client
    wg		*sync.WaitGroup
    ontag	OnTagCallback
}

func(self *TpMqtt)Publish(topic string, payload []byte) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    token := self.c.Publish(topic, 0, false, payload)
    token.Wait()
    return nil
}

func(self *TpMqtt)Subscribe(topic string) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    if self.ontag == nil {
        return errors.New("needs to assign a handler before subscribe topics")
    }
    if token := self.c.Subscribe(topic, 0, self.onMessageReceived); token.Wait() && token.Error() != nil {
        return token.Error()
    }
    return nil
}

func(self *TpMqtt)Unsubscribe(topic string) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    if token := self.c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
        return token.Error()
    }
    return nil
}

func(self *TpMqtt)SubscribeCallback(hnd OnTagCallback) error {
    if self.c == nil {
        return errors.New("tag client not found")
    }
    self.ontag = hnd
    return nil
}

func(self *TpMqtt)Close() error {
    if self.c != nil {
        self.c.Disconnect(250)
    }
    return nil
}

func(self *TpMqtt)onMessageReceived(client mqtt.Client, message mqtt.Message) {
    if self.ontag == nil {
        return
    }
    srcName, tagName := DecodeTopic(message.Topic())
    if srcName == "" || tagName == "" {
        log.Println("wrong topic format")
        return
    }
    t := &Tag{}
    err := DecodePayload(message.Payload(), t)
    if err != nil {
        log.Printf("on message received error (%v)", err)
        return
    }
    self.ontag(t.sourceName, t.tagName, t.val, t.valType, t.ts, t.unit)
}

func (self *TpMqtt)OnConnectHandler(client mqtt.Client) {
    self.wg.Done()
}

func (self *TpMqtt)OnDisconnectHandler(client mqtt.Client, err error) {
    log.Printf("client disconnected")
    if err != nil {
        log.Printf(" error (%v)", err)
    }
    log.Printf("\n")
}

func (self *TpMqtt)OnPublishHandler(client mqtt.Client, message mqtt.Message) {
    log.Printf("Topic: %v\n", message.Topic())
    log.Printf("Msg: %v\n", message.Payload())
}

func NewMqtt(cfg *MQConfig) (*TpMqtt, error) {
    rand.Seed(time.Now().UnixNano())
    t := &TpMqtt{
            wg: &sync.WaitGroup{},
            ontag: nil,
    }

    opts := mqtt.NewClientOptions()
    if cfg.Host == "" {
        opts.AddBroker("tcp://" + getEnv("APPMAN_TAGSERVICE_ADDR", "localhost") + ":" + cfg.Port)
    } else {
        opts.AddBroker("tcp://" + cfg.Host + ":" + cfg.Port)
    }
    opts.SetClientID(genId(8))
    opts.SetKeepAlive(30 * time.Second)
    opts.SetMaxReconnectInterval(3)
    opts.SetCleanSession(true)
    opts.SetOnConnectHandler(t.OnConnectHandler)
    opts.SetConnectionLostHandler(t.OnDisconnectHandler)
    opts.SetDefaultPublishHandler(t.OnPublishHandler)

    t.c = mqtt.NewClient(opts)
    if token := t.c.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }

    waitTimeout(t.wg, 3 * time.Second)
    return t, nil
}
