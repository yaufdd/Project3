package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
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

func (pr *Repo) GetProjectrOwnerIDByPId(ctx context.Context, projectID int) (int, error) {
	project, err := models.Projects(models.ProjectPostWhere.ID.EQ(projectID), qm.Load(qm.Rels(models.ProjectRels.Company, models.CompanyRels.User))).One(ctx, pr.db)
	if err != nil {
		return 0, err
	}
	user := project.R.Company.R.User
	return user.ID, err
}

func (pr *Repo) UpdateProjectTitle(ctx context.Context, project *models.Project) (int64, error) {
	return project.Update(ctx, pr.db, boil.Whitelist("title"))

}

func (pr *Repo) DeleteProject(ctx context.Context, id int) (int64, error) {
	project, err := models.FindProject(ctx, pr.db, id)
	if err != nil {
		return 0, err
	}
	return project.Delete(ctx, pr.db)
}
