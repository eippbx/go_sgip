// 定义sgip中的字段类型

package protocol

import (
	"bytes"
	"encoding/binary"
)

// 不强制以0x00结尾的定长字符串。
// 当位数不足时，在不明确注明的
// 情况下， 应左对齐，右补0x00
type OctetString struct {
	Data     []byte // 数据 未补零/已补零
	FixedLen int    // 协议中该参数的固定长度
}

// 去除补零，转为字符串
func (o *OctetString) String() string {

	end := bytes.IndexByte(o.Data, 0)
	if -1 == end {
		return string(o.Data)
	}

	return string(o.Data[:end])
}

// 按需补零
func (o *OctetString) Byte() []byte {
	if len(o.Data) < o.FixedLen {
		// fill 0x00
		tmp := make([]byte, o.FixedLen-len(o.Data))
		o.Data = append(o.Data, tmp...)
	}

	return o.Data
}

func unpackUi32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func packUi32(n uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, n)
	return b

}

func unpackUi16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func packUi16(n uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, n)
	return b
}

func packUi8(n uint8) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(n))
	return b[1:]
}
