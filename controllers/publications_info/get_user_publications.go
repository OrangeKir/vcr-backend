package publications_info

import (
	"Backend/models/publications_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetUserPublications(conn *pgx.Conn, request publications_info.GetUserPublicationsInfoByLoginRequest) (publications_info.GetUserPublicationsInfoByLoginResponse, error) {
	query := `SELECT
				    p.id,
				    p.topic_name,
				    ph.name
				FROM publications p
				LEFT JOIN publications_to_persons ptp on p.id = ptp.publication_id
				LEFT JOIN publishing_houses ph on ph.id = p.publishing_house_id
				WHERE person_login = '$1`

	response := publications_info.GetUserPublicationsInfoByLoginResponse{}
	rows, err := conn.Query(context.Background(), query, request.Login)

	if err != nil {
		return publications_info.GetUserPublicationsInfoByLoginResponse{}, err
	}

	for rows.Next() {
		record := publications_info.Publication{}
		err := rows.Scan(&record.PublicationId, &record.TopicName, &record.PublishingHouseName)
		if err != nil {
			return publications_info.GetUserPublicationsInfoByLoginResponse{}, err
		}

		response.Publications = append(response.Publications, record)
	}

	return response, err
}
