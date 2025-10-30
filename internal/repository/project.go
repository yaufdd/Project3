package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

// AddProject add project to db
func (pr *Repo) InsertProjectTable(ctx context.Context, newProject *models.Project) error {
	return newProject.Insert(ctx, pr.db, boil.Infer())
}

// ReadProject get project info by id
func (pr *Repo) GetProjectInfo(ctx context.Context, id int) (*models.Project, error) {
	return models.FindProject(ctx, pr.db, id)
}

func (pr *Repo) GetProjectInfoByTitle(ctx context.Context, title string) (*models.Project, error) {
	return models.Projects(models.ProjectWhere.Title.EQ(title)).One(ctx, pr.db)
}

// UpdateProject for update project title. Need to send full json of project object,
// including id. Without ID info wont be updated
func (pr *Repo) UpdateProjectTitle(ctx context.Context, projecID int, newTitle string) error {
	project, err := models.FindProject(ctx, pr.db, projecID)
	if err != nil {
		return err
	}

	project.Title = newTitle

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
