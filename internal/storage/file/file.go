package file

type File struct {
	path string
}

func (f *File) Create(query string) error {
	return nil
}
func (f *File) Read(query string) ([]byte, error) {
	return nil, nil
}
func (f *File) Update(query string) (int, error) {
	return 0, nil
}
func (f *File) Delete(query string) (int, error) {
	return 0, nil
}
