package persons_info

import (
	"Backend/models/persons_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetPersonInfo(conn *pgx.Conn, request persons_info.GetPersonInfoRequest) (persons_info.GetPersonInfoResponse, error) {
	query := `SELECT
				    first_name,
					second_name,
					third_name,
				    date_of_birth,
				    education,
				    academic_degree,
				    academic_rank,
				    contact_number,
				    contact_email
				FROM persons
				WHERE login = $1`

	response := persons_info.GetPersonInfoResponse{}
	err := conn.QueryRow(context.Background(), query, request.Login).Scan(
		&response.FirstName, &response.SecondName, &response.ThirdName, &response.DateOfBirth,
		&response.Education, &response.AcademicDegree, &response.AcademicRank, &response.ContactNumber, &response.ContactEmail)

	return response, err
}
