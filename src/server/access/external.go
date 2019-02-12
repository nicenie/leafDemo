package access

import (
	"server/access/internal"
)

var (
	// Module access module
	Module = new(internal.Module)
	// ChanRPC access chanrpc
	ChanRPC = internal.ChanRPC
)
