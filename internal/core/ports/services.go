package ports

import "github.com/reynaldoqs/x_resolver/internal/core/domain"

type CommunityResolverService interface {
	Validate(resolver *domain.CommunityResolver) error
	ListResolvers() ([]*domain.CommunityResolver, error)
	Create(resolver *domain.CommunityResolver) error
	Update(id string, resolver *domain.CommunityResolver) error
	//make a specific method for add recharges
	GetOne(id string) (*domain.CommunityResolver, error)
	Remove(id string) error
}

type RechargesService interface {
	Validate(recharge *domain.Recharge) error
	Create(recharge *domain.Recharge) error
	ListRecharges() ([]*domain.Recharge, error)
}
