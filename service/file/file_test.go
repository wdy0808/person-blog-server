package file

import (
	"testing"
)

func TestFileExist(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"file", args{"/Users/wudayuan/Documents/2019/person-blog-server/main.go"}, true},
		{"dir", args{"/Users/wudayuan/Documents/2019/person-blog-server"}, true},
		{"no_exist", args{"/Users/wudayuan/Documents/2019/person-blog-se"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExist(tt.args.file); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
