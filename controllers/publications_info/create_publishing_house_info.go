package publications_info

import (
	"Backend/models/publications_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreatePublishingHouse(conn *pgx.Conn, request publications_info.CreatePublishingHouseRequest) error {
	query := `INSERT INTO publishing_houses (name) VALUES ($1)`

	_, err := conn.Exec(context.Background(), query, request.Name)

	return err
}
