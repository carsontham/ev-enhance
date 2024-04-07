package service

import (
	"context"
)

type ExternalPartnerAPI interface {
	GetEVInformation(ctx context.Context, s int, chargerID string) (*EVResponse, error)
	//CreatePayment(obj model.ExternalPartnerObject)
}
