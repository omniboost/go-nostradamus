package nostradamus

type Offices []Office

type Office struct {
	OfficeID   int    `csv:"office_id"`
	OfficeName string `csv:"office_name"`
}

type Departments []Department

type Department struct {
	OfficeID       int    `csv:"office_id"`
	DepartmentID   int    `csv:"department_id"`
	DepartmentName string `csv:"department_name"`
}

type Teams []Team

type Team struct {
	OfficeID     int    `csv:"office_id"`
	DepartmentID int    `csv:"department_id"`
	TeamID       int    `csv:"team_id"`
	TeamName     string `csv:"team_name"`
}

type Staff []Card

type Card struct {
	OfficeID             int       `csv:"office_id"`
	OfficeName           string    `csv:"office_name"`
	CardID               int       `csv:"card_id"`
	CardAccountantNumber string    `csv:"card_accountant_number"`
	CardName             string    `csv:"card_name"`
	CardStart            Date      `csv:"card_start"`
	CardEnd              *Date     `csv:"card_end"`
	CardContractHours    float64   `csv:"card_contract_hours"`
	CardHourlySalary     float64   `csv:"card_hourly_salary"`
	CardSalaryRate       float64   `csv:"card_salary_rate"`
	CardDepartment       int       `csv:"card_department"`
	CardFunction         int       `csv:"card_function"`
	CardContractNumber   CrappyInt `csv:"card_contract_number"`
	CardContractProfile  int       `csv:"card_contract_profile"`
}

type Hours []Hour

type Hour struct {
	OfficeID         int      `csv:"office_id"`
	OfficeName       string   `csv:"office_name"`
	CardID           int      `csv:"card_id"`
	CardName         string   `csv:"card_name"`
	HourID           int      `csv:"hour_id"`
	HourOfficeID     int      `csv:"hour_office_id"`
	HourDepartmentID int      `csv:"hour_department_id"`
	HourTeamID       int      `csv:"hour_team_id"`
	HourDate         Date     `csv:"hour_date"`
	HourHours        float64  `csv:"hour_hours"`
	HourEnter        DateTime `csv:"hour_enter"`
	HourExit         DateTime `csv:"hour_exit"`
	HourState        int      `csv:"hour_state"`
	HourType         int      `csv:"hour_type"`
}
