package event

import "github.com/vmihailenco/msgpack/v4"

type SpecimenEvent struct {
	*Base
}

func (o *SpecimenEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *SpecimenEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
