package datasource

import (
	// "github.com/lovoo/goka"
)

// type GokaProcessCallback interface {
// 	Handle(goka.Context, interface{})
// }

type GokaMessageCodec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}