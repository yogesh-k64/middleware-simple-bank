package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedpass, err := HashedPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedpass)

	err = CheckPassword(hashedpass, password)
	require.NoError(t, err)

	wrongPass := RandomString(6)

	err = CheckPassword(hashedpass, wrongPass)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
