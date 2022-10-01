package memontago

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	type args struct {
		datetime interface{}
		format   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//{"5秒前", args{time.Now().Add(-5 * time.Second)}, "5秒前", false},
		//{"5分钟前", args{time.Now().Add(-5 * time.Minute)}, "5分钟前", false},
		//{"5小时后", args{time.Now().Add(5 * time.Hour)}, "5小时后", false},
		//{"5秒前 special", args{time.Now().Add(5 * time.Second)}, "在线", false},
		//{"29秒前 special", args{time.Now().Add(29 * time.Second)}, "刚刚", false},

		//{"5 seconds ago", args{time.Now().Add(-5 * time.Second)}, "5 seconds ago", false},
		//{"5 seconds later", args{time.Now().Add(5 * time.Second)}, "5 seconds later", false},
		{"test", args{time.Now().Format("2006-01-02 15"), "2006-01-02 15"}, "46分钟前", true},
		{"test", args{time.Now().Add(-5 * time.Minute).Format("2006-01-02 15:04:05"), ""}, "5分钟前", true},
	}
	config.Language = "ch"
	config.Special = true
	UnMarshalToStruct(config.Language)
	SetConfig(config)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.datetime, tt.args.format)
			if err != nil {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
