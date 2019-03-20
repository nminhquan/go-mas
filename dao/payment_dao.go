package dao

import (
	. "mas/db"
	. "mas/model"
)

type PaymentDAO interface {
	CreatePayment(pmInfo PaymentInfo) string
}

type PaymentDAOImpl struct {
	db *AccountDB
}

func NewPaymentDAO(db *AccountDB) *PaymentDAOImpl {
	return &PaymentDAOImpl{db}
}

func (pmDao *PaymentDAOImpl) CreatePayment(pmInfo PaymentInfo) string {
	return pmDao.db.InsertPaymentToDB(pmInfo)
}
