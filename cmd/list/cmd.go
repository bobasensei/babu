package list

import (
	"log"
	"os"

	"github.com/bobasensei/babu/pkg/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var verbose bool
var token string

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pages in storage",
		Args:  cobra.NoArgs,
		RunE:  action,
	}
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "display intermediate information")
	return cmd
}

func action(cmd *cobra.Command, args []string) error {

	dbURL := os.Getenv("BABU_DATABASE")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}

	var pages []models.Page
	db.WithContext(cmd.Context()).Find(&pages)

	for i, p := range pages {
		log.Printf("%d %s\n", i, p.Id)
	}
	return nil
}
