package dingtalk

import (
	"github.com/sunny0826/kubewatch/config"
	"testing"
)

func TestDingContent_sendDingContent(t *testing.T) {
	type fields struct {
		Title   string
		Message string
	}
	type args struct {
		color string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DingContent{
				Title:   tt.fields.Title,
				Message: tt.fields.Message,
			}
			if got := s.sendDingContent(tt.args.color); got != tt.want {
				t.Errorf("sendDingContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{fields: fields{Token: "", Sign: ""}},
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

func TestDingTalk_ObjectCreated(t *testing.T) {
	type fields struct {
		Token string
		Sign  string
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DingTalk{
				Token: tt.fields.Token,
				Sign:  tt.fields.Sign,
			}
		})
	}
}

func TestDingTalk_ObjectDeleted(t *testing.T) {
	type fields struct {
		Token string
		Sign  string
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DingTalk{
				Token: tt.fields.Token,
				Sign:  tt.fields.Sign,
			}
		})
	}
}

func TestDingTalk_ObjectUpdated(t *testing.T) {
	type fields struct {
		Token string
		Sign  string
	}
	type args struct {
		oldObj interface{}
		newObj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DingTalk{
				Token: tt.fields.Token,
				Sign:  tt.fields.Sign,
			}
		})
	}
}

func TestHandler(t *testing.T) {
	type fields struct {
		Token string
		Sign  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DingTalk{
				Token: tt.fields.Token,
				Sign:  tt.fields.Sign,
			}
		})
	}
}

func Test_checkMissingSlackVars(t *testing.T) {
	type args struct {
		d *DingTalk
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkMissingSlackVars(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("checkMissingSlackVars() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_notifySlack(t *testing.T) {
	type args struct {
		d      *DingTalk
		obj    interface{}
		action string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
