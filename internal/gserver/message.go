package gserver

type Message struct {
	MagicNum  uint16
	Body      []byte
	Header    []byte
	Len       uint32
	HeaderLen uint32
}

func NewMessage(body []byte, header []byte) *Message {
	return &Message{
		Body:      body,
		Header:    header,
		HeaderLen: uint32(len(header)),
		Len:       uint32(len(header) + len(body) + 4),
	}
}

// magic num
const magicNum = 0x19a

func (d *Message) GetLen() uint32 {
	return d.Len
}

func (d *Message) GetHeaderLen() uint32 {
	return d.HeaderLen
}

func (d *Message) GetHeader() []byte {
	return d.Header
}

func (d *Message) GetBody() []byte {
	return d.Body
}
