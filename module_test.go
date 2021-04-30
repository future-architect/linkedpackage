package linkedpackage

import (
	"reflect"
	"testing"
)

func TestUniqueModules(t *testing.T) {
	type args struct {
		modules []Module
	}
	tests := []struct {
		name string
		args args
		want []Module
	}{
		{
			name: "unique",
			args: args{
				modules: []Module{
					{
						Lang: "js",
						Name: "trim",
						Path: "/node_modules/trim",
					},
					{
						Lang: "js",
						Name: "trim",
						Path: "/node_modules/trim",
					},
					{
						Lang: "js",
						Name: "trim",
						Path: "/node_modules/trim",
					},
				},
			},
			want: []Module{
				{
					Lang: "js",
					Name: "trim",
					Path: "/node_modules/trim",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueModules(tt.args.modules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueModules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupingModuleByLicense(t *testing.T) {
	type args struct {
		modules []Module
	}
	tests := []struct {
		name string
		args args
		want []GroupedModule
	}{
		{
			name: "group same author and license",
			args: args{
				modules: []Module{
					{
						Lang: "js",
						Name: "sample1",
						Author: "user1",
						LicenseName: "MIT",
					},
					{
						Lang: "js",
						Name: "sample2",
						Author: "user1",
						LicenseName: "MIT",
					},
				},
			},
			want: []GroupedModule{
				{
					Author: "user1",
					License: "MIT",
					Modules: []Module{
						{
							Lang: "js",
							Name: "sample1",
							Author: "user1",
							LicenseName: "MIT",
						},
						{
							Lang: "js",
							Name: "sample2",
							Author: "user1",
							LicenseName: "MIT",
						},
					},
				},
			},
		},
		{
			name: "user is different, they become separated groups",
			args: args{
				modules: []Module{
					{
						Lang: "js",
						Name: "sample1",
						Author: "user1",
						LicenseName: "MIT",
					},
					{
						Lang: "js",
						Name: "sample2",
						Author: "user2",
						LicenseName: "MIT",
					},
				},
			},
			want: []GroupedModule{
				{
					Author: "user1",
					License: "MIT",
					Modules: []Module{
						{
							Lang: "js",
							Name: "sample1",
							Author: "user1",
							LicenseName: "MIT",
						},
					},
				},
				{
					Author: "user2",
					License: "MIT",
					Modules: []Module{
						{
							Lang: "js",
							Name: "sample2",
							Author: "user2",
							LicenseName: "MIT",
						},
					},
				},
			},
		},
		{
			name: "license is different, they become separated groups",
			args: args{
				modules: []Module{
					{
						Lang: "js",
						Name: "sample1",
						Author: "user1",
						LicenseName: "MIT",
					},
					{
						Lang: "js",
						Name: "sample2",
						Author: "user1",
						LicenseName: "BSD",
					},
				},
			},
			want: []GroupedModule{
				{
					Author: "user1",
					License: "MIT",
					Modules: []Module{
						{
							Lang: "js",
							Name: "sample1",
							Author: "user1",
							LicenseName: "MIT",
						},
					},
				},
				{
					Author: "user1",
					License: "BSD",
					Modules: []Module{
						{
							Lang: "js",
							Name: "sample2",
							Author: "user1",
							LicenseName: "BSD",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupingModulesByLicense(tt.args.modules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupingModulesByLicense() = %v, want %v", got, tt.want)
			}
		})
	}
}