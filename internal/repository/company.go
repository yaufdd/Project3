package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (cr *Repo) AddCompanyToFounder(ctx context.Context, newCompany *models.Company) error {
	return newCompany.Insert(ctx, cr.DB, boil.Infer())
}

func (cr *Repo) ReadCompany(ctx context.Context, id int) (*models.Company, error) {
	return models.FindCompany(ctx, cr.DB, id)
}

func (cr *Repo) UpdateCompanyname(ctx context.Context, newCompanyName string, companyID int) error {
	company, err := models.FindCompany(ctx, cr.DB, companyID)
	if err != nil {
		return err
	}
	company.Name = newCompanyName

	_, err = company.Update(ctx, cr.DB, boil.Whitelist("name"))

	return err
}

func (cr *Repo) DeleteCompany(ctx context.Context, id int) (int64, error) {
	company, err := models.FindCompany(ctx, cr.DB, id)
	if err != nil {
		return 0, nil
	}
	return company.Delete(ctx, cr.DB)
}

func (cr *Repo) GetCompanyID(ctx context.Context, companyName string) (int, error) {
	company, err := models.Companies(models.CompanyWhere.Name.EQ(companyName)).One(ctx, cr.DB)
	if err != nil {
		return 0, err
	}
	return company.ID, err
}
