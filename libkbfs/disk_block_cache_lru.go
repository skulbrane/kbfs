package libkbfs

import (
	"time"

	"github.com/keybase/kbfs/kbfsblock"
)

// lruEntry is an entry for sorting LRU times
type lruEntry struct {
	BlockID kbfsblock.ID
	Time    time.Time
}

type blockIDsByTime []lruEntry

func (b blockIDsByTime) Len() int           { return len(b) }
func (b blockIDsByTime) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b blockIDsByTime) Less(i, j int) bool { return b[i].Time.Before(b[j].Time) }

func (b blockIDsByTime) ToBlockIDSlice(numBlocks int) []kbfsblock.ID {
	ids := make([]kbfsblock.ID, 0, numBlocks)
	for _, entry := range b {
		if len(ids) == numBlocks {
			return ids
		}
		ids = append(ids, entry.BlockID)
	}
	return ids
}
