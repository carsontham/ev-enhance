package service

import "ev-enhance/app/domain"

type ExternalPartnerAPI interface {
	GetEVInformation(id string) (domain.EVInformation, error)
	CreatePayment(obj domain.ExternalPartnerObject)
}
