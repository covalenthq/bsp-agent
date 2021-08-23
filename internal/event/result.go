package event

import "github.com/vmihailenco/msgpack/v4"

type ResultEvent struct {
	*Base
}

func (o *ResultEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *ResultEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
