package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bobasensei/babu/pkg/models"
	"github.com/jackc/pgtype"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var verbose bool
var token string

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch PAGE",
		Short: "Fetch a page from Wikipedia",
		Args:  cobra.ExactArgs(1),
		RunE:  action,
	}
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "display intermediate information")
	return cmd
}

func action(cmd *cobra.Command, args []string) error {
	page := args[0]
	token := os.Getenv("BABU_WIKIMEDIA")
	path := "https://api.enterprise.wikimedia.com/v2/structured-contents/" + page
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}
	log.Printf("fetching")
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("%s", res.Status)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(b))

	if false {
		return nil
	}
	var docs []interface{}
	err = json.Unmarshal(b, &docs)
	if err != nil {
		return err
	}
	fmt.Printf("\n\nlen=%d\n", len(docs))

	log.Printf("connecting")
	dbURL := os.Getenv("BABU_DATABASE")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", db)

	if false {
		err = db.AutoMigrate(&models.Page{})
		if err != nil {
			return err
		}
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
			log.Printf("saving")
			result := db.Save(page)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	log.Printf("done")
	return nil
}
