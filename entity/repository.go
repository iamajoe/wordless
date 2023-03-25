package entity

type Repositories interface {
	Close() error
	GetUrl() RepositoryUrl
}
