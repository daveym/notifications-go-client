package handler

// IBaseHandler - Executes http request to gov notify interface.
type IBaseHandler interface {
	handle()
	SetTokenCreator()
}
