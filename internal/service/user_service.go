package service

import (
	"context"
	"errors"

	_crypto "github.com/coelhoedudev/gobit/internal/crypto"
	"github.com/coelhoedudev/gobit/internal/store/pgstore"
	"github.com/coelhoedudev/gobit/internal/usecase/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	queries *pgstore.Queries
	pool    *pgxpool.Pool
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var ErrDuplicatedEmailOrPassword = errors.New("username or email already exists")

func (s *UserService) Create(ctx context.Context, user user.CreateUserDTO) (uuid.UUID, error) {

	hash, err := _crypto.GenerateHashFromPassword(user.Password)
	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: hash,
		Bio:          user.Bio,
	}

	id, err := s.queries.CreateUser(ctx, args)
	if err != nil {
		//verificar o tipo de erro(verificar o tipo ex: 2305)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrPassword
		}

		return uuid.UUID{}, err
	}

	return id, nil
}
