//go:generate mockgen -source auth.go -destination ../testdata/mock_service/auth_gen.go
package service

import "v1_gorm_direct/graph/model"

type AuthService interface {
	GenerateToken(useInfo model.User) (string, error)
	// ValidateToken(tokenString string) error
}
