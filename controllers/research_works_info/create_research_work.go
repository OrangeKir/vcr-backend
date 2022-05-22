package research_works_info

import (
	"Backend/models/research_works_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreateResearchWork(conn *pgx.Conn, request research_works_info.CreateResearchWorkRequest) error {
	query := `INSERT INTO research_works (lead_customer,
                            customer,
                            full_work_title,
                            short_work_title,
                            code,
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
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	_, err := conn.Exec(context.Background(), query,
		request.LeadCustomer, request.Customer, request.FullWorkTitle, request.ShortWorkTitle, request.Code, request.ContractNumber,
		request.ContractDate, request.TopicNumber, request.WorkType, request.StartDate, request.EndDate, request.StagesNumber,
		request.WorkPriceStructureId, request.ResponsibleExecutorLogin)

	return err
}
