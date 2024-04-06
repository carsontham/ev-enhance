package service

import (
	"context"
	"ev-enhance/app/domain/model"
)

type ExternalPartnerAPI interface {
	GetEVInformation(ctx context.Context, s string) (model.EVInformation, error)
	//CreatePayment(obj model.ExternalPartnerObject)
}
