package core

import "encoding/json"

type JSONEncoder struct{}

func (e *JSONEncoder) Encode(msg PoseMessage) ([]byte, error) {
    return json.Marshal(msg)
}

func (e *JSONEncoder) Decode(data []byte) (PoseMessage, error) {
    var msg PoseMessage
    err := json.Unmarshal(data, &msg)
    return msg, err
}
