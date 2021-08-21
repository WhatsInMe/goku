package utilities

import (
	"github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
	"os"
	"time"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(accountId uint) (*TokenDetails, error) {
	tokenDetails := &TokenDetails{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.AccessUuid = uuid.NewV4().String()
	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUuid = uuid.NewV4().String()

	var err error

	// access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["account_id"] = accountId
	atClaims["access_uuid"] = tokenDetails.AccessUuid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDetails.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	
	// refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["account_id"] = accountId
	rtClaims["refresh_uuid"] = tokenDetails.RefreshUuid
	rtClaims["exp"] = tokenDetails.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokenDetails.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}
