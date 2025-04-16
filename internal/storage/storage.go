package storage

type Storage interface {
	Create(query string) error
	Read(query string) ([]byte, error)
	Update(query string) (int, error)
	Delete(query string) (int, error)
}
