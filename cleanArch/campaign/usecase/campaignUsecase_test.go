package usecase_test

import (
	"log"
	"testing"

	"github.com/matsu0228/go_exampels/cleanArch/campaign/usecase"
	"github.com/matsu0228/go_exampels/cleanArch/common/entities"
)

type MockRepos struct {
}

// GetCampaign is mock
func (m MockRepos) GetCampaign(id int) (entities.Campaign, error) {
	log.Println("test GetCampaign() input=", id)
	return entities.Campaign{ID: 1, Title: "title", Text: "text"}, nil
}

// AddCampaign is mock
func (m MockRepos) AddCampaign(text, title string) error {
	log.Println("test AddCampaign() input=", text, title)
	return nil
}

// SetCampaign is mock
func (m MockRepos) SetCampaign(id int, text string) error {
	log.Println("test SetCampaign() input=", id, text)
	return nil
}

// TestSet is tester of Set
func TestSet(t *testing.T) {
	crMock := MockRepos{}
	cu := usecase.NewCampaignUsecase(crMock)
	err := cu.Set(2915, "updated title test")
	if err != nil {
		t.Errorf("Set err= ")
	}
}
