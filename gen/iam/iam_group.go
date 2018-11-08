// Code generated by sdkgen. DO NOT EDIT.

package iam

import (
	"context"

	"google.golang.org/grpc"
)

// IAM provides access to "iam" component of Yandex.Cloud
type IAM struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewIAM creates instance of IAM
func NewIAM(g func(ctx context.Context) (*grpc.ClientConn, error)) *IAM {
	return &IAM{g}
}

// IamToken gets IamTokenService client
func (i *IAM) IamToken() *IamTokenServiceClient {
	return &IamTokenServiceClient{getConn: i.getConn}
}

// Role gets RoleService client
func (i *IAM) Role() *RoleServiceClient {
	return &RoleServiceClient{getConn: i.getConn}
}

// ServiceAccount gets ServiceAccountService client
func (i *IAM) ServiceAccount() *ServiceAccountServiceClient {
	return &ServiceAccountServiceClient{getConn: i.getConn}
}

// UserAccount gets UserAccountService client
func (i *IAM) UserAccount() *UserAccountServiceClient {
	return &UserAccountServiceClient{getConn: i.getConn}
}

// YandexPassportUserAccount gets YandexPassportUserAccountService client
func (i *IAM) YandexPassportUserAccount() *YandexPassportUserAccountServiceClient {
	return &YandexPassportUserAccountServiceClient{getConn: i.getConn}
}
