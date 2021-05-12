package linkedpackage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*func TestParseJSSourcemapFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []Module
	}{
		{
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseJSSourcemapFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSSourcemapFile() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestParseWebPackSourcemapFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []Module
		wantErr bool
	}{
		{
			name: "vue-cli project",
			args: args{
				path: "testdata/sources/vue/vue-mini.js.map",
			},
			want: []Module{
				{
					Lang: "js",
					Name: "@babel/runtime",
					Path: "/node_modules/@babel/runtime",
				},
				{
					Lang: "js",
					Name: "ajv",
					Path: "/node_modules/ajv",
				},
				{
					Lang: "js",
					Name: "asn1",
					Path: "/node_modules/asn1",
				},
				{
					Lang: "js",
					Name: "asn1.js",
					Path: "/node_modules/asn1.js",
				},
				{
					Lang: "js",
					Name: "assert",
					Path: "/node_modules/assert",
				},
				{
					Lang: "js",
					Name: "assert-plus",
					Path: "/node_modules/assert-plus",
				},
				{
					Lang: "js",
					Name: "bn.js",
					Path: "/node_modules/asn1.js/node_modules/bn.js",
				},
				{
					Lang: "js",
					Name: "core-js",
					Path: "/node_modules/core-js",
				},
				{
					Lang: "js",
					Name: "vue",
					Path: "/node_modules/vue"},
				{
					Lang: "js",
					Name: "vue-class-component",
					Path: "/node_modules/vue-class-component",
				},
				{
					Lang: "js",
					Name: "vue-gtag",
					Path: "/node_modules/vue-gtag",
				},
				{
					Lang: "js",
					Name: "vue-loader",
					Path: "/node_modules/vue-loader",
				},
				{
					Lang: "js",
					Name: "vue-property-decorator",
					Path: "/node_modules/vue-property-decorator",
				},
				{
					Lang: "js",
					Name: "vue-router",
					Path: "/node_modules/vue-router",
				},
				{
					Lang: "js",
					Name: "vue-style-loader",
					Path: "/node_modules/vue-style-loader",
				},
				{
					Lang: "js",
					Name: "vuex",
					Path: "/node_modules/vuex",
				},
				{
					Lang: "js",
					Name: "vuex-module-decorators",
					Path: "/node_modules/vuex-module-decorators",
				},
				{
					Lang: "js",
					Name: "xtend",
					Path: "/node_modules/xtend",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSSourcemapFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSSourcemapFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSSourcemapFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNCCSourcemapFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []Module
		wantErr bool
	}{
		{
			name: "ncc production build with source-map",
			args: args{
				path: "testdata/ncc-project/dist/index.js.map",
			},
			want: []Module{
				{
					Lang: "js",
					Name: "@date-io/dayjs",
					Path: "/node_modules/@date-io/dayjs",
				},
				{
					Lang: "js",
					Name: "@vercel/ncc",
					Path: "/node_modules/@vercel/ncc",
				},
				{
					Lang: "js",
					Name: "trim",
					Path: "/node_modules/trim",
				},
			},
			wantErr: false,
		},
		{
			name: "ncc development build with source-map",
			args: args{
				path: "testdata/ncc-project/dist_dev/index.js.map",
			},
			want: []Module{
				{
					Lang: "js",
					Name: "@date-io/dayjs",
					Path: "/node_modules/@date-io/dayjs",
				},
				{
					Lang: "js",
					Name: "@vercel/ncc",
					Path: "/node_modules/@vercel/ncc",
				},
				{
					Lang: "js",
					Name: "trim",
					Path: "/node_modules/trim",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSSourcemapFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSSourcemapFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSSourcemapFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJSWebPack(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []Module
		wantErr bool
	}{
		{
			name: "vendor code",
			args: args{
				path: "testdata/sources/vue-devbuild/chunk-vendors.js",
			},
			want: []Module{
				{
					Lang: "js",
					Name: "@vue/reactivity",
					Path: "/node_modules/@vue/reactivity",
				},
				{
					Lang: "js",
					Name: "@vue/runtime-core",
					Path: "/node_modules/@vue/runtime-core",
				},
				{
					Lang: "js",
					Name: "@vue/runtime-dom",
					Path: "/node_modules/@vue/runtime-dom",
				},
				{
					Lang: "js",
					Name: "@vue/shared",
					Path: "/node_modules/@vue/shared",
				},
				{
					Lang: "js",
					Name: "core-js",
					Path: "/node_modules/core-js",
				},
				{
					Lang: "js",
					Name: "css-loader",
					Path: "/node_modules/css-loader",
				},
				{
					Lang: "js",
					Name: "vue",
					Path: "/node_modules/vue",
				},
				{
					Lang: "js",
					Name: "vue-router",
					Path: "/node_modules/vue-router",
				},
				{
					Lang: "js",
					Name: "vue-style-loader",
					Path: "/node_modules/vue-style-loader",
				},
				{
					Lang: "js",
					Name: "vuex",
					Path: "/node_modules/vuex",
				},
			},
			wantErr: false,
		},
		{
			name: "app code",
			args: args{
				path: "testdata/sources/vue-devbuild/about.js",
			},
			want: []Module{
				{
					Lang: "js",
					Name: "babel-loader",
					Path: "/node_modules/babel-loader",
				},
				{
					Lang: "js",
					Name: "cache-loader",
					Path: "/node_modules/cache-loader",
				},
				{
					Lang: "js",
					Name: "vue-loader-v16",
					Path: "/node_modules/vue-loader-v16",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSWebPack(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSWebPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSWebPack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectJSConfigReader(t *testing.T) {
	type args struct {
		module *Module
		root   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Module
		wantErr bool
	}{
		{
			name: "read package.json with simple author",
			args: args{
				module: &Module{
					Lang: "js",
					Name: "sample1",
					Path: "sample1",
				},
				root: "testdata/license",
			},
			want:    &Module{
				Lang: "js",
				Name: "sample1",
				Path: "sample1",
				Author: "abc",
				LicenseName: "MIT",
				LicenseContent: "MIT",
				Version: "1.0.0",
			},
			wantErr: false,
		},
		{
			name: "read package.json with complex author",
			args: args{
				module: &Module{
					Lang: "js",
					Name: "sample2",
					Path: "sample2",
				},
				root: "testdata/license",
			},
			want:    &Module{
				Lang: "js",
				Name: "sample2",
				Path: "sample2",
				Author: "abc <abc@example.com>",
				LicenseName: "MIT",
				LicenseContent: "MIT",
				Version: "1.0.0",
			},
			wantErr: false,
		},
		{
			name: "read package.json with complex licenses",
			args: args{
				module: &Module{
					Lang: "js",
					Name: "sample3",
					Path: "sample3",
				},
				root: "testdata/license",
			},
			want:    &Module{
				Lang: "js",
				Name: "sample3",
				Path: "sample3",
				Author: "Kris Zyp",
				LicenseName: "AFLv2.1, BSD",
				LicenseContent: "",
				Version: "0.2.3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := projectJSConfigReader(tt.args.module, tt.args.root); (err != nil) != tt.wantErr {
				t.Errorf("projectJSConfigReader() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want, tt.args.module)
		})
	}
}