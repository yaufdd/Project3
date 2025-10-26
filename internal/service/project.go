package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

func (ps *Servise) AddProject(ctx context.Context, project *models.Project) error {
	return ps.repo.AddProject(ctx, project)
}

func (ps *Servise) ReadProject(ctx context.Context, id int) (*models.Project, error) {
	return ps.repo.ReadProject(ctx, id)
}

func (ps *Servise) UpdateProjectTitle(ctx context.Context, newProjectInfo *models.Project) error {
	return ps.repo.UpdateProjectTitle(ctx, newProjectInfo)
}

func (ps *Servise) DeleteProject(ctx context.Context, id int) (int64, error) {
	return ps.repo.DeleteProject(ctx, id)
}

func (ps *Servise) GetProjectID(ctx context.Context, reqeust request.GetProjectID) (int, error) {
	return ps.repo.GetProjectID(ctx, reqeust)
}
