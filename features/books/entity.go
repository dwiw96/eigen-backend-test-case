package books

type Books struct {
	Code   string
	Title  string
	Author string
	Stock  int
}

type RepositoryInterface interface {
	InsertListOfBooks(input []Books) (err error)
}

type ServiceInterface interface {
	InsertListOfBooks(input []Books) (err error)
}
