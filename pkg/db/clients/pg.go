package dbclients

import (
	"context"
	"fmt"
	"github.com/deevins/educational-platform-backend/internal/config"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DBops interface for database
type DBops interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	GetPool(ctx context.Context) *pgxpool.Pool
}

// PGX содержит все операции с базой данных, включая транзакции
type PGX interface {
	DBops
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error)
}

// Tx - транзакция
type Tx interface {
	DBops
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type PostgresqlDatabase struct {
	cluster *pgxpool.Pool
}

func NewPgClient(cluster *pgxpool.Pool) *PostgresqlDatabase {
	return &PostgresqlDatabase{cluster: cluster}
}

func (db PostgresqlDatabase) GetPool(_ context.Context) *pgxpool.Pool {
	return db.cluster
}

func (db PostgresqlDatabase) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

func (db PostgresqlDatabase) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

func (db PostgresqlDatabase) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, sql, args)
}

func (db PostgresqlDatabase) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}

// NewPGClient func returns instance of PostgresSQL database client
func NewPGClient(ctx context.Context, cfg config.Config) (*PostgresqlDatabase, error) {
	dsn := generateDsn(cfg)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return NewPgClient(pool), nil
}

func generateDsn(config config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Dbname, config.SSLMode)
}
