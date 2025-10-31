package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
)

func (ps *Servise) AddProject(ctx context.Context, newProject *models.Project) error {

	return ps.repo.InsertProjectTable(ctx, newProject)
}

func (ps *Servise) ReadProject(ctx context.Context, id int) (*models.Project, error) {
	return ps.repo.GetProjectInfo(ctx, id)
}

func (ps *Servise) UpdateProjectTitle(ctx context.Context, projecID int, newTitle string) error {
	return ps.repo.UpdateProjectTitle(ctx, projecID, newTitle)
}

func (ps *Servise) DeleteProject(ctx context.Context, id int) (int64, error) {
	return ps.repo.DeleteProject(ctx, id)
}

func (ps *Servise) GetProjectID(ctx context.Context, title string) (int, error) {
	project, err := ps.repo.GetProjectInfoByTitle(ctx, title)
	if err != nil {
		return 0, err
	}
	return project.ID, err
}
