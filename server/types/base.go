package types

import jsoniter "github.com/json-iterator/go"

type Base struct {
}

func (b *Base) JSON() string {
	str, _ := jsoniter.ConfigDefault.MarshalToString(b)
	return str
}

func (b *Base) Decode(str string) error {
	return jsoniter.ConfigDefault.UnmarshalFromString(str, b)
}
