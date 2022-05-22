package conferences_info

import (
	"Backend/models/conferences_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetConferencesInfo(conn *pgx.Conn) (conferences_info.GetConferencesResponse, error) {
	query := `	SELECT
					c.id,
    				c.name
				FROM conferences`

	response := conferences_info.GetConferencesResponse{}
	rows, err := conn.Query(context.Background(), query)

	for rows.Next() {
		row := conferences_info.Conference{}

		err := rows.Scan(&row.Id, &row.Name)

		if err != nil {
			return conferences_info.GetConferencesResponse{}, err
		}
		response.Conferences = append(response.Conferences, row)
	}

	return response, err
}
