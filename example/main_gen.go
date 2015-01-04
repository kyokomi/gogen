package main

import (
	"github.com/philhofer/msgp/msgp"
)


// DecodeMsg implements the msgp.Decodable interface
func (z *Hoge) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}

		// TODO:
	}

	return
}
