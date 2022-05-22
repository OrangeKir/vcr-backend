package conferences_info

import (
	"Backend/models/conferences_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetUserConferencesInfo(conn *pgx.Conn, request conferences_info.GetUserConferencesInfoRequest) (conferences_info.GetUserConferencesInfoResponse, error) {
	query := `	SELECT
					c.id,
    				c.name,
    				ctp.topic_name
				FROM conferences_to_persons AS ctp
				LEFT JOIN conferences c on c.id = ctp.conference_id
				WHERE person_login = $1
				GROUP BY c.name, ctp.topic_name`

	var conferences []conferences_info.UserConference
	rows, err := conn.Query(context.Background(), query, request.Login)

	for rows.Next() {
		var row conferences_info.UserConference

		rows.Scan(&row.Id, &row.Name, &row.TopicNames)

		conferences = append(conferences, row)
	}

	response := conferences_info.GetUserConferencesInfoResponse{Conferences: conferences}

	return response, err
}
