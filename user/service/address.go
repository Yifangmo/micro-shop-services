package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/user/global"
	"github.com/Yifangmo/micro-shop-services/user/models"
	"github.com/Yifangmo/micro-shop-services/user/proto"
)

func (*UserServer) GetConsigneeAddressList(ctx context.Context, req *proto.ConsigneeAddressRequest) (*proto.ConsigneeAddressListResponse, error) {
	var address []models.Address
	var resp proto.ConsigneeAddressListResponse
	var addressProto []*proto.ConsigneeAddressResponse
	dbres := global.DB.Where(&models.Address{UserID: req.UserId}).Find(&address)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	resp.Total = dbres.RowsAffected
	for _, address := range address {
		addressProto = append(addressProto, &proto.ConsigneeAddressResponse{
			Id:               address.ID,
			UserId:           address.UserID,
			Province:         address.Province,
			City:             address.City,
			District:         address.District,
			ConsigneeAddress: address.ConsigneeAddress,
			ConsigneeName:    address.ConsigneeName,
			ConsigneeMobile:  address.ConsigneeMobile,
		})
	}
	resp.Data = addressProto

	return &resp, nil
}

func (*UserServer) CreateConsigneeAddress(ctx context.Context, req *proto.ConsigneeAddressRequest) (*proto.IDResponse, error) {
	address := &models.Address{
		UserID:           req.UserId,
		Province:         req.Province,
		City:             req.City,
		District:         req.District,
		ConsigneeAddress: req.ConsigneeAddress,
		ConsigneeName:    req.ConsigneeName,
		ConsigneeMobile:  req.ConsigneeMobile,
	}
	dbres := global.DB.Save(address)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &proto.IDResponse{Id: address.ID}, nil
}

func (*UserServer) DeleteConsigneeAddress(ctx context.Context, req *proto.ConsigneeAddressRequest) (*emptypb.Empty, error) {
	dbres := global.DB.Where("id=? and user=?", req.Id, req.UserId).Delete(&models.Address{})
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "the address record does exist")
	}
	return &emptypb.Empty{}, nil
}

func (*UserServer) UpdateConsigneeAddress(ctx context.Context, req *proto.ConsigneeAddressRequest) (*emptypb.Empty, error) {
	var address models.Address
	dbres := global.DB.Where("id=? and user=?", req.Id, req.UserId).First(&address)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the address record does exist")
	}

	if address.Province != "" {
		address.Province = req.Province
	}

	if address.City != "" {
		address.City = req.City
	}

	if address.District != "" {
		address.District = req.District
	}

	if address.ConsigneeAddress != "" {
		address.ConsigneeAddress = req.ConsigneeAddress
	}

	if address.ConsigneeName != "" {
		address.ConsigneeName = req.ConsigneeName
	}

	if address.ConsigneeMobile != "" {
		address.ConsigneeMobile = req.ConsigneeMobile
	}

	dbres = global.DB.Save(&address)

	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}

	return &emptypb.Empty{}, nil
}
