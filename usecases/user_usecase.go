package usecases

import (
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (t *userUseCase) GetAUserByFirebaseId(user *entities.User, firebaseId string) (err error) {
	// get a user data by firebase id
	if err := t.userRepo.GetAUserByFirebaseId(user, firebaseId); err != nil {
		return err
	}
	return nil
}

func (t *userUseCase) UpdateAUser(user *entities.User, id string) (err error) {
	// find a user is haved
	var checkUser entities.User
	if errRes := t.userRepo.GetAUser(&checkUser, id); errRes != nil {
		return errRes
	}

	// update a user
	if err := t.userRepo.UpdateAUser(user); err != nil {
		return err
	}
	return nil
}

func (t *userUseCase) GetAUserByStripeID(user *entities.User, stripeID string) error {
	if err := t.userRepo.GetAUserByStripeID(user, stripeID); err != nil {
		return err
	}
	return nil
}
