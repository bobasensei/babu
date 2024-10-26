package initcmd

import (
	"errors"
	"os"

	"github.com/bobasensei/babu/pkg/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialize the babu datastore",
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
			return db.AutoMigrate(&models.Page{})
		},
	}
	return cmd
}
