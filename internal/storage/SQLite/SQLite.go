package SQLite

type SQLite struct {
	path string
}

func (s *SQLite) Create(query string) error {
	return nil
}
func (s *SQLite) Read(query string) ([]byte, error) {
	return nil, nil
}
func (s *SQLite) Update(query string) (int, error) {
	return 0, nil
}
func (s *SQLite) Delete(query string) (int, error) {
	return 0, nil
}
