package conn

import (
	"encoding/binary"
	"net"

	"github.com/yedamao/go_sgip/sgip/errors"
	"github.com/yedamao/go_sgip/sgip/protocol"
)

// Conn is a sgip connection can read/write protocol Operation
type Conn struct {
	net.Conn
}

func (c *Conn) Read() (protocol.Operation, error) {
	l := make([]byte, 4)
	_, err := c.Conn.Read(l)
	if err != nil {
		return nil, err

	}

	length := binary.BigEndian.Uint32(l) - 4
	if length > protocol.MAX_OP_SIZE {
		return nil, errors.SgipSizeErr
	}

	data := make([]byte, length)

	i, err := c.Conn.Read(data)
	if err != nil {
		return nil, err
	}

	if i != int(length) {
		return nil, errors.SgipLenErr
	}

	pkt := append(l, data...)

	op, err := protocol.ParseOperation(pkt)
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (c *Conn) Write(op protocol.Operation) error {
	_, err := c.Conn.Write(op.Serialize())

	return err
}

func (c *Conn) Close() {
	c.Conn.Close()
}
