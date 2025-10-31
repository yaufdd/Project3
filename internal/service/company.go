package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
)

func (cs *Servise) AddCompanyToFounder(ctx context.Context, company *models.Company) error {
	return cs.repo.InsertCompanyTable(ctx, company)
}

func (cs *Servise) ReadCompany(ctx context.Context, id int) (*models.Company, error) {
	return cs.repo.GetCompanyInfo(ctx, id)
}

func (cs *Servise) UpdateCompanyName(ctx context.Context, newCompanyName string, companyID int) error {
	return cs.repo.UpdateCompanyName(ctx, newCompanyName, companyID)
}

func (cs *Servise) DeleteCompany(ctx context.Context, id int) (int64, error) {
	return cs.repo.DeleteCompany(ctx, id)
}

func (cs *Servise) GetCompanyID(ctx context.Context, companyName string) (int, error) {
	company, err := cs.repo.GetCompanyInfoByName(ctx, companyName)
	if err != nil {
		return 0, err
	}
	return company.ID, err
}
