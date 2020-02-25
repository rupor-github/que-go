package que

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4"
	// "github.com/jackc/pgx/v4/log/testingadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

var testConnConfig *pgx.ConnConfig

func openTestClientMaxConns(t testing.TB, maxConnections int) *Client {
	dsn := fmt.Sprintf("user=app_crossdevice password=app_crossdevice host=localhost dbname=crossdevice pool_max_conns=%d", maxConnections)
	cfg, err := pgxpool.ParseConfig(dsn)
	cfg.ConnConfig.RuntimeParams["search_path"] = "crossdevice"
	// cfg.ConnConfig.Logger = testingadapter.NewLogger(t)
	// cfg.ConnConfig.LogLevel = pgx.LogLevelDebug
	cfg.AfterConnect = PrepareStatements
	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	testConnConfig = cfg.ConnConfig
	return NewClient(pool)
}

func openTestClient(t testing.TB) *Client {
	return openTestClientMaxConns(t, 5)
}

func truncateAndClose(pool *pgxpool.Pool) {
	if _, err := pool.Exec(context.Background(), "TRUNCATE TABLE que_jobs"); err != nil {
		panic(err)
	}
	pool.Close()
}

func findOneJob(q queryable) (*Job, error) {
	findSQL := `
	SELECT priority, run_at, job_id, job_class, args, error_count, last_error, queue
	FROM que_jobs LIMIT 1`

	j := &Job{}
	err := q.QueryRow(context.Background(), findSQL).Scan(
		&j.Priority,
		&j.RunAt,
		&j.ID,
		&j.Type,
		&j.Args,
		&j.ErrorCount,
		&j.LastError,
		&j.Queue,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return j, nil
}
