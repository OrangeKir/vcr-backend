package staff_info

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func GetSummaryCombinationRates(conn *pgx.Conn, personLogin string) (float32, error) {
	query := `SELECT
				    SUM(working_rate)
				FROM staff
				WHERE person_login = $1 AND status = 'Active' AND (position_type = 'InternalCombination' OR position_type = 'ExternalCombination')`

	var response *float32
	err := conn.QueryRow(context.Background(), query, personLogin).Scan(&response)

	if response == nil {
		return 0, err
	}

	return *response, err
}
