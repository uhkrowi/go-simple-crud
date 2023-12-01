package model

import "time"

type BaseModel struct {
	TglCreated time.Time `gorm:"column:tgl_created;type:datetime;<-:false" json:"tgl_created" search:"tgl_created" swaggerignore:"true"`
	TglUpdated time.Time `gorm:"column:tgl_updated;type:datetime;<-:false" json:"tgl_updated" search:"tgl_updated" swaggerignore:"true"`
}

type MetaData struct {
	TotalData int64 `json:"total_data"`
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	PrevPage  *int  `json:"prev_page"`
	NextPage  *int  `json:"next_page"`
	LastPage  int   `json:"last_page"`
}
