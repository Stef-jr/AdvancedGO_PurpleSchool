package verify

import (
	"net/http"
	"validation/configs"
)

type VerifyHandler struct {
	*configs.Config
}

type DepsVerify struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, ver DepsVerify) {
	handler := &VerifyHandler{
		Config: ver.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (handler *VerifyHandler) Send() http.HandlerFunc{
	return nil
}

func (handler *VerifyHandler) Verify() http.HandlerFunc{
	return nil
}