package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/reynaldoqs/x_resolver/internal/core/services/communityresolverserv"
	"github.com/reynaldoqs/x_resolver/internal/core/services/rechargesserv"
	firebasemsg "github.com/reynaldoqs/x_resolver/internal/infrastructure/cloudmessenger/firebase"
	"github.com/reynaldoqs/x_resolver/internal/infrastructure/repositories/firestore"
	"github.com/reynaldoqs/x_resolver/internal/infrastructure/repositories/mongo"
)

func RegisterRouter() {
	chiDispatcher := chi.NewRouter()
	chiDispatcher.Use(middleware.RequestID)
	chiDispatcher.Use(middleware.RealIP)
	chiDispatcher.Use(middleware.Logger)
	chiDispatcher.Use(middleware.Recoverer)

	mongoClient := mongo.NewMongoClient("mongodb://localhost:27017", 30)
	googleApp := firebasemsg.NewFirebaseApp("./gu-project.json")

	repoRecharges := mongo.NewRechargesRepository(mongoClient, "project-x")
	cloudMsgr := firebasemsg.NewCloudMessaging(googleApp)
	repoCommRecharge := firestore.NewCommunityRechargeRepository(googleApp)
	repoFarmers := firestore.NewFarmersRepository(googleApp)
	repoCommResolver := mongo.NewCommunityResolverRepository(mongoClient, "project-x")

	servRecharges := rechargesserv.NewService(repoRecharges, cloudMsgr, repoCommRecharge, repoCommResolver, repoFarmers)

	getReachargesEndpoint := GetRecharges(servRecharges)
	postRechargeEndpoint := PostRecharge(servRecharges)

	chiDispatcher.Get("/recharges", getReachargesEndpoint)
	chiDispatcher.Post("/recharges", postRechargeEndpoint)

	repoCResolvers := mongo.NewCommunityResolverRepository(mongoClient, "project-x")
	servCResolvers := communityresolverserv.NewService(repoCResolvers)

	getCResolversEndpoint := GetCommunityResolvers(servCResolvers)
	postCResolverEndpoint := PostCommunityResolver(servCResolvers)
	putCResolversEndpoint := PutCommunityResolver(servCResolvers)
	getOneCResolversEndpoint := GetOneCommunityResolver(servCResolvers)
	deleteCResolverEndpoint := DeleteCommunityResolver(servCResolvers)

	chiDispatcher.Get("/cresolvers", getCResolversEndpoint)
	chiDispatcher.Post("/cresolvers", postCResolverEndpoint)
	chiDispatcher.Put("/cresolvers/{id}", putCResolversEndpoint)
	chiDispatcher.Get("/cresolvers/{id}", getOneCResolversEndpoint)
	chiDispatcher.Delete("/cresolvers/{id}", deleteCResolverEndpoint)

	const port string = ":8080"
	fmt.Printf("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, chiDispatcher)
}
