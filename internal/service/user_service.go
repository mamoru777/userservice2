package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mamoru777/userservice2/internal/mylogger"
	"github.com/mamoru777/userservice2/internal/repositories/userrepository"
	chatserviceapi "gitlab.com/mediasoft-internship/internship/mamoru777/chatservice/pkg/gateway-api"

	gatewayapi "github.com/mamoru777/userservice2/pkg/gateway-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	gatewayapi.UnimplementedUserInfoServiceServer
	userrep userrepository.IUserRepository
}

func New(userrep userrepository.IUserRepository) *UserService {
	return &UserService{
		userrep: userrep,
	}
}

func (us *UserService) SignUpUserInfo(ctx context.Context, request *gatewayapi.SignUpUserInfoRequest) (*gatewayapi.SignUpUserInfoResponse, error) {
	xRequestId := ctx.Value("x_request_id")
	xRequestIdString, ok := xRequestId.(string)
	if !ok {
		err := errors.New("Не получилось извлечь xRequestId из контекста")
		mylogger.Logger.Error(err)
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}

	mylogger.Logger.Println("Запрос № ", xRequestIdString, " ", "Использована функция SignUpUserInfo")
	if len(request.GetUser().Fio) < 4 {
		err := errors.New("ФИО должно состоять минимум из 4 символов")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	mylogger.Logger.Println("Запрос № ", xRequestIdString, " ", "Получено значение фамилии")
	if request.GetUser().Post == "" {
		err := errors.New("Поле должность обязательно к заполнению")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	if request.GetUser().Department == "" {
		err := errors.New("Поле отдел обязательно к заполнению")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	mylogger.Logger.Println("Запрос № ", xRequestIdString, " ", "Использована функция получения user id из контекста")
	userId := ctx.Value("user_id")
	mylogger.Logger.Println("Запрос № ", xRequestIdString, " ", "Получилось извлечь user id из контекста")
	mylogger.Logger.Println("Запрос № ", xRequestIdString, " ", userId)
	userIdString, ok := userId.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx user id в string")
		return &gatewayapi.SignUpUserInfoResponse{}, nil
	}
	mylogger.Logger.Println("Запрос № ", xRequestIdString, " ", userIdString)
	userUuid, err := StringToUuuid(userIdString)
	if err != nil {
		err := errors.New("Не удалось преобразовать строку в uuid")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	user := userrepository.User{
		Id:         userUuid,
		Fio:        request.GetUser().Fio,
		Post:       request.GetUser().Post,
		Department: request.GetUser().Department,
	}
	err = us.userrep.Create(ctx, &user)
	if err != nil {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось создать пользователя в бд", err)
		err := errors.New("Не удалось создать пользователя в бд")
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	accessToken := ctx.Value("access_token")
	refreshToken := ctx.Value("refresh_token")
	accessTokenString, ok := accessToken.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx accessToken в string")
		return &gatewayapi.SignUpUserInfoResponse{}, nil
	}
	refreshTokenString, ok := refreshToken.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx refreshToken в string")
		return &gatewayapi.SignUpUserInfoResponse{}, nil
	}
	conn, err := grpc.Dial("localhost:13995", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось подключиться к grpc серверу chat service", err)
		err = errors.New("Не удалось подключиться к grpc серверу chat service")
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	chatservice := chatserviceapi.NewChatServiceClient(conn)
	err = CreateChats(ctx, chatservice, &chatserviceapi.CreateChatsRequest{Userid: userIdString})
	if err != nil {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.SignUpUserInfoResponse{}, err
	}
	return &gatewayapi.SignUpUserInfoResponse{AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}

func (us *UserService) GetUserInfo(ctx context.Context, request *gatewayapi.GetUserInfoRequest) (*gatewayapi.GetUserInfoResponse, error) {
	xRequestId := ctx.Value("x_request_id")
	xRequestIdString, ok := xRequestId.(string)
	if !ok {
		err := errors.New("Не получилось извлечь xRequestId из контекста")
		mylogger.Logger.Error(err)
		return &gatewayapi.GetUserInfoResponse{}, err
	}

	userId := ctx.Value("user_id")
	userIdString, ok := userId.(string)
	if !ok {
		err := errors.New("Не получилось извлечь user id из контекста")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.GetUserInfoResponse{}, err
	}
	userUuid, err := StringToUuuid(userIdString)
	if err != nil {
		err := errors.New("Не удалось преобразовать строку в uuid")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.GetUserInfoResponse{}, err
	}
	user, err := us.userrep.Get(ctx, userUuid)
	if err != nil {
		err := errors.New("Не удалось получить пользователя из бд")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.GetUserInfoResponse{}, err
	}
	responseUser := gatewayapi.User{
		Fio:        user.Fio,
		Post:       user.Post,
		Department: user.Department,
	}
	accessToken := ctx.Value("access_token")
	refreshToken := ctx.Value("refresh_token")
	accessTokenString, ok := accessToken.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx accessToken в string")
		return &gatewayapi.GetUserInfoResponse{}, nil
	}
	refreshTokenString, ok := refreshToken.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx refreshToken в string")
		return &gatewayapi.GetUserInfoResponse{}, nil
	}
	return &gatewayapi.GetUserInfoResponse{User: &responseUser, AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}

func (us *UserService) GetUserList(ctx context.Context, request *gatewayapi.GetUserListRequest) (*gatewayapi.GetUserListResponse, error) {
	xRequestId := ctx.Value("x_request_id")
	xRequestIdString, ok := xRequestId.(string)
	if !ok {
		err := errors.New("Не получилось извлечь xRequestId из контекста")
		mylogger.Logger.Error(err)
		return &gatewayapi.GetUserListResponse{}, err
	}

	userList, err := us.userrep.List(ctx)
	if err != nil {
		err := errors.New("Не удалось получить список пользователей из бд")
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", err)
		return &gatewayapi.GetUserListResponse{}, err
	}
	var userListResponse []*gatewayapi.User
	for _, u := range userList {
		userListResponse = append(userListResponse, &gatewayapi.User{
			Fio:        u.Fio,
			Post:       u.Post,
			Department: u.Department,
		})
	}
	accessToken := ctx.Value("access_token")
	refreshToken := ctx.Value("refresh_token")
	accessTokenString, ok := accessToken.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx accessToken в string")
		return &gatewayapi.GetUserListResponse{}, nil
	}
	refreshTokenString, ok := refreshToken.(string)
	if !ok {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось преобразовать ctx refreshToken в string")
		return &gatewayapi.GetUserListResponse{}, nil
	}
	return &gatewayapi.GetUserListResponse{Result: userListResponse, AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}

func StringToUuuid(value string) (uuid.UUID, error) {
	emptyUUID := uuid.UUID{}
	uuid, err := uuid.Parse(value)
	if err != nil {
		return emptyUUID, err
	}
	return uuid, err
}

func CreateChats(ctx context.Context, client chatserviceapi.ChatServiceClient, req *chatserviceapi.CreateChatsRequest) error {
	xRequestId := ctx.Value("x_request_id")
	xRequestIdString, ok := xRequestId.(string)
	if !ok {
		err := errors.New("Не получилось извлечь xRequestId из контекста")
		mylogger.Logger.Error(err)
		xRequestIdString = ""
	}

	_, err := client.CreateChats(ctx, req)
	if err != nil {
		mylogger.Logger.Error("Запрос № ", xRequestIdString, " ", "Не удалось вызвать удаленную функцию создания чатов, либо функция возвратила ошибку")
		return err
	}
	return nil
}
