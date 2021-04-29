package impl

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	//_ = container.Provide(NewAccountImpl)
	//_ = container.Provide(func(m *AccountImpl) repository.IAccount { return m })
	_ = container.Provide(NewStore)
	_ = container.Provide(func(m *Store) IStore { return m })
	return nil
}