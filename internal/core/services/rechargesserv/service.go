package rechargesserv

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
	"github.com/reynaldoqs/x_resolver/internal/core/ports"
)

type service struct {
	rechargeRepo   ports.RechargesRespository
	cloudMsg       ports.CloudMessaging
	cRechargesRepo ports.CommunityRechargeRepository
	cResolversRepo ports.CommunityResolverRepository
	farmersRepo    ports.FarmersRepository
}

func NewService(
	rr ports.RechargesRespository,
	cm ports.CloudMessaging,
	crecr ports.CommunityRechargeRepository,
	cresr ports.CommunityResolverRepository,
	farRep ports.FarmersRepository) *service {
	return &service{
		rechargeRepo:   rr,
		cloudMsg:       cm,
		cRechargesRepo: crecr,
		cResolversRepo: cresr,
		farmersRepo:    farRep,
	}
}

func (s *service) Validate(recharge *domain.Recharge) error {
	v := validator.New()
	err := v.Struct(recharge)
	if err != nil {
		var errors string
		for _, err := range err.(validator.ValidationErrors) {
			errors += fmt.Sprintf("error: %v isn't aceptable value for %v \n", err.Value(), err.Field())
		}

		return fmt.Errorf(errors)
	}
	return err
}

func (s *service) Create(recharge *domain.Recharge) error {
	recharge.CreatedAt = time.Now()

	s.rechargeRepo.SaveR(recharge)

	// add recharge to our raltime data base for users
	/*cRecharge := domain.CommunityRecharge{
		IDRecharge: recharge.ID,
		Mount:      recharge.Mount,
		State:      1,
		Company:    recharge.Company,
		ExecCode:   "*105#",
		CreatedAt:  recharge.CreatedAt,
	}
	s.cRechargesRepo.Store(&cRecharge)

	// send notification to resolvers
	resolvers, err := s.cResolversRepo.GetAllC()
	if err != nil {
		err = errors.Wrap(err, "service.Create")
	}
	s.cloudMsg.RechargeNotify(recharge, resolvers)
	*/
	//for FARMER
	farmers, err := s.farmersRepo.GetAllFarmers()
	if err != nil {
		err = errors.Wrap(err, "service.Create")
	}

	for _, v := range farmers {
		fmt.Println("Aqui estamos")
		s.cloudMsg.FarmerNotify(v, nil)
	}
	return nil
}
func (s *service) ListRecharges() ([]*domain.Recharge, error) {
	return s.rechargeRepo.GetAllR()
}
