package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type RequestRepository interface {
	AllRequest() []entity.Request
	SimpanRequest(request entity.Request) entity.Request
	UpdateRequest(id_request uint32, request entity.Request) entity.Request
}

type requestRepository struct {
	connection *gorm.DB
}

func (r requestRepository) AllRequest() []entity.Request {
	var requests []entity.Request
	r.connection.Order("req_id asc").Preload("Cat").Find(&requests)
	return requests
}

func (r requestRepository) SimpanRequest(request entity.Request) entity.Request {
	r.connection.Save(&request)
	r.connection.Preload( "Cat").Find(&request)
	return request
}

func (r requestRepository) UpdateRequest(id_request uint32, request entity.Request) entity.Request {
	r.connection.Where("req_id = ?", id_request).Updates(&request)
	r.connection.Where("req_id = ?", id_request).Preload("Cat").Find(&request)
	return request
}

func NewRequestRepository(conn *gorm.DB) RequestRepository {
	return &requestRepository{
		connection: conn,
	}
}
