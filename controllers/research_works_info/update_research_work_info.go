package research_works_info

import (
	"Backend/models/research_works_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func UpdateResearchWorkInfo(conn *pgx.Conn, request research_works_info.UpdateResearchWorkInfoRequest) error {
	query := `UPDATE research_works
				SET
					lead_customer = $1,
					customer = $2,
					full_work_title = $3,
					short_work_title = $4,
					contract_number = $5,
					contract_date = $6,
					topic_number = $7,
					work_type = $8,
					work_price = $9,
					start_date = $10,
					end_date = $11,
					stages_number = $12,
					work_price_structure_id = $13,
					responsible_executor_login = $14) 
				WHERE code = $15`

	_, err := conn.Exec(context.Background(), query,
		request.LeadCustomer, request.Customer, request.FullWorkTitle, request.ShortWorkTitle, request.ContractNumber,
		request.ContractDate, request.TopicNumber, request.WorkType, request.StartDate, request.EndDate, request.StagesNumber,
		request.WorkPriceStructureId, request.ResponsibleExecutorLogin, request.Code)

	return err
}
