package account

import (
	"github.com/a1emax/youngine/store"
)

type Account struct {
	buffer store.Buffer[Data]
}

func New(buffer store.Buffer[Data]) *Account {
	buffer.Data().Check()
	buffer.Push()

	return &Account{
		buffer: buffer,
	}
}

func (a *Account) Debug() bool {
	return a.buffer.Data().Debug
}

func (a *Account) ToggleDebug() {
	data := a.buffer.Data()
	data.Debug = !data.Debug

	a.buffer.Push()
}
