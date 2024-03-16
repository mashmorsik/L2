package main

import "testing"

func TestSort(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "default_sort", args: args{filePath:}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Sort(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
