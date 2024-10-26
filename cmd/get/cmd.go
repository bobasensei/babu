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

var verbose bool
var token string

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get PAGE",
		Short: "Get a page from storage",
		Args:  cobra.ExactArgs(1),
		RunE:  action,
	}
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "display intermediate information")
	return cmd
}

func action(cmd *cobra.Command, args []string) error {
	id := args[0]

	dbURL := os.Getenv("BABU_DATABASE")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}

	var page models.Page
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
	return nil
}
