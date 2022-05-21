package staff_info

import (
	"Backend/models/staff_info"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func CreateStaff(conn *pgx.Conn, request staff_info.CreateStaffRequest) error {

	if request.PositionType == staff_info.Main {
		positionQuery := "SELECT EXISTS (SELECT id FROM staff WHERE person_login = $1 AND position_type = 'Main')"

		var existsMain bool

		err := conn.QueryRow(context.Background(), positionQuery, request.PersonLogin).Scan(&existsMain)

		if err != nil {
			return err
		}

		if existsMain {
			return fmt.Errorf("already contain main position")
		}
	}

	query := `INSERT INTO staff(person_login,
                  				position,
                  				working_rate,
                  				professional_qualification_group,
                  				qualification_level,
                  				position_type,
                  				division_id,
                  				supervisor_login,
                  				is_supervisor,
                  				contract_end_date,
                  				status,
                  				status_change_reason
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := conn.Exec(context.Background(), query, request.PersonLogin, request.Position, request.WorkingRate,
		request.ProfessionalQualificationGroup, request.QualificationLevel, request.PositionType, request.DivisionId,
		request.SupervisorLogin, request.IsSupervisor, request.ContractEndDate, request.Status, request.StatusChangeReason)

	return err
}
