package domain

type RoleName string

const (
	// Автор – может публиковать контент
	Author RoleName = "author"

	// Рецензент – проверяет и комментирует материалы
	Reviewer RoleName = "reviewer"

	// Редактор – управляет контентом, модерирует публикации.
	Editor RoleName = "editor"

	// Читатель – имеет доступ только к просмотру.
	Reader RoleName = "reader"
)

type Role struct {
	ID   string   `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name RoleName `json:"name" gorm:"type:varchar(50);unique;not null"`
}

func (r RoleName) Value() string {
	return string(r)
}
