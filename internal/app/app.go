package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/mamoru777/userservice2/internal/config"
	"github.com/mamoru777/userservice2/internal/databaseinit"
	"github.com/mamoru777/userservice2/internal/middleware"
	"github.com/mamoru777/userservice2/internal/mylogger"
	"github.com/mamoru777/userservice2/internal/repositories/userrepository"
	"github.com/mamoru777/userservice2/internal/service"
	"gitlab.com/mediasoft-internship/internship/mamoru777/foundation/xrequestidmiddleware"

	gateway_api "github.com/mamoru777/userservice2/pkg/gateway-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(dataBaseConfig config.DataBaseConfig, grpcServerConfig config.GrpcServerConfig) error {

	db, err := databaseinit.InitSqlxDB(dataBaseConfig)
	if err != nil {
		mylogger.Logger.Fatal("Не удалось инициализировать базу данных", err)
	}
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(
		xrequestidmiddleware.NewReqInterceptor().RequestIDInterceptor,
		middleware.NewAuthInterceptor().JWTInterceptor),
	)
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go runGrpcServer(grpcServerConfig, s, db)
	go runHTTPServer(ctx, grpcServerConfig, mux)

	gracefullyShutdown(s, cancel)
	return nil
}

func runGrpcServer(grpcServerConfig config.GrpcServerConfig, s *grpc.Server, db *sqlx.DB) {
	serv := service.New(userrepository.New(db))
	gateway_api.RegisterUserInfoServiceServer(s, serv)
	l, err := net.Listen("tcp", grpcServerConfig.GRPCAddr)
	if err != nil {
		mylogger.Logger.Fatalf("не удалось прослушать tcp %s, %v", grpcServerConfig.GRPCAddr, err)
	}
	mylogger.Logger.Printf("запуск grpc сервера по адресу: %s", grpcServerConfig.GRPCAddr) //config.GRPCAddr)
	if err := s.Serve(l); err != nil {
		mylogger.Logger.Fatalf("ошибка сервиса grpc сервера %v", err)
	}
}

func runHTTPServer(
	ctx context.Context, cfg config.GrpcServerConfig, mux *runtime.ServeMux,
) {
	err := gateway_api.RegisterUserInfoServiceHandlerFromEndpoint(
		ctx,
		mux,
		"0.0.0.0"+cfg.GRPCAddr,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		mylogger.Logger.Fatal(err)
	}
	mylogger.Logger.Printf("запуск http сервера по адресу: %s", cfg.HTTPAddr)
	if err := http.ListenAndServe(cfg.HTTPAddr, mux); err != nil {
		mylogger.Logger.Fatalf("ошибка сервиса http сервера %v", err)
	}

}

func gracefullyShutdown(s *grpc.Server, cancel context.CancelFunc) {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Получен сигнал выключения:", sig, "Graceful shutdown выполнен")
	log.Println(errorMessage)
	s.GracefulStop()
	cancel()
}
