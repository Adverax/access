package access

import (
	"context"
	"sync"
)

type Storage interface {
	Load() ([]byte, error)
	Save(data []byte) error
}

type GetterSetterMarshaller interface {
	GetterSetter
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type StorageGetterSetter struct {
	rw      GetterSetterMarshaller
	storage Storage
	loaded  bool
	sync.Mutex
}

func (that *StorageGetterSetter) GetProperty(
	ctx context.Context,
	name string,
) (interface{}, error) {
	err := that.ensureLoaded()
	if err != nil {
		return nil, err
	}
	return that.rw.GetProperty(ctx, name)
}

func (that *StorageGetterSetter) SetProperty(
	ctx context.Context,
	name string,
	value interface{},
) error {
	err := that.ensureLoaded()
	if err != nil {
		return err
	}

	err = that.rw.SetProperty(ctx, name, value)
	if err != nil {
		return err
	}

	return that.save()
}

func (that *StorageGetterSetter) ensureLoaded() error {
	that.Lock()
	defer that.Unlock()

	if that.loaded {
		return nil
	}
	err := that.load()
	if err != nil {
		return err
	}
	that.loaded = true
	return nil
}

func (that *StorageGetterSetter) load() error {
	data, err := that.storage.Load()
	if err != nil {
		return err
	}
	return that.rw.Unmarshal(data)
}

func (that *StorageGetterSetter) save() error {
	that.Lock()
	defer that.Unlock()

	data, err := that.rw.Marshal()
	if err != nil {
		return err
	}
	return that.storage.Save(data)
}

func NewStorageGetterSetter(
	rw GetterSetterMarshaller,
	storage Storage,
) GetterSetter {
	return &StorageGetterSetter{
		rw:      rw,
		storage: storage,
	}
}
