package entity

import (
	"reflect"
	"testing"
)

func TestLoadWikiDocument(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *WikiDocument
		wantErr bool
	}{
		{
			name: "正确读取文件",
			args: args{fileName: "../doc/wiki-test.json"},
			want: &WikiDocument{
				Entries: []WikiEntry{
					{
						ID:    "1",
						Title: "你好",
						Text:  "世界",
					},
				},
				CreateTime: 1653321383,
				ModifyTime: 1653321383,
			},
			wantErr: false,
		}, {
			name:    "错误读取文件",
			args:    args{fileName: "./doc/not-found.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadWikiDocument(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadWikiDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadWikiDocument() got = %v, want %v", got, tt.want)
			}
		})
	}
}
