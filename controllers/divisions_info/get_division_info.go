package divisions_info

import (
	"Backend/models/divisions_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetDivisionInfo(conn *pgx.Conn, request divisions_info.GetDivisionInfoRequest) (divisions_info.GetDivisionInfoResponse, error) {
	query := `SELECT
				    d.organisation_name,
					d.division_name,
					d.head_division_code,
				    d.division_supervisor,
					hd.division_name AS head_division_name
				FROM divisions as d
				LEFT JOIN divisions as hd
				    ON d.head_division_code = hd.division_code

				WHERE d.division_code = $1`

	response := divisions_info.GetDivisionInfoResponse{}
	err := conn.QueryRow(context.Background(), query, request.DivisionCode).Scan(&response.OrganisationName, &response.DivisionName, &response.HeadDivisionCode, &response.DivisionSupervisor, &response.HeadDivisionName)

	return response, err
}
