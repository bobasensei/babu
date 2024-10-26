package list

import (
	"fmt"
	"os"

	"github.com/bobasensei/babu/pkg/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pages in storage",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbURL := os.Getenv("BABU_DATABASE")
			db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
			if err != nil {
				return err
			}
			var pages []models.Page
			response := db.WithContext(cmd.Context()).Find(&pages)
			if response.Error != nil {
				return response.Error
			}
			for _, p := range pages {
				fmt.Printf("%s\n", p.Id)
			}
			return nil
		},
	}
	return cmd
}
