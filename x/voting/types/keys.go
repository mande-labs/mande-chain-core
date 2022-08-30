package types

import "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "voting"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_voting"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func VoteBookIndex(voterAddress string, receiverAddress string) string {
	return fmt.Sprintf("%s-%s", voterAddress, receiverAddress)
}
