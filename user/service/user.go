package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Yifangmo/micro-shop-services/common"
	"github.com/Yifangmo/micro-shop-services/user/global"
	"github.com/Yifangmo/micro-shop-services/user/models"
	"github.com/Yifangmo/micro-shop-services/user/proto"
	"github.com/Yifangmo/micro-shop-services/user/utils"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

func (s *UserServer) GetUserList(ctx context.Context, req *common.PageInfo) (*proto.UserListResponse, error) {
	resp := &proto.UserListResponse{}
	dbres := global.DB.Model(&models.User{}).Count(&resp.Total)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	var users []models.User
	dbres = global.DB.Scopes(utils.Paginate(int(req.PageNumber), int(req.PageSize))).Find(&users)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	for _, user := range users {
		resp.Data = append(resp.Data, BuildUserInfoResp(user))
	}
	return resp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user models.User
	dbres := global.DB.Where(&models.User{Mobile: req.Mobile}).First(&user)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the user does not exist, mobile: %s", req.Mobile)
	}

	return BuildUserInfoResp(user), nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *proto.UserIDRequest) (*proto.UserInfoResponse, error) {
	var user models.User
	dbres := global.DB.First(&user, req.Id)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the user does not exist, user id: %d", req.Id)

	}
	return BuildUserInfoResp(user), nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.IDResponse, error) {
	var user models.User
	dbres := global.DB.Where(&models.User{Mobile: req.Mobile}).Take(&user)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "the user mobile does exist: %s", req.Mobile)
	}
	user.Mobile = req.Mobile
	user.Nickname = req.NickName
	user.Password = utils.GenStorePassword(req.Password)

	dbres = global.DB.Create(&user)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &proto.IDResponse{Id: user.ID}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*empty.Empty, error) {
	var user models.User
	dbres := global.DB.First(&user, req.Id)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the user does not exist, user id: %d", req.Id)
	}
	user.Nickname = req.NickName
	user.Gender = models.Gender(req.Gender)
	if req.Birthday != nil {
		t := req.Birthday.AsTime()
		user.Birthday = &t
	}

	dbres = global.DB.Save(&user)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.CheckPasswordRequest) (*proto.CheckPasswordResponse, error) {
	return &proto.CheckPasswordResponse{Success: utils.VerifyPassword(req.RawPassword, req.EncryptedPassword)}, nil
}

func BuildUserInfoResp(user models.User) *proto.UserInfoResponse {
	resp := &proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		NickName: user.Nickname,
		Gender:   proto.Gender(user.Gender),
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
	}
	if user.Birthday != nil {
		resp.Birthday = timestamppb.New(*user.Birthday)
	}
	return resp
}
