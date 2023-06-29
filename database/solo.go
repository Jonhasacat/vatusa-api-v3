package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SoloCertification struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	Position     string
	ExpiresAt    time.Time
	RequesterCID uint64
}

func soloQuery() *gorm.DB {
	return DB.Model(&SoloCertification{}).Joins("Controller").Joins("Certification")
}

func FetchSoloCertificationById(id uint) (*SoloCertification, error) {
	var model *SoloCertification
	soloQuery().First(model, id)
	if model == nil {
		return nil, errors.New(fmt.Sprintf("Solo Certification %d not found", id))
	}
	return model, nil
}

func FetchActiveSoloCertifications() ([]*SoloCertification, error) {
	rows, err := soloQuery().Where("ExpiresAt <=", time.Now()).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var certs []*SoloCertification

	for rows.Next() {
		var c *SoloCertification
		err := DB.ScanRows(rows, c)
		if err != nil {
			return nil, err
		}
		certs = append(certs, c)
	}
	return certs, nil
}

func FetchActiveSoloCertificationsByFacility(facility string) ([]*SoloCertification, error) {
	rows, err := soloQuery().
		Where("ExpiresAt <= ?", time.Now()).
		Where("Facility = ?", facility).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var certs []*SoloCertification

	for rows.Next() {
		var c *SoloCertification
		err := DB.ScanRows(rows, c)
		if err != nil {
			return nil, err
		}
		certs = append(certs, c)
	}
	return certs, nil
}
