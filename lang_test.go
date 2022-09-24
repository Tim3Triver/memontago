package memontago

import (
	"fmt"
	"testing"
)

func TestMarshalToStruct(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"ch"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnMarshalToStruct(tt.args.lang); err != nil {
				t.Errorf("UnMarshalToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(Lang)
		})
	}
}
