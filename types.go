package nostradamus

import "encoding/xml"

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

type NostradamusEnvelope struct {
	XMLName xml.Name `xml:"nostradamus"`
}

type CreateArticlesDto struct {
	NostradamusEnvelope

	Articles Articles `xml:"articles>article"`
}

type Articles []Article

type Article struct {
	ID         int     `xml:"id,attr"`
	GroupID    int     `xml:"group_id,attr"`
	SubgroupID int     `xml:"subgroup_id,attr"`
	OfficeID   int     `xml:"office_id,attr"`
	Name       string  `xml:"name,attr"`
	Price      float64 `xml:"price,attr"`
	Tax        float64 `xml:"tax,attr"`
	State      int     `xml:"state,attr"`
}

type ArticleSubGroups struct {
	XMLName          xml.Name `xml:"nostradamus"`
	Text             string   `xml:",chardata"`
	Articlesubgroups struct {
		Text  string `xml:",chardata"`
		Group []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Parent   string `xml:"parent,attr"`
			OfficeID string `xml:"office_id,attr"`
			Name     string `xml:"name,attr"`
		} `xml:"group"`
	} `xml:"articlesubgroups"`
}
type ArticleGroups struct {
	XMLName       xml.Name `xml:"nostradamus"`
	Text          string   `xml:",chardata"`
	Articlegroups struct {
		Text  string `xml:",chardata"`
		Group []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			OfficeID string `xml:"office_id,attr"`
			Name     string `xml:"name,attr"`
		} `xml:"group"`
	} `xml:"articlegroups"`
}
type Sales struct {
	XMLName xml.Name `xml:"nostradamus"`
	Text    string   `xml:",chardata"`
	Sales   struct {
		Text         string `xml:",chardata"`
		Erase        string `xml:"erase,attr"`
		BusinessDate string `xml:"business_date,attr"`
		Sale         []struct {
			Text       string `xml:",chardata"`
			OfficeID   string `xml:"office_id,attr"`
			Enter      string `xml:"enter,attr"`
			TicketID   string `xml:"ticket_id,attr"`
			Table      string `xml:"table,attr"`
			Guests     string `xml:"guests,attr"`
			EmployeeID string `xml:"employee_id,attr"`
			StaffID    string `xml:"staff_id,attr"`
			Product    []struct {
				Text       string `xml:",chardata"`
				Datetime   string `xml:"datetime,attr"`
				PosID      string `xml:"pos_id,attr"`
				ArticleID  string `xml:"article_id,attr"`
				Ordering   string `xml:"ordering,attr"`
				Amount     string `xml:"amount,attr"`
				Value      string `xml:"value,attr"`
				EmployeeID string `xml:"employee_id,attr"`
				StaffID    string `xml:"staff_id,attr"`
			} `xml:"product"`
		} `xml:"sale"`
	} `xml:"sales"`
}
