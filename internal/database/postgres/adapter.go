package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"payments-api/internal/config"
	"payments-api/internal/database"

	_ "github.com/lib/pq"
)

var (
	name    string = "postgres"
	connStr string = "%s://%s:%s@%s/%s?sslmode=%s"
)

type PostgresAdapter struct {
	conn *sql.DB
}

func NewAdapter(c *config.Config) (*PostgresAdapter, error) {
	cs := fmt.Sprintf(connStr, name, c.DbUser, c.DbPass, c.DbHost, c.DbName, c.DbSSLM)
	db, err := sql.Open(name, cs)
	if err != nil {
		return nil, err
	}

	return &PostgresAdapter{
		conn: db,
	}, nil
}

func (adp *PostgresAdapter) Insert(ctx context.Context, table string, columns []string, values ...any) error {
	raw := "INSERT INTO %s(%s) VALUES(%s);"
	v := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		v[i] = fmt.Sprintf("$%d", i+1)
	}
	instruction := fmt.Sprintf(raw, table, strings.Join(columns, ", "), strings.Join(v, ", "))

	stmt, err := adp.conn.PrepareContext(ctx, instruction)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, values...); err != nil {
		return err
	}

	return nil
}

func (adp *PostgresAdapter) GetOne(ctx context.Context, table string, columns []string, filters []string, ent database.Entity) (database.Entity, error) {
	raw := "SELECT %s FROM %s "
	instruction := fmt.Sprintf(raw, strings.Join(columns, ", "), table)
	if len(filters) > 0 {
		instruction += strings.Join(filters, " ")
	}
	instruction += ";"

	stmt, err := adp.conn.PrepareContext(ctx, instruction)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n\nAAAAAAAAAAAAAAAAAAAAAAA %#v\n\n", res)

	return nil, nil
}

func (adp *PostgresAdapter) Ping(ctx context.Context) (bool, error) {
	err := adp.conn.PingContext(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}
