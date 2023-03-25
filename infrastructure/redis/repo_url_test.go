package redis

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/iamajoe/wordless/config"
	"github.com/iamajoe/wordless/entity"
)

func TestRepositoryUrl_GetByIDs(t *testing.T) {
	c, err := config.Get(os.Getenv)
	if err != nil {
		t.Fatal(err)
	}

	db, err := Connect(c.RedisHost, c.RedisSecret, 10)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.FlushAll(context.Background())
	defer db.Close()

	repo, err := createRepositoryUrl(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		url entity.Url
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			urlValue := fmt.Sprintf("tmp_url_%d", rand.Intn(100000))
			id, err := repo.Create(urlValue)
			if err != nil {
				t.Fatal(err)
			}

			url := entity.NewModelUrl(id, urlValue)

			return testStruct{
				name: "runs",
				args: args{url},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByIDs([]string{tt.args.url.ID})
			if err != nil {
				t.Fatal(err)
				return
			}

			if len(got) == 0 {
				t.Errorf("RepositoryUser.GetByIDs() = %v, want %v", len(got), 1)
			}

			if got[0].Value != tt.args.url.Value {
				t.Errorf("RepositoryUser.GetByIDs() = %v, want %v", got[0].Value, tt.args.url.Value)
			}
		})
	}
}

func TestRepositoryUrl_Create(t *testing.T) {
	c, err := config.Get(os.Getenv)
	if err != nil {
		t.Fatal(err)
	}

	db, err := Connect(c.RedisHost, c.RedisSecret, 11)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.FlushAll(context.Background())
	defer db.Close()

	repo, err := createRepositoryUrl(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		url entity.Url
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			urlValue := fmt.Sprintf("tmp_url_%d", rand.Intn(100000))
			url := entity.NewModelUrl("", urlValue)

			return testStruct{
				name: "runs",
				args: args{url},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Create(tt.args.url.Value)
			if err != nil {
				t.Fatal(err)
				return
			}

			if len(got) == 0 {
				t.Errorf("RepositoryUser.Create() = %v, want %v", got, true)
			}

			// check if url is in
			url, err := db.db.Get(context.Background(), got).Result()
			if err != nil {
				t.Fatal(err)
			}

			if url != tt.args.url.Value {
				t.Errorf("url = %v, want %v", url, tt.args.url.Value)
			}
		})
	}
}

func Test_createRepositoryUrl(t *testing.T) {
	c, err := config.Get(os.Getenv)
	if err != nil {
		t.Fatal(err)
	}

	db, err := Connect(c.RedisHost, c.RedisSecret, 12)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.FlushAll(context.Background())
	defer db.Close()

	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryUrl(db)
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryUrl() = %v, want %v", got, nil)
			}
		})
	}
}
