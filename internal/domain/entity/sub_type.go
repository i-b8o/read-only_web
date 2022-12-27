package entity

type DocSubType struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name"`
	TypeID uint64 `json:"type_id,omitempty"`
}
