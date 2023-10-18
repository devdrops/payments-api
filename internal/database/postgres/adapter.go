package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"payments-api/internal/config"

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

func (adp *PostgresAdapter) Update(ctx context.Context, table string, columns []string, filters []string, values ...any) error {
	raw := "UPDATE %s SET %s "
	values := make([]string, len(values))
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

// GetOne reads zero or one result that match a criteria.
func (adp *PostgresAdapter) GetOne(ctx context.Context, table string, columns []string, filters []string) *sql.Row {
	instruction := adp.prepareSelect(table, columns, filters)

	return adp.conn.QueryRowContext(ctx, instruction)
}

// GetMany reads zero or more results that match a criteria.
func (adp *PostgresAdapter) GetMany(ctx context.Context, table string, columns []string, filters []string) (*sql.Rows, error) {
	instruction := adp.prepareSelect(table, columns, filters)

	return adp.conn.QueryContext(ctx, instruction)
}

// prepareSelect provides a SELECT instruction, with the given inputs.
func (adp *PostgresAdapter) prepareSelect(table string, columns []string, filters []string) string {
	raw := "SELECT %s FROM %s "
	instruction := fmt.Sprintf(raw, strings.Join(columns, ", "), table)
	if len(filters) > 0 {
		instruction += strings.Join(filters, " ")
	}
	instruction += ";"

	return instruction
}

func (adp *PostgresAdapter) Ping(ctx context.Context) (bool, error) {
	err := adp.conn.PingContext(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}
