package dbase

import (
	"server/dbase/internal"
)

var (
	// Module leaf db module
	Module = new(internal.Module)
	// ChanRPC db chanrpc
	ChanRPC = internal.ChanRPC
)
