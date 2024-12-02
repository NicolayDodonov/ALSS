package fileSt

import (
	"os"
	"strconv"
)

type FileSt struct {
	pathAge   string
	pathTrain string
}

func New(pathAge, pathTrain string) *FileSt {
	return &FileSt{
		pathAge,
		pathTrain,
	}
}

func (fs *FileSt) WorldAgeSave(year int) error {
	//открыть файл
	file, err := os.OpenFile(fs.pathAge, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//записать данные
	_, err = file.WriteString(strconv.Itoa(year) + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileSt) TrainGenSave(data []string) error {
	//открыть файл
	file, err := os.OpenFile(fs.pathAge, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//записать данные
	for _, str := range data {
		_, err = file.WriteString(str)
		if err != nil {
			return err
		}
	}

	return nil
}
