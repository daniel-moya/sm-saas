package services

import (
    "context"
    "errors"
    "user-service/internal/models"
    "user-service/internal/repositories"
    "user-service/internal/auth"
    "time"
)

type UserService interface {
    RegisterUser(ctx context.Context, email, password string) (*models.User, error)
    LoginUser(ctx context.Context, email, password string) (string, error)
    Delete(ctx context.Context, id int64) (error)
}

type userService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, email, password string) (*models.User, error) {
    existingUser, err := s.userRepo.GetUserByEmail(ctx, email)

    if existingUser != nil {
        return nil, errors.New("Error registering user")
    }

    if existingUser != nil {
        return nil, errors.New("email already in use")
    }

    hashedPassword, err := auth.HashPassword(password)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:     email,
        Password:  hashedPassword,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    err = s.userRepo.CreateUser(ctx, user)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *userService) LoginUser(ctx context.Context, email, password string) (string, error) {
    user, err := s.userRepo.GetUserByEmail(ctx, email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    if !auth.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid email or password")
    }

    token, err := auth.GenerateJWT(user.ID, user.Email)
    if err != nil {
        return "", err
    }

    return token, nil
}


func (s *userService) Delete(ctx context.Context, id int64) (error) {
    // if err != nil {
    //     return err
    // }

    return nil
}
