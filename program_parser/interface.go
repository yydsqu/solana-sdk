package program_parser

import (
	binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type Instruction interface {
	TypeID() binary.TypeID
	SetAccounts(accounts solana.PublicKeySlice) (err error)
	NewInstance() Instruction
}

type Parser interface {
	ProgramID() solana.PublicKey
	InstructionIdType() binary.TypeIDEncoding
	Instructions() map[binary.TypeID]Instruction
	Process(instruction any, tx Transaction) error
}

type Transaction interface {
	Raw() bool
	GetSlot() uint64
	Signature() solana.Signature
	FeePayer() solana.PublicKey
	GetAccount(index int) (solana.PublicKey, error)
	GetPublicKeys(indexes []byte) (solana.PublicKeySlice, error)
	GetAccountMetas(indexes []byte) (solana.AccountMetaSlice, error)
	GetFinalSOLBalanceAmount(account solana.PublicKey) uint64
	GetFinalTokenBalanceAmount(account solana.PublicKey) uint64
}
