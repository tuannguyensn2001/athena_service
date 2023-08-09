package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthDTO(t *testing.T) {
	t.Run("when role is not valid should return error", func(t *testing.T) {
		input := RegisterInput{
			Phone:    "08123456789",
			Password: "password",
			Email:    "hello",
			Role:     "admin",
		}

		validate := validator.New()
		err := validate.Struct(input)

		require.NotNil(t, err)
	})

	t.Run("when role is valid", func(t *testing.T) {
		input := RegisterInput{
			Phone:    "08123456789",
			Password: "password",
			Email:    "hello",
			Role:     "teacher",
			Username: "username",
		}

		validate := validator.New()
		err := validate.Struct(input)

		require.Nil(t, err)
	})
}
