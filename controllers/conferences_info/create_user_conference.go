package conferences_info

import (
	"Backend/models/conferences_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreateUserConferenceTopic(conn *pgx.Conn, request conferences_info.CreateUserConferenceTopicRequest) error {
	query := "INSERT INTO conferences_to_persons (person_login, conference_id, topic_name) VALUES ($1, $2, $3)"

	_, err := conn.Exec(context.Background(), query, request.Login, request.ConferenceId, request.TopicName)

	return err
}
