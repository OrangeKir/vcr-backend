package price_structures_info

import (
	"Backend/models/price_structures_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func UpdatePriceStructureInfo(conn *pgx.Conn, request price_structures_info.UpdatePriceStructureInfoRequest) error {
	query := `	UPDATE price_structures
				SET
				    topic_number = $1,
                    laboriousness = $2,
                    basic_salary = $3,
                    additional_salary = $4,
                    payroll_fund = $5,
                    materials = $6,
                    special_equipment = $7,
                    social_payments = $8,
                    overhead_costs = $9,
                    business_trips = $10,
                    other_direct_expenses = $11,
                    work_own_cost = $12,
                    profit = $13,
                    work_cost = $14,
                    value_added_tax = $15,
                    work_price = $16, 
                    price_type = $17
				WHERE id = $18`
	_, err := conn.Exec(context.Background(), query, request.TopicNumber, request.Laboriousness, request.BasicSalary,
		request.AdditionalSalary, request.PayrollFund, request.Materials, request.SpecialEquipment, request.SocialPayments,
		request.OverheadCosts, request.BusinessTrips, request.OtherDirectExpenses, request.WorkOwnCost, request.Profit,
		request.WorkCost, request.ValueAddedTax, request.WorkPrice, request.PriceType, request.Id)

	return err
}
