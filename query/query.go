package query

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MenuItem struct {
	ID              string        `json:"id"`
	QtyRemaining    interface{}   `json:"qty_remaining"`
	OnlineTimeSteps []interface{} `json:"online_time_steps"`
}

type Response struct {
	MenuItems []MenuItem `json:"menu_items"`
}

func CheckAvailability(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println(resp)
		return false
	}

	// Check if any menu item has non-empty online_time_steps
	for _, item := range response.MenuItems {
		if len(item.OnlineTimeSteps) > 0 {
			return true
		}
	}
	return false
}
