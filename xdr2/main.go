package xdr

type Enum interface {
	ValidEnum(int32) bool
}
