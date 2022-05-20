package divisions_info

import (
	"Backend/models/divisions_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreateDivision(conn *pgx.Conn, request divisions_info.CreateDivisionRequest) error {
	query := "INSERT INTO divisions(division_code, organisation_name, division_name, head_division_code, division_supervisor) VALUES($1, $2, $3, $4, $5)"

	_, err := conn.Exec(context.Background(), query,
		request.DivisionCode, request.OrganisationName, request.DivisionName, request.HeadDivisionCode, request.DivisionSupervisor)

	return err
}
