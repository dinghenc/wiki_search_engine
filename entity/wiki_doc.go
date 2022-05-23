package entity

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// WikiDocument wiki文档结构
type WikiDocument struct {
	Entries    []WikiEntry `json:"entries"`
	CreateTime int64       `json:"create_time"`
	ModifyTime int64       `json:"modify_time"`
}

// WikiEntry wiki词条
type WikiEntry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// LoadWikiDocument 从文件中加载wiki文档
func LoadWikiDocument(fileName string) (*WikiDocument, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "open file=%s failed", fileName)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "read all content from file failed")
	}
	doc := new(WikiDocument)
	if err := json.Unmarshal(bytes, doc); err != nil {
		return nil, errors.Wrap(err, "json unmarshal wiki document failed")
	}
	return doc, nil
}
