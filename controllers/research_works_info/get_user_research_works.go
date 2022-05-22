package research_works_info

import (
	"Backend/models/research_works_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetUserResearchWorks(conn *pgx.Conn, request research_works_info.GetUserResearchWorksRequest) (research_works_info.GetUserResearchWorksResponse, error) {
	query := `SELECT
					code,
					short_work_title
				WHERE responsible_executor_login = '$1`

	response := research_works_info.GetUserResearchWorksResponse{}
	rows, err := conn.Query(context.Background(), query, request.Login)

	if err != nil {
		return research_works_info.GetUserResearchWorksResponse{}, err
	}

	for rows.Next() {
		record := research_works_info.ResearchWorkMainInfo{}
		err := rows.Scan(&record.Code, &record.ShortWorkTitle)
		if err != nil {
			return research_works_info.GetUserResearchWorksResponse{}, err
		}

		response.ResearchWorksMainInfo = append(response.ResearchWorksMainInfo, record)
	}

	return response, err
}
