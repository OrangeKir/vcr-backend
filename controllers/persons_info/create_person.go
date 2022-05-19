package persons_info

import (
	"Backend/models/persons_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreatePerson(conn *pgx.Conn, request persons_info.CreatePersonRequest) error {
	query := "INSERT INTO persons(login, first_name, second_name, third_name, date_of_birth, education, academic_degree, academic_rank, contact_number, contact_email) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	_, err := conn.Exec(context.Background(), query,
		request.Login, request.FirstName, request.SecondName, request.ThirdName,
		request.DateOfBirth, request.Education, request.AcademicDegree, request.AcademicRank,
		request.ContactNumber, request.ContactEmail)

	return err
}
