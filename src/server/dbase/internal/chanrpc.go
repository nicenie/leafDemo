package internal

func init() {
	skeleton.RegisterChanRPC("login", rpcProcLogin)
}

func rpcProcLogin(args []interface{}) {
}
