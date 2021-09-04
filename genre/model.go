package genre

type Genre struct {
  ID   uint `json:"id"`
	Name string `gorm:"unique_index;not null" json:"name"`
}
