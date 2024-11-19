package types

type MemeStruct struct {
	Owner   string `json:"owner" binding:"required"`
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type MemeMeta struct {
	File  string
	Start uint64
	Size  uint64
}
