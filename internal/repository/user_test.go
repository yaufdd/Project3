package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/yaufdd/project3/internal/models"
)

// TestCreateUserWithoutReturnsError positive test without error
func TestCreateUserWithoutReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sql := `INSERT INTO "users"`
	mock.ExpectQuery(regexp.QuoteMeta(sql)).
		WithArgs("Eric", "test@ex.com", "founder", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(100))

	repo := NewRepository(db)
	u := &models.User{Username: "Eric", Email: "test@ex.com", Role: "founder"}
	err := repo.CreateUser(context.Background(), u)

	require.NoError(t, err, "failed to create user")
	require.Equal(t, 100, u.ID)
	require.NoError(t, mock.ExpectationsWereMet())
}

// TestCreateUserTryToInsertNull Expect not null insert error.
func TestCreateUserTryToInsertNull(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	notNullError := &pq.Error{
		Code:    "23502",
		Message: `null value in column "username" violates not-null constraint`,
		Column:  "username",
		Table:   "users",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs("", "test@ex.com", "founder", sqlmock.AnyArg()).
		WillReturnError(notNullError)

	repo := NewRepository(db)
	u := &models.User{Username: "", Email: "test@ex.com", Role: "founder"}

	err := repo.CreateUser(context.Background(), u)
	require.Error(t, err)
	var pqe *pq.Error
	require.True(t, errors.As(err, &pqe))
	require.Equal(t, pq.ErrorCode("23502"), pqe.Code)
	require.NoError(t, mock.ExpectationsWereMet())
}

// TestCreateUserInsertNotUniqueUsername Expect not unique error.
func TestCreateUserInsertNotUniqueUsername(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	uniqueViolationError := &pq.Error{
		Code:    "23505",
		Message: "unique_violation",
		Column:  "username",
		Table:   "users",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs("test", "test@ex.com", "founder", sqlmock.AnyArg())

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs("test", "test@ex.com", "founder", sqlmock.AnyArg()).
		WillReturnError(uniqueViolationError)

	repo := NewRepository(db)
	u := &models.User{Username: "test", Email: "test@ex.com", Role: "founder"}

	repo.CreateUser(context.Background(), u)
	err := repo.CreateUser(context.Background(), u)

	require.Error(t, err)
	var pqe *pq.Error
	require.True(t, errors.As(err, &pqe))
	require.Equal(t, pq.ErrorCode("23505"), pqe.Code)
	require.NoError(t, mock.ExpectationsWereMet())

}

// TestCreateUserInsertRoleNotFromEnum Expect insert role not from enum error
func TestCreateUserInsertRoleNotFromEnum(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	enumErr := &pq.Error{
		Code:    "22P02",
		Message: `invalid input value for enum user_role: "hacker"`,
		Column:  "role",
		Table:   "users",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs("test", "test@ex.com", "hacker", sqlmock.AnyArg()).
		WillReturnError(enumErr)

	repo := NewRepository(db)
	u := &models.User{Username: "test", Email: "test@ex.com", Role: "hacker"}

	err := repo.CreateUser(context.Background(), u)
	require.Error(t, err)

	var pqe *pq.Error
	require.True(t, errors.As(err, &pqe))
	require.Equal(t, pq.ErrorCode("22P02"), pqe.Code)
	require.NoError(t, mock.ExpectationsWereMet())
}
