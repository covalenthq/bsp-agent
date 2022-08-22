package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

// BigInt is a big.Int wrapper which marshals/unmarshals into byte arrays
type BigInt struct {
	*big.Int
}

// SetUint64 sets BigInt from a uint64 value
func (x *BigInt) SetUint64(value uint64) *BigInt {
	if x.Int == nil {
		x.Int = new(big.Int)
	}

	_ = x.Int.SetUint64(value)

	return x
}

// MarshalText implements TextMarshaler
func (x *BigInt) MarshalText() (text []byte, err error) {
	if x == nil {
		return []byte("[<nil>]"), nil
	}

	slice := []byte("\"0x")

	slice = append(slice, []byte(x.Int.Text(16))...)
	slice = append(slice, []byte("\"")...)

	return slice, nil
}

// UnmarshalText implements TextUnmarshaler
func (x *BigInt) UnmarshalText(text []byte) error {
	// ignore the opening and end quotes
	if x.Int == nil {
		x.Int = new(big.Int)
	}

	text = text[1 : len(text)-1]
	if _, success := x.Int.SetString(string(text), 0); !success {
		return fmt.Errorf("failed to unmarshal text")
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (x *BigInt) MarshalJSON() ([]byte, error) {
	return x.MarshalText()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (x *BigInt) UnmarshalJSON(text []byte) error {
	return x.UnmarshalText(text)
}

// DecodeRLP implements rlp.Decoder
func (x *BigInt) DecodeRLP(s *rlp.Stream) error {
	if x.Int == nil {
		x.Int = new(big.Int)
	}
	err := decodeBigInt(s, x.Int)
	if err != nil {
		return err
	}

	return nil
}

func decodeBigInt(s *rlp.Stream, val *big.Int) error {
	internal, err := s.BigInt()
	if err != nil {
		return fmt.Errorf("%w, %v", err, val)
	}

	if val != nil {
		val.Set(internal)
	} else {
		return fmt.Errorf("val is nil, can't set big.Int value")
	}

	return nil
}
