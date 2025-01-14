package dto

type Path struct {
	Id     int    `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Guid string `json:"guid"`
    Description   string `json:"description"`
}