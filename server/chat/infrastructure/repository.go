package infrastructure

type Repositories struct {
	User *UserRepository
	Group *GroupRepository
}

func NewRepositories() *Repositories {
	return &Repositories{User:NewUserRepository(), Group:NewGroupRepository()}
}