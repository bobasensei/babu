package get

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bobasensei/babu/pkg/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Cmd() *cobra.Command {
	var structured bool
	cmd := &cobra.Command{
		Use:   "get PAGE",
		Short: "Get a page from storage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			dbURL := os.Getenv("BABU_DATABASE")
			db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
			if err != nil {
				return err
			}
			if structured {
				var page models.Content
				result := db.First(&page, "id = ?", id)
				if result.Error != nil {
					return result.Error
				}
				v := page.Document.Get()
				b, err := json.MarshalIndent(v, "", "  ")
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", string(b))
			} else {
				var page models.Article
				result := db.First(&page, "id = ?", id)
				if result.Error != nil {
					return result.Error
				}
				v := page.Document.Get()
				b, err := json.MarshalIndent(v, "", "  ")
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", string(b))
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&structured, "structured-contents", false, "fetch structured contents")
	return cmd
}
