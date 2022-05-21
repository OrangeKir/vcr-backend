package staff_info

import (
	"Backend/models/staff_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetStaffInfo(conn *pgx.Conn, request staff_info.GetStaffInfoRequest) (staff_info.GetStaffInfoResponse, error) {
	query := `SELECT
				    s.position,
				    s.working_rate,
				    s.professional_qualification_group,
				    s.qualification_level,
				    s.position_type,
				    d.organisation_name,
				    s.division_id,
				    d.division_name,
				    s.supervisor_login,
				    p.first_name as supervisor_first_name,
				    p.second_name as supervisor_second_name,
				    p.third_name as supervisor_third_name,
				    s.is_supervisor,
				    s.contract_end_date,
				    s.status,
				    s.status_change_reason
				FROM staff s
				LEFT JOIN persons p ON p.login = s.person_login
				LEFT JOIN divisions d ON s.division_id = d.division_code
				WHERE person_login = $1`

	response := staff_info.GetStaffInfoResponse{}
	rows, err := conn.Query(context.Background(), query, request.PersonLogin)

	if err != nil {
		return staff_info.GetStaffInfoResponse{}, err
	}

	for rows.Next() {
		record := staff_info.StaffRecord{}
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
