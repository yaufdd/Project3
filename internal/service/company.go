package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
)

func (cs *Servise) AddCompanyToFounder(ctx context.Context, company *models.Company) error {
	return cs.repo.AddCompanyToFounder(ctx, company)
}

func (cs *Servise) ReadCompany(ctx context.Context, id int) (*models.Company, error) {
	return cs.repo.ReadCompany(ctx, id)
}

func (cs *Servise) UpdateCompanyname(ctx context.Context, newCompanyInfo *models.Company) error {
	return cs.repo.UpdateCompanyname(ctx, newCompanyInfo)
}

func (cs *Servise) DeleteCompany(ctx context.Context, id int) (int64, error) {
	return cs.repo.DeleteCompany(ctx, id)
}

func (cs *Servise) GetCompanyID(ctx context.Context, companyName string) (int, error) {
	return cs.repo.GetCompanyID(ctx, companyName)
}
