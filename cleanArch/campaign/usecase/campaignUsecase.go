package usecase

import (
	"github.com/matsu0228/go_exampels/cleanArch/campaign"
)

// CampaignUsecase is struct
type CampaignUsecase struct {
	campaignRepos campaign.Repositories
}

// NewCampaignUsecase is constructor
func NewCampaignUsecase(cr campaign.Repositories) CampaignUsecase {
	return CampaignUsecase{campaignRepos: cr}
}

// Set is business logic
func (c CampaignUsecase) Set(id int, text string) error {
	//
	// do something
	//
	err := c.campaignRepos.SetCampaign(id, "usecase work: "+text)
	return err
}

// // Get is func
// func (c Client) Get(id int) (Campaign, error) {
//
// 	// dosomething
//
// 	return c.Repository.GetCampaign(id)
// }
