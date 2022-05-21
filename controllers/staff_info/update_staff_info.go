package staff_info

import (
	"Backend/models/staff_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func UpdateStaffInfo(conn *pgx.Conn, request staff_info.UpdateStaffInfoRequest) error {
	query := `	UPDATE staff
				SET
					position = $1,
					working_rate = $2,
					professional_qualification_group = $3,
					qualification_level = $4,
					position_type = $5,
					organisation_name = $6,
					division_id = $7,
					supervisor_login = $8,
					is_supervisor = $9,
					contract_end_date = $10,
					status = $11,
					status_change_reason = $12
				WHERE id = $13`
	_, err := conn.Exec(context.Background(), query, request.Position, request.WorkingRate, request.ProfessionalQualificationGroup,
		request.QualificationLevel, request.OrganisationName, request.DivisionId, request.SupervisorLogin, request.IsSupervisor,
		request.ContractEndDate, request.Status, request.StatusChangeReason, request.Id)

	return err
}
