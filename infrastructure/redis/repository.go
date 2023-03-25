package redis

import (
	"errors"
	"fmt"

	"github.com/iamajoe/wordless/entity"
)

type repositories struct {
	db  *DB
	url entity.RepositoryUrl
}

func (r *repositories) GetUrl() entity.RepositoryUrl {
	return r.url
}

func (r *repositories) Close() error {
	if r.db != nil {
		err := r.db.Close()
		if err != nil {
			return err
		}

		r.db = nil
	}

	return nil
}

func InitRepos(host string, secret string) (repos entity.Repositories, err error) {
  // connect url db to the right index
	urlDbIndex := uint(1)
	db, err := Connect(host, secret, urlDbIndex)
	if err != nil {
		err = errors.New(fmt.Sprintf("error initializing db redis: %v", err))
		return repos, err
	}
	url, err := createRepositoryUrl(db)
	if err != nil {
		return repos, err
	}

	return &repositories{db, url}, nil
}
