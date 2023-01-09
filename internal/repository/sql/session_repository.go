package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/domain"
	"time"
)

type sessionSQLRepository struct {
	db *sql.DB
}

func (repository sessionSQLRepository) Find(
	ctx context.Context,
	key domain.FindWith,
	val any,
) (data *domain.Session, err error) {
	q := "SELECT id, raw, created_at FROM sessions "
	if key == domain.FindWithID {
		q += " WHERE id = ?"
	}
	q += " AND WHERE deletet_at IS NULL"
	row := repository.db.QueryRowContext(ctx, q, val)

	data = &domain.Session{}
	if err := row.Scan(
		&data.ID, &data.Raw,
		&data.CreatedAt,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repository sessionSQLRepository) Create(
	ctx context.Context,
	param *domain.Session,
) error {
	q := "INSERT INTO sessions (raw, created_at) VALUES (?, ?) RETURNING id, raw, created_at"
	now := time.Now().UnixMilli()
	row := repository.db.QueryRowContext(ctx, q, param.Raw, now)

	data := &domain.Session{}
	if err := row.Scan(
		&data.ID, &data.Raw,
		&data.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (repository sessionSQLRepository) Delete(
	ctx context.Context,
	param *domain.Session,
) error {
	q := "UPDATE sessions SET deleted_at = ? WHERE id = ? RETURNING id, raw, created_at"
	now := time.Now().UnixMilli()
	row := repository.db.QueryRowContext(ctx, q, now, param.ID)

	data := &domain.Session{}
	if err := row.Scan(
		&data.ID, &data.Raw,
		&data.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func NewSessionSQLRepository() domain.ISessionRepository {
	return &sessionSQLRepository{db: configs.DbPool}
}
