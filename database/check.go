package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func CheckRequirements(conn *pgx.Conn) {
	err := conn.QueryRow(context.Background(), "SELECT * FROM pg_stat_statements ORDER BY mean_time DESC;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "DB check is failed: %v\n", err)
		os.Exit(1)
	}
}
