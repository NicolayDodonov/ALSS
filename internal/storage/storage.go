package storage

type Storage interface {
	WorldAgeSave(int) error
	TrainGenSave([]string) error
}
