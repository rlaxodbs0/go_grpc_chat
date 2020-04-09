package infrastructure

import (
	"errors"
	"go_grpc_chat/server/chat/domain/model"
	"go_grpc_chat/server/chat/domain/repository"
	"sync"
)

type GroupRepository struct {
	mu sync.RWMutex
	entity uint64
	groupMap *DB
}

var _ repository.GroupRepositoryImpl = &GroupRepository{} //interface impl

func NewGroupRepository() *GroupRepository {
	info := sync.Map{}
	return &GroupRepository{groupMap: &DB{info: info}}
}

func (gr *GroupRepository) CreateGroup(group *model.Group) (uint64, error) {
	gr.mu.Lock()
	group.ID = gr.entity
	gr.groupMap.info.Store(gr.entity, &group)
	gr.entity ++
	gr.mu.Lock()
	return group.ID, nil
}

func (gr *GroupRepository) DeleteGroupByID(id uint64) error {
	gr.groupMap.info.Delete(id)
	return nil
}

func (gr *GroupRepository) GetGroupByID(id uint64) (*model.Group,error) {
	if res, ok := gr.groupMap.info.Load(id); ok {
		return res.(*model.Group), nil
	}
	return nil, errors.New("No group")
}