package entity

// WikiDocument wiki文档结构
type WikiDocument struct {
	Entries    []WikiEntry `json:"entries"`
	ModifyTime int64       `json:"modify_time"`
	CreateTime int64       `json:"create_time"`
}

// WikiEntry wiki词条
type WikiEntry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
