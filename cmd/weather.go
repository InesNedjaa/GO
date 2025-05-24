/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

const ApiKey = "fb65ba1e04e3e67bc1a388d62d64d9e9"

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt   int64 `json:"dt"`
	Sys  struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}


type WeatherParam struct {
	Main        string
	Temperature float64
	Humidity    int
	WindSpeed   float64
}

type Coordinates struct {
	Name  string `json:"name"`
	LocalNames map[string]string   `json:"local_names"`
	Lat float64   `json:"lat"`
	Long float64  `json:"lon"`
	Country string  `json:"country"`
	State string     `json:"state"`

}
type LongLat struct {
	Long float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}


func ParseStringToLangLat(city string) (*LongLat, error) { 
	var coordinates []Coordinates 
	GeocordinatesParserBaseUrl := "http://api.openweathermap.org/geo/1.0/direct"
	url := fmt.Sprintf("%s?q=%s&limit=1&appid=%s", GeocordinatesParserBaseUrl, url.QueryEscape(city), ApiKey)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in parsing city name")
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in reading response body")
		return nil, err
	}
	
	err = json.Unmarshal(responseBody , &coordinates)
	if err != nil {
		fmt.Println("Error in marshing response body" , err)
		return nil, err
	}
	if len(coordinates) == 0 {return nil , nil }
	return &LongLat{Long: coordinates[0].Long, Lat: coordinates[0].Lat}, nil
    
}
func GetWeather(long float64 , lat float64)(*WeatherParam , error) {
	var WeatherResponse WeatherResponse
	WeatherBaseUrl := "https://api.openweathermap.org/data/2.5/weather"
	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=metric", WeatherBaseUrl, lat, long, ApiKey)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in getting weather information")
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in reading response body")
		return nil, err
	}
	
	err = json.Unmarshal(responseBody , &WeatherResponse)
	if err != nil {
		fmt.Println("Error in marshing response body")
		return nil, err
	}
	return &WeatherParam{Main: WeatherResponse.Weather[0].Main , Temperature:WeatherResponse.Main.Temp ,Humidity: WeatherResponse.Main.Humidity , WindSpeed: WeatherResponse.Wind.Speed }, nil
    


}

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Provide the weather of a giving city",
	Long:  `It allows you to know the weather of a giving city name.`,
	Run: func(cmd *cobra.Command, args []string) {
		city, err := cmd.Flags().GetString("city")
		if err != nil {
			fmt.Println("Error in getting city name")
			return
		}
		coordinates, err := ParseStringToLangLat(city)
		if err != nil || coordinates == nil {
			fmt.Println("Error in getting city coordinates, please check the city name")
			return
		}
		weather , err := GetWeather(coordinates.Long , coordinates.Lat)
		if err != nil {
			fmt.Println("Error in getting city weather")
			return
		}
		fmt.Println(".Main :" , weather.Main)
	    fmt.Println(".Temperature :" , weather.Temperature)
		fmt.Println(".Humidity :" , weather.Humidity)
		fmt.Println(".Wind Speed :" , weather.WindSpeed)

	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
	weatherCmd.Flags().String("city", " ", "City name")

}
