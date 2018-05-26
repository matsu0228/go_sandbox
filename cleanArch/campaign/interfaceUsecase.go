package campaign

import (
	entities "github.com/matsu0228/go_exampels/cleanArch/common/entities"
)

// Usecase is interface
type Usecase interface {
	GetCampaign(id int) (entities.Campaign, error)
	AddCampaign(text, title string) error
	SetCampaign(id int, text string) error
}
