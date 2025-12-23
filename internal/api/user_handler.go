package api

import (
	"net/http"

	"github.com/coelhoedudev/gobit/internal/jsonutils"
	"github.com/coelhoedudev/gobit/internal/usecase/user"
)

func (api *Api) HandleLogin(w http.ResponseWriter, r *http.Request) {
	panic("NOT implemented")
}

func (api *Api) HandleSignup(w http.ResponseWriter, r *http.Request) {
	if _, mappedErrs, err := jsonutils.DecodeValidJson[*user.CreateUserDTO](r); err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, mappedErrs)
	}
}

func (api *Api) HandleLogout(w http.ResponseWriter, r *http.Request) {
	panic("NOT implemented")
}
