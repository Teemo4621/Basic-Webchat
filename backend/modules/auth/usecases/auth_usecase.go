package usecases

import (
	"errors"
	"fmt"

	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"golang.org/x/crypto/bcrypt"
)

type authUse struct {
	AuthRepo entities.AuthRepository
	UserRepo entities.UserRepository
}

func NewAuthUsecase(authRepo entities.AuthRepository, userRepo entities.UserRepository) entities.AuthUsecase {
	return &authUse{AuthRepo: authRepo, UserRepo: userRepo}
}

func (a *authUse) Login(cfg *configs.Config, req *entities.AuthLoginRequest) (*entities.AuthLoginResponse, error) {
	user, err := a.UserRepo.FindOneUser(req.Username)
	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("username or password is invalid")
	}

	accessToken, err := utils.GenerateAccessToken(cfg, &entities.Jwtpassport{
		Id:       user.ID,
		Username: user.Username,
	})

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	refreshToken, err := utils.GenerateRefreshToken(cfg, &entities.Jwtpassport{
		Id:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	if err := a.AuthRepo.SaveRefreshToken(user.ID, refreshToken); err != nil {
		return nil, errors.New("something went wrong")
	}

	return &entities.AuthLoginResponse{
		Id:           user.ID,
		Username:     user.Username,
		ProfileURL:   user.ProfileURL,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authUse) Register(req *entities.AuthRegisterRequest) (*entities.AuthRegisterResponse, error) {
	checkUser, _ := a.UserRepo.FindOneUser(req.Username)

	if checkUser != nil {
		return nil, errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashed)

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	newUser, err := a.UserRepo.Create(&user)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &entities.AuthRegisterResponse{
		Id:        newUser.ID,
		Username:  newUser.Username,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (a *authUse) Me(cfg *configs.Config, id uint) (*entities.AuthMeResponse, error) {
	user, err := a.UserRepo.FindOneUserById(id)
	if err != nil {
		return nil, err
	}

	return &entities.AuthMeResponse{
		Id:         user.ID,
		Username:   user.Username,
		ProfileURL: user.ProfileURL,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (a *authUse) RefreshToken(cfg *configs.Config, req *entities.AuthRefreshTokenRequest) (*entities.AuthRefreshTokenResponse, error) {
	data, err := utils.ParseRefreshToken(cfg, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := a.UserRepo.FindOneUserById(uint(data.Id))
	if err != nil {
		return nil, err
	}

	if user.RefreshToken != req.RefreshToken {
		return nil, errors.New("refresh token is invalid")
	}

	accessToken, err := utils.GenerateAccessToken(cfg, &entities.Jwtpassport{
		Id:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(cfg, &entities.Jwtpassport{
		Id:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}

	if err := a.AuthRepo.SaveRefreshToken(user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &entities.AuthRefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
