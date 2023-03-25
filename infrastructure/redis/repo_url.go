package redis

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/iamajoe/wordless/entity"
)

// ------------------------------
// model

// ------------------------------
// repository

type repositoryUrl struct {
	db        *DB
}

func (repo *repositoryUrl) GetByIDs(ids []string) ([]entity.Url, error) {
	if len(ids) == 0 {
		return []entity.Url{}, nil
	}

	// TODO: probably the context could come from outside so it is cancellable
	ctx := context.Background()

	urls := []entity.Url{}
	for _, id := range ids {
		val, err := repo.db.db.Get(ctx, id).Result()
		if err != nil {
			return urls, nil
		}

		url := entity.NewModelUrl(id, val)
		urls = append(urls, url)
	}

	return urls, nil
}

func (repo *repositoryUrl) Create(url string) (string, error) {
	// TODO: generated id should be different
	id := fmt.Sprintf("%d", rand.Intn(100000))

	// TODO: probably the context could come from outside so it is cancellable
	ctx := context.Background()
	err := repo.db.db.Set(ctx, id, url, 0).Err()
	if err != nil {
		return "", err
	}

	return id, err
}

func createRepositoryUrl(db *DB) (entity.RepositoryUrl, error) {
	repo := repositoryUrl{db}
	return &repo, nil
}
