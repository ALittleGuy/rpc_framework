package giface

import "io"

type IDataPack interface {
	Pack(message IMessage) ([]byte, error)
	UnPack(reader io.Reader) (IMessage, error)
}
