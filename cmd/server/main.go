package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	spotifyoauth2 "golang.org/x/oauth2/spotify"
	"google.golang.org/grpc"

	"github.com/gorkagg10/lovify-user-service/database"
	service "github.com/gorkagg10/lovify-user-service/grpc/user-service"
	"github.com/gorkagg10/lovify-user-service/internal/domain/oauth"
	"github.com/gorkagg10/lovify-user-service/internal/domain/profile"
	"github.com/gorkagg10/lovify-user-service/internal/infra/mongodb"
	"github.com/gorkagg10/lovify-user-service/internal/infra/server"
	"github.com/gorkagg10/lovify-user-service/internal/infra/spotify"
)

func main() {
	port := 8082

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	dbClient, err := database.Connect(ctx)
	if err != nil {
		slog.Error("failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			slog.Error("failed to disconnect from database: %v", err)
			os.Exit(1)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error("failed to listen", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("listening", slog.String("port", fmt.Sprintf(":%d", port)))

	userServer := setupUserServer(dbClient)
	srv := SetupGrpcServer(userServer)

	go func() {
		if err = srv.Serve(lis); err != nil {
			slog.Error("failed to serve", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	srv.GracefulStop()
	slog.Info("shutdown completed", slog.String("port", fmt.Sprintf(":%d", port)))
}

func SetupGrpcServer(userServer *server.UserServer) *grpc.Server {
	grpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(grpcServer, userServer)
	return grpcServer
}

func setupUserServer(dbClient *mongo.Client) *server.UserServer {
	userCollection := dbClient.Database("userService").Collection("profiles")

	userRepository := mongodb.NewUserRepository(userCollection)
	oAuthRepository := spotify.NewOAuthRepository(
		&oauth2.Config{
			ClientID:     "f4ed25e807ab4b74b981cd606a75699b",
			ClientSecret: "4b8515bf00ed4f67bbcd9a77d7486bdb",
			Endpoint:     spotifyoauth2.Endpoint,
			RedirectURL:  "http://127.0.0.1:8082/callback/spotify",
			Scopes:       []string{"user-read-email", "user-read-recently-played", "user-top-read"},
		},
	)

	profileManager := profile.NewManager(userRepository)
	oAuthService := oauth.NewService(oAuthRepository)
	return server.NewUserServer(profileManager, oAuthService)
}
