package service

import (
	"evo_fintech/internal/entity"
	"fmt"
	"strings"

	"github.com/gocarina/gocsv"
)

type DataStorage interface {
	Insert(data *entity.Data) error
	Select(filter string) ([]*entity.Data, error)
}

type DataService struct {
	storage DataStorage
}

func NewDataService(storage DataStorage) *DataService {
	return &DataService{storage: storage}
}

func (d DataService) Upload(file *string) error {
	var dataSlice []*entity.Data
	if err := gocsv.UnmarshalString(*file, &dataSlice); err != nil {
		return err
	}
	for _, data := range dataSlice {
		err := d.storage.Insert(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d DataService) Download(filter *entity.Filter) ([]*entity.Data, error) {
	var filterString string
	if filter.TransactionId != 0 {
		filterString += fmt.Sprintf("AND transaction_id = %d ", filter.TransactionId)
	}
	if len(filter.TerminalId) != 0 {
		filterString += fmt.Sprintf("AND terminal_id in (%s) ", filter.TerminalId)
	}
	if filter.Status != "" {
		filterString += fmt.Sprintf("AND status = '%s' ", filter.Status)
	}
	if filter.PaymentType != "" {
		filterString += fmt.Sprintf("AND payment_type = '%s' ", filter.PaymentType)
	}
	if len(filter.DatePost) != 0 {
		filterString += fmt.Sprintf("AND date_post BETWEEN '%s'  and '%s' ", filter.DatePost[0], filter.DatePost[1])
	}
	if filter.PaymentNarrative != "" {
		filterString += `AND payment_narrative LIKE '%` + filter.PaymentNarrative + `%' `
	}
	filterString = strings.Replace(filterString, "AND", "WHERE", 1)
	data, err := d.storage.Select(filterString)
	if err != nil {
		return nil, err
	}
	return data, nil
}
