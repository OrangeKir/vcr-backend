package persons_info

import (
	"Backend/models/persons_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetPersons(conn *pgx.Conn) (persons_info.GetPersonsResponse, error) {
	query := `SELECT
					login,
				    first_name,
					second_name,
					third_name
				FROM persons`

	response := persons_info.GetPersonsResponse{}
	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return persons_info.GetPersonsResponse{}, err
	}

	for rows.Next() {
		record := persons_info.PersonMainInfo{}
		err := rows.Scan(&record.Login, &record.FirstName, &record.SecondName, &record.ThirdName)

		if err != nil {
			return persons_info.GetPersonsResponse{}, err
		}

		response.PersonsInfo = append(response.PersonsInfo, record)
	}

	return response, err
}
