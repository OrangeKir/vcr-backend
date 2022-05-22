package publications_info

import (
	"Backend/models/publications_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetPublishingHousesInfo(conn *pgx.Conn) (publications_info.GetPublishingHousesInfoResponse, error) {
	query := `SELECT
					id,
					name
				FROM publishing_houses`

	response := publications_info.GetPublishingHousesInfoResponse{}

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return publications_info.GetPublishingHousesInfoResponse{}, err
	}

	for rows.Next() {
		record := publications_info.PublishingHouseInfo{}
		err := rows.Scan(&record.Id, &record.Name)
		if err != nil {
			return publications_info.GetPublishingHousesInfoResponse{}, err
		}

		response.PublishingHousesInfo = append(response.PublishingHousesInfo, record)
	}
	return response, err
}
