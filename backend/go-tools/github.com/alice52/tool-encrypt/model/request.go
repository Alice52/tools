package model

type Request struct {
	Value string `form:"value" json:"value" binding:"required"`
	Key   string `form:"key" json:"key"`
}
