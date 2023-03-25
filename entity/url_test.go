package entity

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_newModelUrl(t *testing.T) {
	type args struct {
		id  string
		url string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			return testStruct{
				name: "runs",
				args: args{
					id:  fmt.Sprintf("%d", rand.Intn(100000)),
					url: fmt.Sprintf("%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewModelUrl(tt.args.id, tt.args.url)

			if got.Value != tt.args.url {
				t.Errorf("got.Value = %v, want %v", got.Value, tt.args.url)
			}
		})
	}
}
