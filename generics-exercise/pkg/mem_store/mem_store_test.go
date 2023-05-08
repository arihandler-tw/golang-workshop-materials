package mem_store

import (
	"reflect"
	"testing"
)

func TestMemStore_Find(t *testing.T) {
	type args[K comparable] struct {
		id K
	}

	type testCase[K comparable, V any] struct {
		name      string
		s         MemStore[K, V]
		args      args[K]
		want      *V
		wantFound bool
	}
	tests := []testCase[string, string]{
		{
			name: "Find a present value",
			s: MemStore[string, string]{
				store: map[string]string{"id-1": "id-1"},
			},
			args: args[string]{
				id: "id-1",
			},
			want:      GetPtr("id-1"),
			wantFound: true,
		},
		{
			name: "Returns not found for a missing value",
			s: MemStore[string, string]{
				store: map[string]string{"id-1": "id-1"},
			},
			args: args[string]{
				id: "id-2",
			},
			want:      nil,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valueGot, foundGot := tt.s.Find(tt.args.id)
			if !reflect.DeepEqual(valueGot, tt.want) {
				t.Errorf("Find() got = %v, want %v", valueGot, tt.want)
			}
			if foundGot != tt.wantFound {
				t.Errorf("Find() wantFound = %v, want %v", foundGot, tt.wantFound)
			}
		})
	}
}

func TestMemStore_Store(t *testing.T) {
	type args[K comparable, V any] struct {
		id    K
		value V
	}
	type testCase[K comparable, V any] struct {
		name       string
		s          MemStore[K, V]
		args       args[K, V]
		want       *V
		wantErr    error
		finalStore MemStore[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "Store a value is returned",
			s: MemStore[string, string]{
				store: map[string]string{},
			},
			args: args[string, string]{
				id:    "id-1",
				value: "value-1",
			},
			want:    GetPtr("value-1"),
			wantErr: nil,
			finalStore: MemStore[string, string]{
				store: map[string]string{"id-1": "value-1"},
			},
		},
		{
			name: "Overwrites the value if the key is already present",
			s: MemStore[string, string]{
				store: map[string]string{"id-1": "value-1"},
			},
			args: args[string, string]{
				id:    "id-1",
				value: "value-2",
			},
			want:    GetPtr("value-2"),
			wantErr: nil,
			finalStore: MemStore[string, string]{
				store: map[string]string{"id-1": "value-2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Store(tt.args.id, tt.args.value)
			if err != nil && err != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.s, tt.finalStore) {
				t.Errorf("Store() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func GetPtr[T any](t T) *T {
	return &t
}
