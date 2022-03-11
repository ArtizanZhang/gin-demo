package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	db.Model(&Tag{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func ExitTagName(tagName string) bool {
	var tag Tag
	db.Model(&Tag{}).Where(map[string]interface{}{"name": tagName}).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false

}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true

}
