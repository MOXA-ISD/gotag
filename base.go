package gotag

type MQConfig struct {
    Id          string
    Host        string
    Port        string
    Debug       string
    Qos         byte
    Retained    bool
}

type OnTagCallback func(
    moduleName  string,
    sourceName  string,
    tagName     string,
    value       *Value,
    valueType   uint16,
    timestamp   uint64)
