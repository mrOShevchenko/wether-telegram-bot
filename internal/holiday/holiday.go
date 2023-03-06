package holiday

import "time"

type Country struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Flag     string `json:"flag"`
	Holiday  string `json:"holiday"`
}

type HolidayData struct {
	Australia Country `json:"Australia"`
	Ukraine   Country `json:"Ukraine"`
	China     Country `json:"China"`
	Canada    Country `json:"Canada"`
	Georgia   Country `json:"Georgia"`
	France    Country `json:"France"`
	Date
}

type Date struct {
	Today time.Time
}