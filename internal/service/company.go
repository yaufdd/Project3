package service

import (
	"context"
	"errors"

	"github.com/yaufdd/project3/internal/models"
)

func (cs *Servise) AddCompanyToFounder(ctx context.Context, newCompanyName string, uesrIDFromToken int) error {
	newCompany := models.Company{
		Name: newCompanyName,
	}
	newCompany.UserID = uesrIDFromToken
	return cs.repo.InsertCompanyTable(ctx, &newCompany)
}

func (cs *Servise) ReadCompany(ctx context.Context, id int) (*models.Company, error) {
	return cs.repo.GetCompanyInfo(ctx, id)
}

func (cs *Servise) UpdateCompanyName(ctx context.Context, newCompanyName string, companyID, uesrIDFromToken int) error {
	company, err := cs.repo.GetCompanyInfo(ctx, companyID)
	if err != nil {
		return err
	}
	if company.UserID != uesrIDFromToken {
		return errors.New("forbidden")
	}
	company.Name = newCompanyName
	return cs.repo.UpdateCompanyName(ctx, company)
}

func (cs *Servise) DeleteCompany(ctx context.Context, companyID, uesrIDFromToken int) (int64, error) {
	company, err := cs.repo.GetCompanyInfo(ctx, companyID)
	if err != nil {
		return 0, err
	}
	company.UserID = uesrIDFromToken
	return cs.repo.DeleteCompany(ctx, company)
}

func (cs *Servise) GetCompanyID(ctx context.Context, companyName string) (int, error) {
	company, err := cs.repo.GetCompanyInfoByName(ctx, companyName)
	if err != nil {
		return 0, err
	}
	return company.ID, err
}
