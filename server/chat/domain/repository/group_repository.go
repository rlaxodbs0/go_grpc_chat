package repository

import "go_grpc_chat/server/chat/domain/model"

type GroupRepositoryImpl interface  {
	GetGroupByID(id uint64) *model.Group
}
