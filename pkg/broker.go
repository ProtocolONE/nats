package nats_manager

import (
	"encoding/json"
	"errors"
	"github.com/nats-io/stan.go"
	"sync"
	"time"
)

type NatsManagerInterface interface {
	Publish(string, interface{}, bool) error
	QueueSubscribe(string, string, MsgHandler, ...stan.SubscriptionOption) (Subscription, error)
	Close() error
}

type NatsManager struct {
	client  stan.Conn
	options *Options
}

func (m NatsManager) Publish(subject string, msg interface{}, async bool) error {
	var (
		glock sync.Mutex
		guid  string
		ch    = make(chan bool)
	)

	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	acb := func(lguid string, err error) {
		glock.Lock()
		defer glock.Unlock()

		ch <- true
	}

	if !async {
		if err = m.client.Publish(subject, message); err != nil {
			return err
		}
	} else {
		glock.Lock()

		if guid, err = m.client.PublishAsync(subject, message, acb); err != nil {
			return err
		}

		glock.Unlock()

		if guid == "" {
			return errors.New("expected non-empty guid to be returned from the message broker")
		}

		select {
		case <-ch:
			break
		case <-time.After(5 * time.Second):
			return errors.New("timeout to publish message to the message broker")
		}
	}

	return nil
}

func (m NatsManager) QueueSubscribe(subject string, qgroup string, handler MsgHandler, opts ...stan.SubscriptionOption) (Subscription, error) {
	return m.client.QueueSubscribe(subject, qgroup, stan.MsgHandler(handler), opts...)
}

func (m NatsManager) Close() error {
	return m.client.Close()
}
