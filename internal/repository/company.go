package repository

import (
	"context"
	"errors"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
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

func (cr *Repo) UpdateCompanyName(ctx context.Context, company *models.Company) error {
	affRow, err := company.Update(ctx, cr.db, boil.Whitelist("name"))
	if err != nil {
		return err
	}
	if affRow == 0 {
		return errors.New("no row was updated")
	}
	return err
}

func (cr *Repo) DeleteCompany(ctx context.Context, company *models.Company) (int64, error) {
	n, err := models.Companies(qm.Where("id=? AND user_id=?", company.ID, company.UserID)).DeleteAll(ctx, cr.db)
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return 0, errors.New("user not found")
	}
	return company.Delete(ctx, cr.db)
}
