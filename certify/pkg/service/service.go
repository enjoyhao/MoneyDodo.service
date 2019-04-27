package service

import (
	"context"
	"fmt"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// CertifyService describes the service.
type CertifyService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetAllUnAuth(ctx context.Context) (status bool, errinfo string, data []model.User)
	PostAuthInfo(ctx context.Context, id string, img []byte) (status bool, errinfo string, data model.User)
	PostCertifyInfo(ctx context.Context, id string, pass bool) (status bool, errinfo string, data model.User)
}

type basicCertifyService struct {
	*db.DBService
}

func (b *basicCertifyService) GetAllUnAuth(ctx context.Context) (status bool, errinfo string, data []model.User) {
	// TODO implement the business logic of GetAllUnAuth
	user := model.User{}
	rows, err := b.Engine().Where("certificationStatus = ?", 1).Rows(user)
	if err == nil {
		for rows.Next() {
			err1 := rows.Scan(&user)
			if err1 != nil {
				return false, err1.Error(), data
			}
			fmt.Println(user.Id)
			data = append(data, user)
		}
		return true, "", data
	} else {
		return false, err.Error(), data
	}
}
func (b *basicCertifyService) PostAuthInfo(ctx context.Context, id string, img []byte) (status bool, errinfo string, data model.User) {
	// TODO implement the business logic of PostAuthInfo
	user := model.User{
		Id: id,
	}
	status, err := b.Engine().Get(user)
	if status == false || err != nil {
		return false, err.Error(), data
	}
	user.CertifiedPic = img
	user.CertificationStatus = 1
	_, err = b.Engine().Id(id).Update(user)
	if err != nil {
		return false, err.Error(), data
	}
	_, err = b.Engine().Id(id).Get(data)
	if err != nil {
		return false, err.Error(), data
	}
	return true, "", data
}
func (b *basicCertifyService) PostCertifyInfo(ctx context.Context, id string, pass bool) (status bool, errinfo string, data model.User) {
	// TODO implement the business logic of PostCertifyInfo
	user := model.User{
		Id: id,
	}
	status, err := b.Engine().Get(user)
	if status == false || err != nil {
		return false, err.Error(), data
	}
	if pass {
		user.CertificationStatus = 2
	} else {
		user.CertificationStatus = 3
	}
	_, err = b.Engine().Id(id).Update(user)
	if err != nil {
		return false, err.Error(), data
	}
	_, err = b.Engine().Id(id).Get(data)
	if err != nil {
		return false, err.Error(), data
	}
	return true, "", data
}

// NewBasicCertifyService returns a naive, stateless implementation of CertifyService.
func NewBasicCertifyService() CertifyService {
	basicCertifyService := &basicCertifyService{
		&db.DBService{},
	}
	err := basicCertifyService.Bind("conf/conf.users.yml")
	if err != nil {
		log.Printf("The UserService failed to bind with mysql")
	}
	return basicCertifyService
}

// New returns a CertifyService with all of the expected middleware wired in.
func New(middleware []Middleware) CertifyService {
	var svc CertifyService = NewBasicCertifyService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}