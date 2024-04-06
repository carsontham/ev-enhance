package service

import (
	"ev-enhance/app/domain/model"
)

type ExternalPartnerAPI interface {
	GetEVInformation(id string) (model.EVInformation, error)
	CreatePayment(obj model.ExternalPartnerObject)
}
