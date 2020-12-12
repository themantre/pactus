package store

import (
	"github.com/zarbchain/zarb-go/account"
	"github.com/zarbchain/zarb-go/block"
	"github.com/zarbchain/zarb-go/crypto"
	"github.com/zarbchain/zarb-go/tx"
	"github.com/zarbchain/zarb-go/validator"
)

type StoreReader interface {
	BlockByHeight(height int) (*block.Block, error)
	BlockByHash(hash crypto.Hash) (*block.Block, int, error)
	BlockHeight(hash crypto.Hash) (int, error)
	Transaction(hash crypto.Hash) (*tx.CommittedTx, error)
	HasAccount(crypto.Address) bool
	Account(addr crypto.Address) (*account.Account, error)
	TotalAccounts() int
	HasValidator(crypto.Address) bool
	Validator(addr crypto.Address) (*validator.Validator, error)
	TotalValidators() int
}
