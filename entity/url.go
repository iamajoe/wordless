package entity

type Url struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func NewModelUrl(id string, value string) Url {
	return Url{id, value}
}

type RepositoryUrl interface {
	GetByIDs(ids []string) ([]Url, error)
	Create(value string) (string, error)
}
