package server

import (
	"context"
	"go.ajitem.com/pb/user"
)

type UserHandler struct {
	Allowed []string
}

func (uh *UserHandler) Register(context.Context, *user.RegisterRequest) (*user.RegisterResponse, error) {

}

func (uh *UserHandler) Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error) {

}

func (uh *UserHandler) ForgotPassword(context.Context, *user.ForgotPasswordRequest) (*user.ForgotPasswordResponse, error) {

}

func (uh *UserHandler) ResetPassword(context.Context, *user.ResetPasswordRequest) (*user.ResetPasswordResponse, error) {

}

func (uh *UserHandler) ChangePassword(context.Context, *user.ChangePasswordRequest) (*user.ChangePasswordResponse, error) {

}
