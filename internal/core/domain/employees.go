package domain

import "time"

type Employees struct {
	Id        string    `json:"id" db:"id" `
	Fullname  string    `json:"fullname" db:"fullname" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required,lowercase"`
	Password  string    `json:"password,omitempty" db:"password" binding:"required"`
	RoleId    string    `json:"role_id,omitempty" db:"role_id" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type EmployeePlants struct {
	Id         string    `db:"id" json:"id"`
	EmployeeId string    `db:"employee_id" json:"employee_id" binding:"required"`
	PlantId    string    `db:"plant_id" json:"plant_id" binding:"required"`
	Position   string    `db:"position" json:"position" binding:"required"`
	JoinDate   string    `db:"join_date" json:"join_date" binding:"required"`
	CreatedAt  time.Time `db:"create_at" json:"create_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type PlantWitEmplResponse struct {
	Plants            Plants
	EmplPlantResponse []EmplPlantResponse
}

type EmplPlantResponse struct {
	Id        string    `json:"id"`
	Fullname  string    `json:"fullname"`
	Position  string    `json:"position"`
	PlantName string    `json:"plant_name" db:"plant_name"`
	JoinDate  time.Time `json:"join_date" db:"join_date"`
}
