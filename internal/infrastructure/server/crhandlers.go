package server

import (
	"encoding/json"
	"net/http"

	"github.com/reynaldoqs/x_resolver/internal/core/domain"
	"github.com/reynaldoqs/x_resolver/internal/core/ports"
)

func GetCommunityResolvers(service ports.CommunityResolverService) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		resolvers, err := service.ListResolvers()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(&resolvers)
	}
}

func PostCommunityResolver(service ports.CommunityResolverService) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		var resolver domain.CommunityResolver

		err := json.NewDecoder(req.Body).Decode(&resolver)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		err = service.Validate(&resolver)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		err = service.Create(&resolver)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		res.WriteHeader(http.StatusNoContent)

	}
}
