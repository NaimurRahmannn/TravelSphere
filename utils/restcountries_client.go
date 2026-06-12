package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"TravelSphere/models"
)

var restCountriesBase = "https://api.restcountries.com/countries/v5"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// v5 lists at most 100 records per request
const restCountriesPageLimit = 100

//response structure
type v5Response struct {
	Data struct {
		Objects []v5Country `json:"objects"`
		Meta    struct {
			Total int  `json:"total"`
			Count int  `json:"count"`
			More  bool `json:"more"`
		} `json:"meta"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}


type v5Country struct {
	Names struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"names"`
	Codes struct {
		Alpha2 string `json:"alpha_2"`
	} `json:"codes"`
	Capitals []struct {
		Name string `json:"name"`
	} `json:"capitals"`
	Population int    `json:"population"`
	Region     string `json:"region"`
	Subregion  string `json:"subregion"`
	Flag       struct {
		URLPNG      string `json:"url_png"`
		Description string `json:"description"`
	} `json:"flag"`
	Coordinates struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinates"`
	Currencies []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Languages []struct {
		Name    string `json:"name"`
		ISO6393 string `json:"iso639_3"`
	} `json:"languages"`
}

func (v v5Country) toRawCountry() models.RawCountry {
	raw := models.RawCountry{
		Population: v.Population,
		Region:     v.Region,
		Subregion:  v.Subregion,
		CCA2:       v.Codes.Alpha2,
	}
	raw.Name.Common = v.Names.Common
	raw.Name.Official = v.Names.Official
	raw.Flags.PNG = v.Flag.URLPNG
	raw.Flags.Alt = v.Flag.Description

	for _, cap := range v.Capitals {
		raw.Capital = append(raw.Capital, cap.Name)
	}

	raw.LatLng = []float64{v.Coordinates.Lat, v.Coordinates.Lng}

	if len(v.Currencies) > 0 {
		raw.Currencies = make(map[string]struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}, len(v.Currencies))
		for _, cur := range v.Currencies {
			raw.Currencies[cur.Code] = struct {
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
			}{Name: cur.Name, Symbol: cur.Symbol}
		}
	}

	if len(v.Languages) > 0 {
		raw.Languages = make(map[string]string, len(v.Languages))
		for _, lang := range v.Languages {
			key := lang.ISO6393
			if key == "" {
				key = lang.Name
			}
			raw.Languages[key] = lang.Name
		}
	}

	return raw
}


func fetchPage(reqURL string) (countries []models.RawCountry, more bool, err error) {
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, false, fmt.Errorf("build request failed: %w", err)
	}
	if key := beego.AppConfig.DefaultString("RESTCOUNTRIES_API_KEY", ""); key != "" {
		req.Header.Set("Authorization", "Bearer "+key)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, false, fmt.Errorf("country not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, fmt.Errorf("read body failed: %w", err)
	}

	var parsed v5Response
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, false, fmt.Errorf("decode failed: %w", err)
	}
	if len(parsed.Errors) > 0 {
		return nil, false, fmt.Errorf("api error: %s", parsed.Errors[0].Message)
	}

	countries = make([]models.RawCountry, 0, len(parsed.Data.Objects))
	for _, obj := range parsed.Data.Objects {
		countries = append(countries, obj.toRawCountry())
	}
	return countries, parsed.Data.Meta.More, nil
}

func GetAllCountries() ([]models.RawCountry, error) {
	var all []models.RawCountry
	for offset := 0; ; offset += restCountriesPageLimit {
		reqURL := fmt.Sprintf("%s?limit=%d&offset=%d", restCountriesBase, restCountriesPageLimit, offset)
		page, more, err := fetchPage(reqURL)
		if err != nil {
			return nil, err
		}
		all = append(all, page...)
		if !more {
			break
		}
	}
	return all, nil
}

func GetCountryByName(name string) ([]models.RawCountry, error) {
	reqURL := fmt.Sprintf("%s?q=%s", restCountriesBase, url.QueryEscape(name))
	countries, _, err := fetchPage(reqURL)
	return countries, err
}
