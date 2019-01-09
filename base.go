package gotag

type MQConfig struct {
	Id			string
	Host		string
	Port		string
	Qos			byte
	Retained	bool
}

type OnTagCallback func(
	sourceName	string,
	tagName		string,
	value		*Value,
	valueType	int32,
	timestamp	uint64,
	unit		string)


type MsgQueueBase interface {
	Publish(topic string, payload []byte)	error
	Subscribe(topic string)					(int32, error)
	Unsubscribe(topic string)				(int32, error)
	SubscribeCallback(ontag OnTagCallback)	(int32, error)
	Close()									(int32, error)
}
