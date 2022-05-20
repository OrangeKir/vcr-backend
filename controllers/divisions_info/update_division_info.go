package divisions_info

import (
	"Backend/models/divisions_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func UpdateDivisionInfo(conn *pgx.Conn, request divisions_info.UpdateDivisionInfoRequest) error {
	query := `	UPDATE divisions
				SET
				    organisation_name = $1,
				    division_name = $2,
				    head_division_code = $3,
				    division_supervisor = $4
				WHERE division_code = $5`
	_, err := conn.Exec(context.Background(), query, request.OrganisationName, request.DivisionName, request.HeadDivisionCode, request.DivisionSupervisor, request.DivisionCode)

	return err
}
