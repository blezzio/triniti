package interfaces

type CallLog interface {
	Called(f any) (int, []any)
}
