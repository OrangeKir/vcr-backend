package persons_info

import (
	"Backend/models/persons_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func UpdatePersonInfo(conn *pgx.Conn, request persons_info.UpdatePersonInfoRequest) error {
	query := `	UPDATE persons
				SET
				    education = $1,
				    academic_degree = $2,
				    academic_rank = $3,
				    contact_number = $4,
				    contact_email = $5
				WHERE login = $6`
	_, err := conn.Exec(context.Background(), query, request.Education, request.AcademicDegree, request.AcademicRank,
		request.ContactNumber, request.ContactEmail, request.Login)

	return err
}
