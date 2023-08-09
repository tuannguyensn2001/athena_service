package auth

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/entities"
	"athena_service/utils"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type usecase struct {
	repository repo
	secretKey  string
}

func NewUsecase(repo repo) usecase {
	return usecase{repository: repo, secretKey: "secret"}
}

func (u usecase) Register(ctx context.Context, input dto.RegisterInput) error {
	_, err := u.repository.FindByPhone(ctx, input.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return app.NewBadRequestError("user registered")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = u.repository.Transaction(ctx, func(ctx context.Context) error {
		user := entities.User{
			Phone:    input.Phone,
			Email:    input.Email,
			Password: string(password),
			Role:     input.Role,
		}
		err := u.repository.Create(ctx, &user)
		if err != nil {
			return err
		}
		profile := entities.Profile{
			UserId:   user.Id,
			Username: input.Username,
			School:   input.School,
			Birthday: input.Birthday,
		}
		err = u.repository.Create(ctx, &profile)
		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err
	}

	return nil

}

func (u usecase) Login(ctx context.Context, input dto.LoginInput) (dto.LoginOutput, error) {
	result := dto.LoginOutput{}
	db := u.repository.GetDB(ctx)
	var user entities.User
	if err := db.Where("phone = ?", input.Phone).First(&user).Error; err != nil {
		return result, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return result, ErrPhoneOrPasswordNotValid.WithError(err)
	}

	if input.Role != user.Role {
		return result, ErrPhoneOrPasswordNotValid
	}

	claims := AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
		},
		UserId: user.Id,
		Role:   user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return result, nil
	}

	result.AccessToken = accessToken
	return result, nil

}

func (u usecase) Verify(ctx context.Context, token string) (entities.User, error) {
	var user entities.User
	t, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.secretKey), nil
	})
	if err != nil {
		return user, ErrTokenNotValid
	}
	if !t.Valid {
		return user, ErrTokenNotValid
	}
	claims, ok := t.Claims.(*AuthClaims)
	if !ok {
		return user, ErrTokenNotValid
	}
	err = u.repository.GetDB(ctx).Where("id = ?", claims.UserId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u usecase) GetMe(ctx context.Context) (entities.User, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return user, err
	}
	var profile entities.Profile
	if err := u.repository.GetDB(ctx).Where("user_id = ?", user.Id).First(&profile).Error; err != nil {
		return user, err
	}

	user.Profile = &profile

	return user, nil
}
