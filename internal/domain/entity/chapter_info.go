package entity

type ChapterInfo struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name"`
	Num      string `json:"num,omitempty"`
	OrderNum uint32 `json:"order_num"`
}

func (c ChapterInfo) IsEmpty() bool {
	if (c.Name == "") && (c.Num == "") {
		return true
	}
	return false
}
