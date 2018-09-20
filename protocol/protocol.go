package protocol

type IProtocol interface {
	DealData([]byte) (Packet, error)
	Name() string
}

type Packet interface {
	Byte() []byte
	String() string
	IsEmpty() bool
}
