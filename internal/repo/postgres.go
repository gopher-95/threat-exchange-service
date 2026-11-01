package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/gopher-95/threat-exchange-service/internal/threat"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB реализует интерфейс Storage для работы с PostgreSQL.
type DB struct {
	db *pgxpool.Pool
}

// NewDB создает новый экземпляр DB с указанным пулом соединений.
func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{db: pool}
}

// GetInfo извлекает информацию об угрозе из базы данных по MD5.
func (d *DB) GetInfo(ctx context.Context, md5 string) (*threat.Threat, error) {
	t := threat.Threat{}
	query := "SELECT md5, name, severity, type, description, created_at, updated_at from threats WHERE md5 = $1"
	err := d.db.QueryRow(ctx, query, md5).Scan(&t.MD5, &t.Name, &t.Severity, &t.Type, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, threat.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get threat info: %w", err)
	}

	return &t, nil
}

// CreateInfo сохраняет новую угрозу в базе данных.
func (d *DB) CreateInfo(ctx context.Context, t threat.Threat) error {
	query := "INSERT INTO threats (md5, name, severity, type, description) VALUES ($1, $2, $3, $4, $5)"

	_, err := d.db.Exec(ctx, query, t.MD5, t.Name, t.Severity, t.Type, t.Description)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return threat.ErrConflict
		}
		return fmt.Errorf("failed to create threat: %w", err)
	}

	return nil
}

// UpdateInfo обновляет существующую угрозу в базе данных.
func (d *DB) UpdateInfo(ctx context.Context, t threat.Threat) error {
	query := "UPDATE threats SET name = $1, severity = $2, type = $3, description = $4, updated_at=NOW() WHERE md5 = $5"
	p, err := d.db.Exec(ctx, query, t.Name, t.Severity, t.Type, t.Description, t.MD5)
	if err != nil {
		return fmt.Errorf("failed to update threat: %w", err)
	}

	if p.RowsAffected() == 0 {
		return threat.ErrNotFound
	}

	return nil
}
