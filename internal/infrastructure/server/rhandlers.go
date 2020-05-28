package server

import (
	"encoding/json"
	"net/http"

	"github.com/reynaldoqs/x_resolver/internal/core/domain"
	"github.com/reynaldoqs/x_resolver/internal/core/ports"
)

func GetRecharges(service ports.RechargesService) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		recharges, err := service.ListRecharges()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(&recharges)
	}
}

func PostRecharge(service ports.RechargesService) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		//podria ser byte y en el servico cambiarlo a domain.recahrge
		var recharge domain.Recharge

		err := json.NewDecoder(req.Body).Decode(&recharge)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		err = service.Validate(&recharge)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		err = service.Create(&recharge)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		res.WriteHeader(http.StatusNoContent)

	}
}
