package service

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestZipFiles(t *testing.T) {
	type args struct {
		c     *gin.Context
		files []string
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"ZipFiles test 1", args{c, []string{"../sampleDirectory/user/log.txt"}}, 1, false},
		{"ZipFiles test 2", args{c, []string{"../sampleDirectory/user/messages.txt"}}, 1, false},
		{"ZipFiles test 3", args{c, []string{"../sampleDirectory/user/nofile.txt"}}, 0, false},
		{"ZipFiles test 4", args{c, []string{"../sampleDirectory/user/log.txt", "../sampleDirectory/user/messages.txt"}}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZipFiles(tt.args.c, tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZipFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZipFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadNzipFile(t *testing.T) {
	type args struct {
		zw  *zip.Writer
		val string
	}
	a := zip.NewWriter(new(gin.Context).Writer)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"readNzipFile test 1", args{a, "../sampleDirectory/user/log.txt"}, false},
		{"readNzipFile test 2", args{a, "../sampleDirectory/user/messages.txt"}, false},
		{"readNzipFile test 3", args{a, "./nofile"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := readNzipFile(tt.args.zw, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZipFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
