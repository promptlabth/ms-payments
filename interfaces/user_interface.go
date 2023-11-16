package interfaces

import "github.com/promptlabth/ms-payments/entities"

type UserUseCase interface {
	GetAUserByFirebaseId(out *entities.User, firebaseId string) (err error)
	UpdateAUser(out *entities.User, id string) (err error)
}

type UserRepository interface {
	GetAUser(out *entities.User, id string) (err error)
	GetAUserByFirebaseId(out *entities.User, firebaseId string) (err error)
	UpdateAUser(out *entities.User) (err error)
}
