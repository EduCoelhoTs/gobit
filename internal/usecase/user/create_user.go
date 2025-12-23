package user

import (
	"context"

	_error "github.com/coelhoedudev/gobit/internal/error"
	"github.com/coelhoedudev/gobit/internal/validator"
)

type CreateUserDTO struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req *CreateUserDTO) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", _error.NotBlankMsg)
	eval.CheckField(validator.NotBlank(req.Email), "email", _error.NotBlankMsg)
	eval.CheckField(validator.Matches(req.Email, validator.EmailRx), "email", _error.EmailMaskMsg)
	eval.CheckField(validator.NotBlank(req.Password), "password", _error.NotBlankMsg)
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "field must be bigger than 8 chars")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", _error.NotBlankMsg)
	eval.CheckField(
		validator.MinChars(req.Bio, 8) && validator.MaxChars(req.Bio, 255),
		"bio",
		"this field has to be the length between 8 and 255 caracters",
	)

	return eval
}
