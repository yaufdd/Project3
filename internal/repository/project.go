package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

// AddProject add project to db
func (pr *Repo) AddProject(ctx context.Context, newProject *models.Project) error {
	return newProject.Insert(ctx, pr.DB, boil.Infer())
}

// ReadProject get project info by id
func (pr *Repo) ReadProject(ctx context.Context, id int) (*models.Project, error) {
	return models.FindProject(ctx, pr.DB, id)
}

// UpdateProject for update project title. Need to send full json of project object,
// including id. Without ID info wont be updated
func (pr *Repo) UpdateProjectTitle(ctx context.Context, projecID int, newTitle string) error {
	project, err := models.FindProject(ctx, pr.DB, projecID)
	if err != nil {
		return err
	}

	project.Title = newTitle

	_, err = project.Update(ctx, pr.DB, boil.Whitelist("title", "description", "goal_amount", "collected_amount"))

	return err
}

func (pr *Repo) DeleteProject(ctx context.Context, id int) (int64, error) {
	project, err := models.FindProject(ctx, pr.DB, id)
	if err != nil {
		return 0, err
	}
	return project.Delete(ctx, pr.DB)
}

func (pr *Repo) GetProjectID(ctx context.Context, title string) (int, error) {
	project, err := models.Projects(models.ProjectWhere.Title.EQ(title)).One(ctx, pr.DB)
	if err != nil {
		return 0, err
	}
	return project.ID, err
}
