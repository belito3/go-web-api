package impl

import (
	"github.com/belito3/go-api-codebase/app/repository"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewAccountImpl)
	_ = container.Provide(func(m *AccountImpl) repository.IAccount { return m })
	return nil
}