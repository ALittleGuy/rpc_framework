package giface

type IMessage interface {

	GetLen() uint32
	GetHeaderLen() uint32
	GetHeader() []byte
	GetBody() []byte
}
