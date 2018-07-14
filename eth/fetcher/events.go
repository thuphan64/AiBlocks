package fetcher

import "github.com/aiblocksproject/go-aiblocks/core/types"

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
