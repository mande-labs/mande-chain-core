package types

import "fmt"

import "encoding/binary"

var _ binary.ByteOrder

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

	// VoteBookKeyPrefix is the prefix to retrieve all VoteBook
	VoteBookKeyPrefix = "VoteBook/value/"

	// AggregateVoteCountKeyPrefix is the prefix to retrieve all AggregateVoteCount
	AggregateVoteCountKeyPrefix = "AggregateVoteCount/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func VoteBookIndex(voterAddress string, receiverAddress string) string {
	return fmt.Sprintf("%s-%s", voterAddress, receiverAddress)
}

// VoteBookKey returns the store key to retrieve a VoteBook from the index fields
func VoteBookKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// AggregateVoteCountKey returns the store key to retrieve a AggregateVoteCount from the index fields
func AggregateVoteCountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
