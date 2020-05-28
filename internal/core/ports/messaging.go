package ports

import "github.com/reynaldoqs/x_resolver/internal/core/domain"

type CloudMessaging interface {
	RechargeNotify(recharge *domain.Recharge, resolvers []*domain.CommunityResolver) error
}
