package redis

import (
	"os"
	"testing"

	"github.com/iamajoe/wordless/config"
)

func TestConnect(t *testing.T) {
	c, err := config.Get(os.Getenv)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		host   string
		secret string
	}
	tests := []struct {
		name string
		args args
	}{
		{"runs", args{c.RedisHost, c.RedisSecret}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.host, tt.args.secret, 10)
			if err != nil {
				t.Fatal(err)
			}

			if got == nil {
				t.Errorf("Connect() = %v, want %v", got, nil)
			}

			err = got.Close()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
