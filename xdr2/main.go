package xdr

type Enum interface {
	ValidEnum(int32) bool
}

type Union interface {
	ArmForSwitch(int32) (string, bool)
	SwitchFieldName() string
}
