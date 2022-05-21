package price_structures_info

import (
	"Backend/models/price_structures_info"
	"context"
	"github.com/jackc/pgx/v4"
)

func GetPriceStructureInfo(conn *pgx.Conn, request price_structures_info.GetPriceStructureInfoRequest) (price_structures_info.GetPriceStructureInfoResponse, error) {
	query := `SELECT
					topic_number,
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
				FROM price_structures
				WHERE id = $1`

	response := price_structures_info.GetPriceStructureInfoResponse{}
	err := conn.QueryRow(context.Background(), query, request.Id).Scan(
		&response.TopicNumber, &response.Laboriousness, &response.BasicSalary, &response.AdditionalSalary,
		&response.PayrollFund, &response.Materials, &response.SpecialEquipment, &response.SocialPayments,
		&response.OverheadCosts, &response.BusinessTrips, &response.OtherDirectExpenses, &response.WorkOwnCost,
		&response.Profit, &response.WorkCost, &response.ValueAddedTax, &response.WorkPrice, &response.PriceType)

	return response, err
}
