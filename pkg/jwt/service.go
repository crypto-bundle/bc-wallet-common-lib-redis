package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	LongLiveTokenDuration = time.Hour * 8760 * 50 // 50 years
)

var (
	ErrWrongTokenSigningMethod = errors.New("unexpected token signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrInvalidTokenClaims      = errors.New("invalid token claims")
	ErrWrongUUID               = errors.New("wrong uuid format")
	ErrWrongDateFormat         = errors.New("wrong date format")
	ErrCanNotCreateUUID        = errors.New("can not create uuid")
)

type service struct {
	logger *zap.Logger
	secret string
}

func (manager *service) GetMerchantUUID(accessToken string) (mID uuid.UUID, err error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrWrongTokenSigningMethod
			}

			return []byte(manager.secret), nil
		},
	)

	if err != nil {
		return mID, ErrInvalidToken
	}

	claim, ok := token.Claims.(*CustomClaims)
	if !ok {
		return mID, ErrInvalidTokenClaims
	}

	return claim.MerchantUUID, nil
}

func (manager *service) GenerateJWT(merchantUUID, expTime string) (string, error) {
	var (
		mUUID    uuid.UUID
		mExpTime time.Time
		err      error
	)
	if merchantUUID != "" {
		mUUID, err = uuid.Parse(merchantUUID)
		if err != nil {
			manager.logger.Error("wrong merchant uuid format", zap.Error(err))
			return "", ErrWrongUUID
		}
	} else {
		mUUID, err = uuid.NewRandom()
		if err != nil {
			return "", ErrCanNotCreateUUID
		}
	}

	if expTime != "" {
		mExpTime, err = time.Parse("2006-01-02", expTime)
		if err != nil {
			manager.logger.Error("wrong date format. Should be '2006-01-02'", zap.Error(err))
			return "", ErrWrongDateFormat
		}
	} else {
		mExpTime = time.Now().Add(LongLiveTokenDuration)
	}

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(mExpTime),
		},
		MerchantUUID: mUUID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	manager.logger.Info("Success generate token with parameters",
		zap.String("merchant_uuid", mUUID.String()))

	return token.SignedString([]byte(manager.secret))
}

func NewService(secret string,
	logger *zap.Logger,
) (s *service) {
	s = &service{
		secret: secret,
		logger: logger,
	}
	return
}
