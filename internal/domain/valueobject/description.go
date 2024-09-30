package valueobject

type Description struct {
	value string
}

func NewDescription(value string) *Description {
	return &Description{value: value}
}

func (description *Description) GetValue() string {
	return description.value
}