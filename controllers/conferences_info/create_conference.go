package conferences_info

import (
	"Backend/models/conferences_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreateConference(conn *pgx.Conn, request conferences_info.CreateConferenceRequest) error {
	query := "INSERT INTO conferences(name) VALUES($1)"

	_, err := conn.Exec(context.Background(), query, request.Name)

	return err
}
