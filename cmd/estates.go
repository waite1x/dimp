package cmd

import (
	"dexp/pkg/file"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Estate struct {
	ID                 int32  `json:"id" gorm:"column:id;type:int;primary_key;autoIncrement:false"`
	Name               string `json:"name" gorm:"column:name;type:varchar(255)"`
	Link               string `json:"link" gorm:"column:link;type:varchar(255)"`
	Description        string `json:"description" gorm:"column:description;type:varchar(255)"`
	CoverImage         string `json:"cover_image" gorm:"column:coverImage;type:varchar(255)"`
	CreatedAt          int64  `json:"created_at" gorm:"column:createdAt;not null;autoCreateTime:milli"`
	UpdatedAt          int64  `json:"updated_at" gorm:"column:updatedAt;autoUpdateTime:milli"`
	DeletedAt          int64  `json:"deleted_at" gorm:"column:deletedAt"`
	OwnerWalletAddress string `json:"owner_wallet_address" gorm:"column:ownerWalletAddress;type:varchar(255)"`
	Geography          string `json:"geography" gorm:"column:geography;type:text"`
	Model              string `json:"model" gorm:"column:model;type:text"`
}

type Data struct {
}

func EstatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estates",
		Short: "estates",
		Long:  "estates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEstates(cmd)
		},
	}
	cmd.PersistentFlags().StringP("dbconn", "d", "", "数据库连接字符串")
	cmd.PersistentFlags().StringP("file", "f", "", "导入的数据文件")
	cmd.MarkFlagRequired("dbconn")
	cmd.MarkFlagRequired("file")
	return cmd
}

func runEstates(cmd *cobra.Command) error {
	connStr, err := cmd.Flags().GetString("dbconn")
	if err != nil {
		return err
	}
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	if connStr == "" {
		return errors.New("数据库连接字符串不能为空")
	}
	data := make(map[string]interface{})
	err = file.ReadJson(filePath, &data)
	if err != nil {
		return err
	}
	estates, err := parse(data)
	if err != nil {
		return err
	}
	insertDb(connStr, estates)
	return nil
}

func parse(data map[string]interface{}) ([]Estate, error) {
	features := data["features"].([]interface{})
	estates := make([]Estate, 0, len(features))
	for _, feature := range features {
		item := feature.(map[string]interface{})
		estate := Estate{
			ID: int32(item["id"].(float64)),
		}
		geometry, err := json.Marshal(item["geometry"])
		if err != nil {
			return nil, err
		}
		estate.Geography = string(geometry)

		estates = append(estates, estate)
	}
	return estates, nil
}

func insertDb(connStr string, data []Estate) error {
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	//err = db.AutoMigrate(&Estate{})
	if err != nil {
		return err
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).
		Model(&Estate{}).
		Create(data)
	return nil
}
