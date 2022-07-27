package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type RequestService interface {
	AllRequest() []entity.Request
	SimpanRequest(request dto.CreateRequest) entity.Request
	UpdateRequest(id_request uint32, request dto.UpdateRequest) entity.Request
}

type requestService struct {
	requestRepository repository.RequestRepository
}

func (r requestService) AllRequest() []entity.Request {
	return r.requestRepository.AllRequest()
}

func (r requestService) SimpanRequest(request dto.CreateRequest) entity.Request {
	newRequest := entity.Request{}
	errSmap := smapping.FillStruct(&newRequest, smapping.MapFields(&request))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := r.requestRepository.SimpanRequest(newRequest)
	return res
}

func (r requestService) UpdateRequest(id_request uint32, request dto.UpdateRequest) entity.Request {
	upRequest := entity.Request{}
	errSmap := smapping.FillStruct(&upRequest, smapping.MapFields(&request))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := r.requestRepository.UpdateRequest(id_request, upRequest)
		return res
	}
}

func NewRequestService(requestRep repository.RequestRepository) RequestService {
	return &requestService{
		requestRepository: requestRep,
	}
}