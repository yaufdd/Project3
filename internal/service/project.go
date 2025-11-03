package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/yaufdd/project3/internal/models"
)

func (ps *Servise) AddProject(ctx context.Context, newProject *models.Project, userIDFromToken int) error {
	fmt.Println(newProject.CompanyID)
	company, err := ps.repo.GetCompanyInfo(ctx, newProject.CompanyID)
	if err != nil {
		return err
	}
	if company.UserID != userIDFromToken {
		return errors.New("forbidden")
	}
	return ps.repo.InsertProjectTable(ctx, newProject)
}

func (ps *Servise) ReadProject(ctx context.Context, id int) (*models.Project, error) {
	return ps.repo.GetProjectInfo(ctx, id)
}

func (ps *Servise) UpdateProjectTitle(ctx context.Context, projectID int, newTitle string, userIDFromToken int) (int64, error) {
	project, err := ps.repo.GetProjectInfo(ctx, projectID)
	if err != nil {
		return 0, err
	}
	company, err := ps.repo.GetCompanyInfo(ctx, project.CompanyID)
	if err != nil {
		return 0, err
	}
	if company.UserID != userIDFromToken {
		return 0, errors.New("forbidden")
	}
	project.Title = newTitle
	return ps.repo.UpdateProjectTitle(ctx, project)
}

func (ps *Servise) DeleteProject(ctx context.Context, projectID int, userIDFromToken int) (int64, error) {
	project, err := ps.repo.GetProjectInfo(ctx, projectID)
	if err != nil {
		return 0, err
	}
	company, err := ps.repo.GetCompanyInfo(ctx, project.CompanyID)
	if err != nil {
		return 0, err
	}
	if company.UserID != userIDFromToken {
		return 0, errors.New("forbidden")
	}
	return ps.repo.DeleteProject(ctx, projectID)
}

func (ps *Servise) GetProjectID(ctx context.Context, title string) (int, error) {
	project, err := ps.repo.GetProjectInfoByTitle(ctx, title)
	if err != nil {
		return 0, err
	}
	return project.ID, err
}
