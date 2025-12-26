package api

import (
	"errors"
	"fmt"
	"net/http"

	_error "github.com/coelhoedudev/gobit/internal/error"
	"github.com/coelhoedudev/gobit/internal/jsonutils"
	"github.com/coelhoedudev/gobit/internal/service"
	"github.com/coelhoedudev/gobit/internal/usecase/user"
)

func (api *Api) HandleLogin(w http.ResponseWriter, r *http.Request) {
	panic("NOT implemented")
}

func (api *Api) HandleSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bateu")
	data, mappedErrs, err := jsonutils.DecodeValidJson[user.CreateUserDTO](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, mappedErrs)
		api.Logger.Error(err.Error())
		return
	}

	id, err := api.UserService.Create(r.Context(), &data)
	if err != nil {
		if errors.Is(err, service.ErrDuplicatedEmailOrPassword) {
			_ = jsonutils.EncodeJson(w, r, int(http.StatusBadRequest), map[string]string{
				"error": "invalid email or password",
			})
			return
		}

		http.Error(w, _error.ServerInternalErrorMsg, http.StatusInternalServerError)
		api.Logger.Error(err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]string{
		"id": id.String(),
	})
}

func (api *Api) HandleLogout(w http.ResponseWriter, r *http.Request) {
	panic("NOT implemented")
}
