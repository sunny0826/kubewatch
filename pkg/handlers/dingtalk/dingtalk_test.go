package dingtalk

import (
	"github.com/sunny0826/kubewatch/config"
	"testing"
)

func TestDingTalk_Init(t *testing.T) {
	type fields struct {
		Token string
		Sign  string
	}
	type args struct {
		c *config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{fields: fields{Token: "39587a4c5cd671e1f53ac706e96271eb3855f8c71c3b05f844efa8b3bf4da833", Sign: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DingTalk{
				Token: tt.fields.Token,
				Sign:  tt.fields.Sign,
			}
			if err := d.Init(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
