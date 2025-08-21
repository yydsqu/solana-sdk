package program_parser

import (
	binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type Instruction interface {
	TypeID() binary.TypeID
	SetAccounts(accounts solana.PublicKeySlice) (err error)
	Copy() Instruction
}
