package giface

type IRequest interface {
	Encode(msg *IMessage) ([]byte, error)
	Decode(request interface{}) error
}
