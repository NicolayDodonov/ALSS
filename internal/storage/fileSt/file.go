package fileSt

type FileSt struct {
	path string
}

func New(path string) *FileSt {
	return &FileSt{
		path,
	}
}

func (fs *FileSt) WorldAgeSave(year int) error {
	return nil
}

func (fs *FileSt) TrainGenSave(data []string) error {
	return nil
}
