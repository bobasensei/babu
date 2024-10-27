package list

import (
	"errors"
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
		Use:   "list",
		Short: "List pages in storage",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbURL := os.Getenv("BABU_DATABASE")
			if dbURL == "" {
				return errors.New("BABU_DATABASE should be set to a Postgres URL (postgres://USER:PASSWORD@HOST:PORT/DATABASE)")
			}
			db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
			if err != nil {
				return err
			}
			if structured {
				var pages []models.Content
				response := db.WithContext(cmd.Context()).Find(&pages)
				if response.Error != nil {
					return response.Error
				}
				for _, p := range pages {
					fmt.Printf("%s\n", p.Id)
				}
			} else {
				var pages []models.Article
				response := db.WithContext(cmd.Context()).Find(&pages)
				if response.Error != nil {
					return response.Error
				}
				for _, p := range pages {
					fmt.Printf("%s\n", p.Id)
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&structured, "structured-contents", false, "fetch structured contents")
	return cmd
}
