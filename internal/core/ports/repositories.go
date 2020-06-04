package ports

import "github.com/reynaldoqs/x_resolver/internal/core/domain"

type CommunityResolverRepository interface {
	GetAllC() ([]*domain.CommunityResolver, error)
	SaveC(resolver *domain.CommunityResolver) error
	UpdateC(id string, resolver *domain.CommunityResolver) error
	GetOneC(id string) (*domain.CommunityResolver, error)
	RemoveC(id string) error
}

// CommunityRechargeRepository it needs to be a real time data for users
type CommunityRechargeRepository interface {
	Store(recharge *domain.CommunityRecharge) error
}

type RechargesRespository interface {
	GetAllR() ([]*domain.Recharge, error)
	SaveR(recharge *domain.Recharge) error
	//Update()
}

type FarmersRepository interface {
	GetAllFarmers() ([]*domain.Farmer, error)
}
