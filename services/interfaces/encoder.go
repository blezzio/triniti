package interfaces

type Encoding interface {
	EncodeToString(src []byte) string
}
