package lstoremem

import (
	core "github.com/singyiu/go-threads/core/logstore"
	lstore "github.com/singyiu/go-threads/logstore"
)

// NewLogstore creates an in-memory threadsafe collection of thread logs.
func NewLogstore() core.Logstore {
	return lstore.NewLogstore(
		NewKeyBook(),
		NewAddrBook(),
		NewHeadBook(),
		NewThreadMetadata())
}
