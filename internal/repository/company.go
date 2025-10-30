package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (cr *Repo) InsertCompanyTable(ctx context.Context, newCompany *models.Company) error {
	return newCompany.Insert(ctx, cr.db, boil.Infer())
}

func (cr *Repo) GetCompanyInfo(ctx context.Context, id int) (*models.Company, error) {
	return models.FindCompany(ctx, cr.db, id)
}

func (cr *Repo) GetCompanyInfoByName(ctx context.Context, companyName string) (*models.Company, error) {
	return models.Companies(models.CompanyWhere.Name.EQ(companyName)).One(ctx, cr.db)
}

func (cr *Repo) GetCompanyOwnerIDbyCname(ctx context.Context, companyName string) (int, error) {
	company, err := models.Companies(models.CompanyWhere.Name.EQ(companyName)).One(ctx, cr.db)
	if err != nil {
		return 0, nil
	}
	return company.UserID, err
}

func (cr *Repo) UpdateCompanyName(ctx context.Context, newCompanyName string, companyID int) error {
	company, err := models.FindCompany(ctx, cr.db, companyID)
	if err != nil {
		return err
	}
	company.Name = newCompanyName

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
