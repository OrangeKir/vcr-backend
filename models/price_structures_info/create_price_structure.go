package price_structures_info

type CreatePriceStructureRequest struct {
	TopicNumber         int
	Laboriousness       int
	BasicSalary         int
	AdditionalSalary    int
	PayrollFund         int
	Materials           int
	SpecialEquipment    int
	SocialPayments      int
	OverheadCosts       int
	BusinessTrips       int
	OtherDirectExpenses int
	WorkOwnCost         int
	Profit              int
	WorkCost            int
	ValueAddedTax       int
	WorkPrice           int
	PriceType           PriceType
}
