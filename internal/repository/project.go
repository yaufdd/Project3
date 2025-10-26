package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

// AddProject add project to db
func (pr *Repo) AddProject(ctx context.Context, project *models.Project) error {
	return project.Insert(ctx, pr.db, boil.Infer())
}

// ReadProject get project info by id
func (pr *Repo) ReadProject(ctx context.Context, id int) (*models.Project, error) {
	return models.FindProject(ctx, pr.db, id)
}

// UpdateProject for update project title. Need to send full json of project object,
// including id. Without ID info wont be updated
func (pr *Repo) UpdateProjectTitle(ctx context.Context, newInfoProject *models.Project) error {
	project, err := models.FindProject(ctx, pr.db, newInfoProject.ID)
	if err != nil {
		return err
	}

	project.Title = newInfoProject.Title

	_, err = project.Update(ctx, pr.db, boil.Whitelist("title", "description", "goal_amount", "collected_amount"))

	return err
}

func (pr *Repo) DeleteProject(ctx context.Context, id int) (int64, error) {
	project, err := models.FindProject(ctx, pr.db, id)
	if err != nil {
		return 0, err
	}
	return project.Delete(ctx, pr.db)
}

func (pr *Repo) GetProjectID(ctx context.Context, reqeust request.GetProjectID) (int, error) {
	project, err := models.Projects(models.ProjectWhere.Title.EQ(reqeust.Title)).One(ctx, pr.db)
	if err != nil {
		return 0, err
	}
	return project.ID, err
}
