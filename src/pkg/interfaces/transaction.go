package interfaces

import (
	"DigitalAsset/src/pkg/transaction"
)

type Transaction interface {
	New(operation string,
		asset transaction.Asset,
		metadata transaction.Metadata,
		inputs []transaction.Input,
		outputs []transaction.Output,
	) (*transaction.Transaction, error)
	NewCreateTransaction(
		asset transaction.Asset,
		metadata transaction.Metadata,
		outputs []transaction.Output,
		issuers []string,
	) (*transaction.Transaction, error)
	//NewTransferTransaction() (*transaction.Transaction, error)

	Sign(keyPairs []transaction.KeyPair) error
}
