package safenet

import (
	"log"
	"os"
	"testing"
)

var verifyOutput = `
The following Luna SA Slots/Partitions were found:
Slot    Serial #        Label
====    ========        =====
 1      345435345435       A-Company

`

var pwd string

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Panic(err)
	}
}

func Test_runVtl(t *testing.T) {

	type args struct {
		dir  string
		app  string
		args string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{
			dir:  pwd,
			app:  "vtl.exe",
			args: "-verify",
		}, verifyOutput, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vtl{
				dir: tt.args.dir,
				app: tt.args.app,
			}
			_, _, err := v.runVtl(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runVtl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_vtl_Verify(t *testing.T) {
	testDir := pwd

	type fields struct {
		dir  string
		app  string
		args string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"", fields{testDir, "vtl.exe", "123213"}, true},
		{"", fields{testDir, "vtl.exe", "535788010"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vtl{
				dir: tt.fields.dir,
				app: tt.fields.app,
			}
			err := v.Verify(tt.fields.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
