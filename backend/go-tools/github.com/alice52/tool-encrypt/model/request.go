package model

type Request struct {
	Value string `form:"value" json:"value" binding:"required"`
	Key   string `form:"key" json:"key"`
}

type AesRequest struct {
	Value string `form:"value" json:"value" binding:"required"`
	Key   string `form:"key" json:"key"`
	Iv    string `form:"iv" json:"iv"`
	// Alg   string `form:"alg" json:"alg"` cbc
}
