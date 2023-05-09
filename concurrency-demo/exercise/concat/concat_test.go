package concat

import (
	"reflect"
	"sort"
	"testing"
)

func Test_concat(t *testing.T) {
	type args struct {
		dir     string
		workers int8
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "Given a directory, outputs each line of each file in uppercase",
			args: args{
				dir:     "test_files",
				workers: 10,
			},
			want:    []string{"APPLE", "BANANA", "ORANGE", "CAR", "BUS", "TRAIN"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Concat(tt.args.dir, tt.args.workers)
			if err != nil && tt.wantErr != nil {
				t.Errorf("Concat got an unexpected error %v", err)
			}
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat got %v, wanted %v", got, tt.want)
			}
		})
	}
}
