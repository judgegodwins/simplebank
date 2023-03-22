package token

import (
	"testing"
	"time"

	// "github.com/golang-jwt/jwt"
	"github.com/judgegodwins/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}

func TestInvalidPasetoToken(t *testing.T) {
	maker1, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	maker2, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker1.CreateToken(util.RandomOwner(), time.Minute)
	
	require.NotEmpty(t, token)
	require.NoError(t, err)

	payload, err := maker2.VerifyToken(token)

	require.Nil(t, payload)
	require.EqualError(t, err, ErrInvalidToken.Error())

}

// func TestInvalidPasetoTokenAlgNone(t *testing.T) {
// 	payload, err := NewPayload(util.RandomOwner(), time.Minute)
// 	require.NoError(t, err)

// 	PasetoToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
// 	token, err := PasetoToken.SignedString(jwt.UnsafeAllowNoneSignatureType)

// 	require.NoError(t, err)

// 	maker, err := NewPasetoMaker(util.RandomString(32))
// 	require.NoError(t, err)

// 	payload, err = maker.VerifyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrInvalidToken.Error())
// 	require.Nil(t, payload)
// }