package domain

import "github.com/iamajoe/wordless/entity"

func CreateUrl(url string, isMini bool, urlRepo entity.RepositoryUrl) (string, error) {
	// TODO: need to actually set the mini
	return urlRepo.Create(url)
}

func FetchId(id string, urlRepo entity.RepositoryUrl) (string, error) {
  urls, err := urlRepo.GetByIDs([]string{id})
  if len(urls) == 0 {
    return "", err
  }

  return urls[0].Value, nil
}
