package entity

import "time"

type Search struct {
	RID       *uint64    `json:"r_id"`
	RName     *string    `json:"r_name"`
	CID       *uint64    `json:"c_id"`
	CName     *string    `json:"c_name"`
	UpdatedAt *time.Time `json:"updated_at"`
	PID       *uint64    `json:"p_id"`
	Text      *string    `json:"text"`
	Count     *uint64    `json:"count"`
}
