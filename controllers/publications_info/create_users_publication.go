package publications_info

import (
	"Backend/models/publications_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreateUsersPublication(conn *pgx.Conn, request publications_info.CreateUsersPublicationRequest) error {
	publicationQuery := "INSERT INTO publications (publishing_house_id, topic_name) VALUES ($1, $2) RETURNING id"
	authorsQuery := "INSERT INTO publications_to_persons (person_login, publication_id)	VALUES ($1, $2)"
	var id int

	err := conn.QueryRow(context.Background(), publicationQuery, request.PublishingHouseId, request.TopicName).Scan(&id)

	if err != nil {
		return err
	}

	for _, author := range request.AuthorsLogins {
		_, err := conn.Exec(context.Background(), authorsQuery, author, id)

		if err != nil {
			return err
		}
	}
	return err
}
