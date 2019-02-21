package nostradamus

import "time"

type Offices []Office

type Office struct {
	ID   int    `csv:"office_id"`
	Name string `csv:"office_name"`
}

type Departments []Department

type Department struct {
	OfficeID int    `csv:"office_id"`
	ID       int    `csv:"department_id"`
	Name     string `csv:"department_name"`
}

type Teams []Team

type Team struct {
	OfficeID     int    `csv:"office_id"`
	DepartmentID int    `csv:"department_id"`
	ID           int    `csv:"team_id"`
	Name         string `csv:"team_name"`
}

type Staff []Card

type Card struct {
	OfficeID         int     `csv:"office_id"`
	OfficeName       string  `csv:"office_name"`
	ID               int     `csv:"card_id"`
	AccountantNumber string  `csv:"card_accountant_number"`
	Name             string  `csv:"card_name"`
	Start            Date    `csv:"card_start"`
	End              Date    `csv:"card_end"`
	ContractHours    float64 `csv:"card_contract_hours"`
	HourlySalary     float64 `csv:"card_hourly_salary"`
	SalaryRate       float64 `csv:"card_salary_rate"`
	Department       int     `csv:"card_department"`
	Function         int     `csv:"card_function"`
	ContractNumber   int     `csv:"card_contract_number"`
	ContractProfile  int     `csv:"card_contract_profile"`
}

type Hours []Hour

type Hour struct {
	OfficeID     int       `csv:"office_id"`
	OfficeName   string    `csv:"office_name"`
	CardID       int       `csv:"card_id"`
	CardName     string    `csv:"card_name"`
	ID           int       `csv:"hour_id"`
	OfficeID     int       `csv:"hour_office_id"`
	DepartmentID int       `csv:"hour_department_id"`
	TeamID       int       `csv:"hour_team_id"`
	Date         Date      `csv:"hour_date"`
	Hours        int       `csv:"hour_hours"`
	Enter        time.Time `csv:"hour_enter"`
	Exit         time.Time `csv:"hour_exit"`
	State        int       `csv:"hour_state"`
	Type         int       `csv:"hour_type"`
}
