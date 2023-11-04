package service

import (
	"context"

	"github.com/Yifangmo/micro-shop-services/user/global"
	"github.com/Yifangmo/micro-shop-services/user/models"
	"github.com/Yifangmo/micro-shop-services/user/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*UserServer) GetUserMessageList(ctx context.Context, req *proto.UserMessageRequest) (*proto.UserMessageListResponse, error) {
	var resp proto.UserMessageListResponse
	var messages []models.LeavingMessage
	var messagesProto []*proto.UserMessageResponse

	dbres := global.DB.Where(&models.LeavingMessage{User: req.UserId}).Find(&messages)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	resp.Total = dbres.RowsAffected

	for _, message := range messages {
		messagesProto = append(messagesProto, &proto.UserMessageResponse{
			Id:          message.ID,
			UserId:      message.User,
			MessageType: int32(message.MessageType),
			Subject:     message.Subject,
			Message:     message.Message,
			File:        message.File,
		})
	}

	resp.Data = messagesProto
	return &resp, nil
}

func (*UserServer) CreateUserMessage(ctx context.Context, req *proto.UserMessageRequest) (*proto.IDResponse, error) {
	msg := models.LeavingMessage{
		User:        req.UserId,
		MessageType: models.MessageType(req.MessageType),
		Message:     req.Message,
		Subject:     req.Subject,
		File:        req.File,
	}
	dbres := global.DB.Create(&msg)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &proto.IDResponse{Id: msg.ID}, nil
}
