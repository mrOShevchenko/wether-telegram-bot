package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"task2.3.3/internal/config"
	"time"
)

type CountryData struct {
	Country  string `json:"country"`
	Location string `json:"location"`
	Flag     string
	Name     string `json:"name"`
}

type HolidayData struct {
	Australia CountryData `json:"Australia"`
	Ukraine   CountryData `json:"Ukraine"`
	China     CountryData `json:"China"`
	Canada    CountryData `json:"Canada"`
	Georgia   CountryData `json:"Georgia"`
	France    CountryData `json:"France"`
	Date
}

type Date struct {
	Today time.Time
}

func newHoliday() *HolidayData {
	defaultHolidays := &HolidayData{
		Australia: CountryData{
			Country:  "AU",
			Location: "Australia",
			Flag:     "🇦🇺",
			Name:     "Not any Holiday today",
		},
		Ukraine: CountryData{
			Country:  "UA",
			Location: "Ukraine",
			Flag:     "🇺🇦",
			Name:     "Not any Holiday today",
		},
		China: CountryData{
			Country:  "CN",
			Location: "China",
			Flag:     "🇨🇳",
			Name:     "Not any Holiday today",
		},
		Canada: CountryData{
			Country:  "CA",
			Location: "Canada",
			Flag:     "🇨🇦",
			Name:     "Not any Holiday today",
		},
		Georgia: CountryData{
			Country:  "GE",
			Location: "Georgia",
			Flag:     "🇬🇪",
			Name:     "Not any Holiday today",
		},
		France: CountryData{
			Country:  "FR",
			Location: "France",
			Flag:     "🇫🇷",
			Name:     "Not any Holiday today",
		},
		Date: Date{time.Now()},
	}
	return defaultHolidays
}

func (tg *telegramBot) holidayRequest() (*HolidayData, error) {
	cfg := tg.c.NewConfig()
	countries := []string{"AU", "UA", "CN", "CA", "GE", "FR"}
	countryMap := make(map[string]CountryData)

	for _, country := range countries {
		time.Sleep(time.Second * 1)
		holiday, err := updateHolidays(cfg, country)
		if err != nil {
			return nil, errors.Wrapf(err, "can't updateHolidays %+v")
		}
		countryMap[country] = *holiday
	}

	holidayData := &HolidayData{
		Australia: countryMap["AU"],
		Ukraine:   countryMap["UA"],
		China:     countryMap["CN"],
		Canada:    countryMap["CA"],
		Georgia:   countryMap["GE"],
		France:    countryMap["FR"],
	}

	return holidayData, nil
}

func updateHolidays(cfg *config.Config, country string) (*CountryData, error) {
	url := fmt.Sprintf("%vapi_key=%vcountry=%s&year=%d&month=%d&day=%d", cfg.AbstractApiBaseUrl,
		cfg.AbstractApiTokenEnv, country, time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Printf("\n url %v \n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get response from url, %+w")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "can't ReadAll Response body %+w")
	}

	var countryDataResp []CountryData
	err = json.Unmarshal(body, &countryDataResp)
	if err != nil {
		return nil, errors.Wrapf(err, "error in Unmarshal body %+w")
	}

	if len(countryDataResp) == 0 {
		return &CountryData{Name: "Not any holiday today"}, nil
	}
	fmt.Printf("\n\n countryDataResp : %v \n\n", countryDataResp)
	cdResp := countryDataResp[0]
	cdResp.Country = country
	cdResp.Location = resp.Header.Get("Location")

	return &cdResp, nil

}

func (tg *telegramBot) holidayRequestforCountry(country string) (string, error) {
	cfg := tg.c.NewConfig()

	url := fmt.Sprintf("%vapi_key=%vcountry=%s&year=%d&month=%d&day=%d", cfg.AbstractApiBaseUrl,
		cfg.AbstractApiTokenEnv, country, time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Printf("\n url %v \n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "error with Get response %+w")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "error in ReadAll %+w")
	}
	var countryValue []CountryData
	err = json.Unmarshal(body, &countryValue)
	if err != nil {
		return "", errors.Wrapf(err, "cann't Unmrashul %+w")
	}

	holidayName := "Not any holiday today :("

	if len(countryValue) > 0 {
		c := countryValue[0]
		holidayName = c.Name
		return holidayName, nil
	}
	return holidayName, nil

}
