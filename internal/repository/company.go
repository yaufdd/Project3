package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

func (cr *Repo) AddCompanyToFounder(ctx context.Context, company *models.Company) error {
	return company.Insert(ctx, cr.db, boil.Infer())
}

func (cr *Repo) ReadCompany(ctx context.Context, id int) (*models.Company, error) {
	return models.FindCompany(ctx, cr.db, id)
}

func (cr *Repo) UpdateCompanyname(ctx context.Context, newCompanyInfo *models.Company) error {
	company, err := models.FindCompany(ctx, cr.db, newCompanyInfo.ID)
	if err != nil {
		return err
	}
	company.Name = newCompanyInfo.Name

	_, err = company.Update(ctx, cr.db, boil.Whitelist("name"))

	return err
}

func (cr *Repo) DeleteCompany(ctx context.Context, id int) (int64, error) {
	company, err := models.FindCompany(ctx, cr.db, id)
	if err != nil {
		return 0, nil
	}
	return company.Delete(ctx, cr.db)
}

func (cr *Repo) GetCompanyID(ctx context.Context, request request.GetCompanyID) (int, error) {
	company, err := models.Companies(models.CompanyWhere.Name.EQ(request.CompanyName)).One(ctx, cr.db)
	if err != nil {
		return 0, err
	}
	return company.ID, err
}
