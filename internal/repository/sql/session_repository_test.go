package sql_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/domain"
	sqlRepo "github.com/aasumitro/gowa/internal/repository/sql"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type sessionSQLRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.ISessionRepository
}

func TestSessionSQLRepository(t *testing.T) {
	suite.Run(t, new(sessionSQLRepositoryTestSuite))
}

func (suite *sessionSQLRepositoryTestSuite) SetupSuite() {
	var err error

	configs.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))

	suite.NoError(err)

	suite.repo = sqlRepo.NewSessionSQLRepository()
}

func (suite *sessionSQLRepositoryTestSuite) TestRepository_Find_ExpectedSuccess() {
	data := suite.mock.
		NewRows([]string{"id", "raw", "created_at"}).
		AddRow(1, "loremipsumdolorasmet", 1672654808919)
	q := "SELECT id, raw, created_at FROM sessions WHERE id = ? AND WHERE deletet_at IS NULL"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithID, 1)
	suite.Nil(err)
	suite.NoError(err)
	suite.NotNil(res)
	suite.Equal("loremipsumdolorasmet", res.Raw)
}

func (suite *sessionSQLRepositoryTestSuite) TestRepository_Find_ExpectedError() {
	data := suite.mock.
		NewRows([]string{"id", "raw", "created_at"}).
		AddRow(nil, nil, nil)
	q := "SELECT id, raw, created_at FROM sessions WHERE id = ? AND WHERE deletet_at IS NULL"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithID, 1)
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "sql: Scan error on column")
}

func (suite *sessionSQLRepositoryTestSuite) TestRepository_Create_ExpectedSuccess() {
	data := &domain.Session{Raw: "loremipsumdolorasmet"}
	row := suite.mock.
		NewRows([]string{"id", "raw", "created_at"}).
		AddRow(1, "loremipsumdolorasmet", 1672654808919)
	q := "INSERT INTO sessions (raw, created_at) VALUES (?, ?) RETURNING id, raw, created_at"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(row)
	err := suite.repo.Create(context.TODO(), data)
	suite.Nil(err)
	suite.NoError(err)
}

func (suite *sessionSQLRepositoryTestSuite) TestRepository_Create_ExpectedError() {
	data := &domain.Session{Raw: "loremipsumdolorasmet"}
	row := suite.mock.
		NewRows([]string{"id", "raw", "created_at"}).
		AddRow(nil, nil, nil)
	q := "INSERT INTO sessions (raw, created_at) VALUES (?, ?) RETURNING id, raw, created_at"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(row)
	err := suite.repo.Create(context.TODO(), data)
	suite.NotNil(err)
	suite.Error(err)
	suite.Contains(err.Error(), "sql: Scan error on column")
}

func (suite *sessionSQLRepositoryTestSuite) TestRepository_Delete_ExpectedSuccess() {
	data := &domain.Session{ID: 1}
	row := suite.mock.
		NewRows([]string{"id", "raw", "created_at"}).
		AddRow(1, "loremipsumdolorasmet", 1672654808919)
	q := "UPDATE sessions SET deleted_at = ? WHERE id = ? RETURNING id, raw, created_at"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(row)
	err := suite.repo.Delete(context.TODO(), data)
	suite.Nil(err)
	suite.NoError(err)
}

func (suite *sessionSQLRepositoryTestSuite) TestRepository_Delete_ExpectedError() {
	data := &domain.Session{ID: 1}
	row := suite.mock.
		NewRows([]string{"id", "raw", "created_at"}).
		AddRow(nil, nil, nil)
	q := "UPDATE sessions SET deleted_at = ? WHERE id = ? RETURNING id, raw, created_at"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(row)
	err := suite.repo.Delete(context.TODO(), data)
	suite.NotNil(err)
	suite.Error(err)
	suite.Contains(err.Error(), "sql: Scan error on column")
}
