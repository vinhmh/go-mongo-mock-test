package forms

type FilterDetail struct {
	FieldName string
	Value     interface{}
}
type Filter struct {
	Limit         int64
	Offset        int64
	FilterDetails []FilterDetail
}
