package price_structures_info

import (
	"Backend/models/price_structures_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func CreatePriceStructure(conn *pgx.Conn, request price_structures_info.CreatePriceStructureRequest) error {
	query := `INSERT INTO price_structures (topic_number,
                              				laboriousness,
                              				basic_salary,
                              				additional_salary,
                              				payroll_fund,
                              				materials,
                              				special_equipment,
                              				social_payments,
                              				overhead_costs,
                              				business_trips,
                              				other_direct_expenses,
                              				work_own_cost,
                              				profit,
                              				work_cost,
                              				value_added_tax,
                              				work_price, 
                              				price_type
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`

	_, err := conn.Exec(context.Background(), query,
		request.TopicNumber, request.Laboriousness, request.BasicSalary, request.AdditionalSalary,
		request.PayrollFund, request.Materials, request.SpecialEquipment, request.SocialPayments,
		request.OverheadCosts, request.BusinessTrips, request.OtherDirectExpenses, request.WorkOwnCost, request.Profit,
		request.WorkCost, request.ValueAddedTax, request.WorkPrice, request.PriceType)

	return err
}
