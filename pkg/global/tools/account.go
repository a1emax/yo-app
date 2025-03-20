package tools

import (
	"github.com/a1emax/youngine/x/scope"

	"yo-app/pkg/domain/account"
)

var Account *account.Account

func initAccount(lc scope.Lifecycle) {
	Account = account.New(StoreBuffer)
}
