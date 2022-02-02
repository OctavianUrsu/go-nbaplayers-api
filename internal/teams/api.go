package teams

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Team struct {
	TeamId       int    `json:"teamId"`
	Abbreviation string `json:"abbreviation"`
	TeamName     string `json:"teamName"`
	SimpleName   string `json:"simpleName"`
	Location     string `json:"location"`
}

func ReturnAllTeams(w http.ResponseWriter, r *http.Request) {
	// Open JSON file and if err => handle it
	teamsJson, err := os.Open("./internal/teams/teams.json")
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer teamsJson.Close()

	// Read the file as a byte array
	byteValue, _ := io.ReadAll(teamsJson)

	// Initialize the array that will store players
	var allTeams []Team

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allTeams)

	// Encode our array into the response
	json.NewEncoder(w).Encode(allTeams)

}
