package publications_info

import (
	"Backend/models/publications_info"
	"Backend/models/staff_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetPriceStructureInfo(conn *pgx.Conn) (publications_info.GetPublishingHouseInfoResponse, error) {
	query := `SELECT
					name
				FROM publishing_houses`

	response := publications_info.GetPublishingHouseInfoResponse{}
	err := conn.QueryRow(context.Background(), query).Scan(&response.Name)
	rows, err := conn.Query(context.Background(), query, request.PersonLogin)

	if err != nil {
		return publications_info.GetPublishingHouseInfoResponse{}, err
	}

	for rows.Next() {
		var record string
		err := rows.Scan(&record.Position, &record.WorkingRate, &record.ProfessionalQualificationGroup, &record.QualificationLevel,
			&record.PositionType, &record.OrganisationName, &record.DivisionId, &record.DivisionName, &record.SupervisorLogin,
			&record.SupervisorFirstName, &record.SupervisorSecondName, &record.SupervisorThirdName, &record.IsSupervisor,
			&record.ContractEndDate, &record.Status, &record.StatusChangeReason)
		if err != nil {
			return staff_info.GetStaffInfoResponse{}, err
		}

		response.StaffRecords = append(response.StaffRecords, record)
	}
	return response, err
}
