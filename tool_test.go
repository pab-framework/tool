package tool

import (
	"reflect"
	"testing"
)

func TestBytes2File(t *testing.T) {
	type args struct {
		data   []byte
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_Bytes2File_1",
			args: args{
				data:   []byte("abc=123"),
				target: "bytes2file.properties",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bytes2File(tt.args.data, tt.args.target)
		})
	}
}

func TestCopyFile(t *testing.T) {
	type args struct {
		source string
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_CopyFile_1",
			args: args{
				source: "bytes2file.properties",
				target: "bytes2file_copy.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyFile(tt.args.source, tt.args.target)
		})
	}
}

func TestFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_FileExist_1",
			args: args{path: "bytes2file.properties"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExist(tt.args.path); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoroutineID(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{
			name: "test_GetGoroutineID_1",
			want: uint64(7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGoroutineID(); got != tt.want {
				t.Errorf("GetGoroutineID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties2Map(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "test_Properties2Map_1",
			args:    args{path: "bytes2file.properties"},
			want:    getMap(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Properties2Map(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Properties2Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Properties2Map() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func getMap() map[string]string {
	m := make(map[string]string)
	m["abc"] = "123"
	return m
}
