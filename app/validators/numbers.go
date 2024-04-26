package validators

type Number struct {
	A int32 `json:"a" validate:"required,numeric"`
	B int32 `json:"b" validate:"required,numeric"`
}
