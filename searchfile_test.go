package linkedpackage

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		dir        string
		extensions []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty folder",
			args: args{
				dir: "testdata/searchtest/empty",
				extensions: []string{".js.map"},
			},
			want: []string{},
		},
		{
			name: "search single folder",
			args: args{
				dir: "testdata/searchtest/subdir",
				extensions: []string{".js.map"},
			},
			want: []string{
				"testdata/searchtest/subdir/test1.js.map",
				"testdata/searchtest/subdir/test3.js.map",
			},
		},
		{
			name: "search folder recursively",
			args: args{
				dir: "testdata/searchtest",
				extensions: []string{".js.map"},
			},
			want: []string{
				"testdata/searchtest/subdir/test1.js.map",
				"testdata/searchtest/subdir/test3.js.map",
				"testdata/searchtest/test2.js.map",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.dir, tt.args.extensions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}