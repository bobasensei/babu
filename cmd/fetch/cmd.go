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

var text bool

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch PAGE",
		Short: "Fetch a page from Wikipedia",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			page := args[0]
			token := os.Getenv("BABU_WIKIMEDIA")

			if text {
				path := "https://api.enterprise.wikimedia.com/v2/articles/" + page
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
				return nil
			}

			path := "https://api.enterprise.wikimedia.com/v2/structured-contents/" + page
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				return err
			}
			fmt.Printf("fetching\n")
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

			fmt.Printf("response headers: %+v\n", res.Header)

			var docs []interface{}
			err = json.Unmarshal(b, &docs)
			if err != nil {
				return err
			}

			dbURL := os.Getenv("BABU_DATABASE")
			db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
			if err != nil {
				return err
			}

			for i, doc := range docs {
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
					fmt.Printf("%d: %s... saving\n", i, identifier)
					result := db.Save(page)
					if result.Error != nil {
						return result.Error
					}
				} else {
					fmt.Printf("%d: %s\n", i, identifier)
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&text, "text", "t", false, "fetch article text (and just return it without saving)")
	return cmd
}
