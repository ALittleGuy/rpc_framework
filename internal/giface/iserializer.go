package giface

type ISerializer interface {
	Encoder(obj struct{}) []byte
	Decoder([]byte) struct{}
}
