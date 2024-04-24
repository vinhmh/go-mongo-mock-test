package models

type FilterDetail struct {
	FieldName string      `json:"fieldName" form:"fieldName"`
	Value     interface{} `json:"value" form:"value"`
}
type Filter struct {
	Limit         int64          `form:"limit" binding:"required"`
	Offset        int64          `form:"offset"`
	FilterDetails []FilterDetail `form:"filterDetails"`
}
