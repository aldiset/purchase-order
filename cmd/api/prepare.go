package main

import (
	"time"

	userusecase "purchase-order/business/user"
	userhandler "purchase-order/cmd/api/controller/user"
	"purchase-order/cmd/api/middlewares"
	usermodule "purchase-order/modules/database/user"

	"purchase-order/cmd/api/routes"
)

func prepareService(handler *routes.HandlerConfig, depends *dependencies) {
	// Persiapan proses authenticator
	jwtMid := middlewares.NewAuthenticator(handler.Config.JWTSecretKey)

	// Persiapan repository, business dan handler
	userRepo := usermodule.NewRepository(depends.db)
	userUseCase := userusecase.NewService(userRepo)
	userUseCase.SetJWTConfig(
		handler.Config.JWTSecretKey,
		time.Duration(handler.Config.JWTExpiredTime)*time.Minute,
	)
	// paymentUseCase := paymentusecase.NewService(nil, nil)
	// productUseCase := productusecase.NewService(repo, userRepo, paymentUseCase)
	// productUseCase := productusecase.NewService(nil, nil)

	// Controller
	handler.User = userhandler.NewController(userUseCase, jwtMid)
}
