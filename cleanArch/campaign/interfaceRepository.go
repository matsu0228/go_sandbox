package campaign

import (
	entities "github.com/matsu0228/go_exampels/cleanArch/common/entities"
)

// Repositories is interface
type Repositories interface {
	GetCampaign(id int) (entities.Campaign, error)
	AddCampaign(text, title string) error
	SetCampaign(id int, text string) error
}
