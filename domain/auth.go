package domain

import (
	"context"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
)

type AuthService interface{
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}