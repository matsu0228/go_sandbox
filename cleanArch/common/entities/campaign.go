package entities

// Campaign is struct on entities
type Campaign struct {
	ID    int
	Title string `db:"sTitle"`
	Text  string `db:"sText"`
}
