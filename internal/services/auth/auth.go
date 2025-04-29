package auth

import (
	"log/slog"
	"sso/internal/domain/models"
	"time"

	"golang.org/x/net/context"
)

type Auth struct {
	log          *slog.Logger
	tokenTTL     time.Duration
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvicer
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type AppProvicer interface {
	App(ctx context.Context, appId int) (models.App, error)
}

// New returns a new instance of Auth service
func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider, appProvider AppProvicer, tokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: userProvider,
		log:          log,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appId int) (string, error) {
	panic("not implemented")
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")
}

func (a *Auth) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	panic("not implemented")
}
