package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/matsu0228/go_exampels/cleanArch/campaign"
	entities "github.com/matsu0228/go_exampels/cleanArch/common/entities"
)

type mssqlCampaignRepositories struct {
	DB *sqlx.DB
}

// NewMssqlRepositories is consructor
func NewMssqlRepositories(db *sqlx.DB) campaign.Repositories {
	return &mssqlCampaignRepositories{DB: db}
}

// AddCampaign is func of insert sample
func (m mssqlCampaignRepositories) AddCampaign(text, title string) error {
	ctx := context.Background()
	err := m.DB.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := "INSERT INTO dxdb..Campaign  (sText, sTitle) VALUES ( @Text, @Title)"
	result, err := m.DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Text", text),
		sql.Named("Title", title),
	)
	fmt.Print("insert res=", result)
	return err
}

// SetCampaign is func of update
// https://docs.microsoft.com/ja-jp/azure/sql-database/sql-database-connect-query-go
func (m mssqlCampaignRepositories) SetCampaign(id int, text string) error {
	ctx := context.Background()
	err := m.DB.PingContext(ctx)
	if err != nil {
		// log.Fatal("Error pinging database: " + err.Error())
		return err
	}

	tsql := "UPDATE dxdb..Campaign Set sText = @Text WHERE ID = @ID"
	result, err := m.DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Text", text),
		sql.Named("ID", id),
	)
	fmt.Print("update res=", result)
	return err
}

// GetCampaign is func of selecting with id
//  -ref: http://lab.aratana.jp/entry/2016/12/17/golang-sqlx
func (m mssqlCampaignRepositories) GetCampaign(id int) (entities.Campaign, error) {
	ctx := context.Background()
	tsql := "SELECT Top 1 id, sTitle, sText FROM dxdb..Campaign WHERE id = @ID "

	rows, err := m.DB.QueryxContext(
		ctx,
		tsql,
		sql.Named("ID", id),
	)
	if err != nil {
		return entities.Campaign{}, err
	}
	fmt.Println("rows=", rows)
	defer rows.Close()

	camp := entities.Campaign{}
	for rows.Next() {
		err = rows.StructScan(&camp)
		if err != nil {
			return entities.Campaign{}, err
		}
	}
	return camp, err
}
