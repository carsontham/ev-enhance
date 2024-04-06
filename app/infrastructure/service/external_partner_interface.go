package service

import (
	"context"
)

type ExternalPartnerAPI interface {
	GetEVInformation(ctx context.Context, s int) (*EVResponse, error)
	//CreatePayment(obj model.ExternalPartnerObject)
}
