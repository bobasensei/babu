package store

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bobasensei/babu/pkg/models"
	"github.com/jackc/pgtype"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store FILE DBURL",
		Short: "store a file in postgres",
		Args:  cobra.ExactArgs(2),
		RunE:  action,
	}
	return cmd
}

func action(cmd *cobra.Command, args []string) error {
	file := args[0]

	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var docs []interface{}
	err = json.Unmarshal(b, &docs)
	if err != nil {
		return err
	}
	fmt.Printf("\n\nlen=%d\n", len(docs))

	dbURL := args[1]
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", db)

	err = db.AutoMigrate(&models.Page{})
	if err != nil {
		return err
	}

	for _, doc := range docs {
		name := doc.(map[string]interface{})["name"].(string)
		isPartOf := doc.(map[string]interface{})["is_part_of"].(map[string]interface{})
		identifier := isPartOf["identifier"].(string)
		if identifier == "enwiki" {
			jsonData := pgtype.JSONB{}
			b, err := json.Marshal(doc)
			if err != nil {
				return err
			}
			err = jsonData.Set(b)
			if err != nil {
				return err
			}
			page := &models.Page{
				Id:       name,
				Document: jsonData,
			}
			result := db.Save(page)
			if result.Error != nil {
				return result.Error
			}
		}
	}

	return nil
}
