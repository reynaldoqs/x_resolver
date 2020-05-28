package communityresolverserv

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
	"github.com/reynaldoqs/x_resolver/internal/core/ports"
)

type service struct {
	repo ports.CommunityResolverRepository
}

func NewService(repo ports.CommunityResolverRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Validate(resolver *domain.CommunityResolver) error {
	v := validator.New()
	err := v.Struct(resolver)
	if err != nil {
		var errors string
		for _, err := range err.(validator.ValidationErrors) {
			errors += fmt.Sprintf("error: %v isn't aceptable value for %v \n", err.Value(), err.Field())
		}

		return fmt.Errorf(errors)
	}
	return err
}

func (s *service) Create(resolver *domain.CommunityResolver) error {

	return s.repo.SaveC(resolver)
}
func (s *service) ListResolvers() ([]*domain.CommunityResolver, error) {
	return s.repo.GetAllC()
}
