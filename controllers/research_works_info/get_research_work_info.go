package research_works_info

import (
	"Backend/models/research_works_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetResearchWorkInfo(conn *pgx.Conn, request research_works_info.GetResearchWorkInfoRequest) (research_works_info.GetResearchWorkInfoResponse, error) {
	query := `UPDATE research_works
				SET
					lead_customer,
					customer,
					full_work_title,
					short_work_title,
					contract_number,
					contract_date,
					topic_number,
					work_type,
					work_price,
					start_date,
					end_date,
					stages_number,
					work_price_structure_id,
					responsible_executor_login) 
				WHERE code = $1`

	response := research_works_info.GetResearchWorkInfoResponse{}

	err := conn.QueryRow(context.Background(), query, request.Code).Scan(&response.LeadCustomer, &response.Customer,
		&response.FullWorkTitle, &response.ShortWorkTitle, &response.ContractNumber, &response.ContractDate,
		&response.TopicNumber, &response.WorkType, &response.StartDate, &response.EndDate, &response.StagesNumber,
		&response.WorkPriceStructureId, &response.ResponsibleExecutorLogin)

	return response, err
}
