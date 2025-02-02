package model

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"golang.org/x/net/html/charset"
)

type WeatherForecast struct {
	Nombre   string `xml:"nome"`
	Previsao []struct {
		Dia    string `xml:"dia"`
		Tempo  string `xml:"tempo"`
		Maxima string `xml:"maxima"`
		Minima string `xml:"minima"`
		IUV    string `xml:"iuv"`
	} `xml:"previsao"`
}

type WaveForecast struct {
	Nombre string `xml:"nome"`
	Mañana []struct {
		Dia             string `xml:"dia"`
		Agitacion       string `xml:"agitacao"`
		Altura          string `xml:"altura"`
		Direcao         string `xml:"direcao"`
		Viento          string `xml:"vento"`
		DireccionViento string `xml:"vento_dir"`
	} `xml:"manha"`
	Tarde []struct {
		Dia             string `xml:"dia"`
		Agitacion       string `xml:"agitacao"`
		Altura          string `xml:"altura"`
		Direcao         string `xml:"direcao"`
		Viento          string `xml:"vento"`
		DireccionViento string `xml:"vento_dir"`
	} `xml:"tarde"`
	Noche []struct {
		Dia             string `xml:"dia"`
		Agitacion       string `xml:"agitacao"`
		Altura          string `xml:"altura"`
		Direcao         string `xml:"direcao"`
		Viento          string `xml:"vento"`
		DireccionViento string `xml:"vento_dir"`
	} `xml:"noite"`
}

func GetWeatherForecast(cityID string) (*WeatherForecast, error) {
	resp, err := http.Get(fmt.Sprintf("http://servicos.cptec.inpe.br/XML/cidade/%s/previsao.xml", cityID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var forecast WeatherForecast
	if err := decoder.Decode(&forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}

func GetWaveForecast(cityID, day string) (*WaveForecast, error) {
	resp, err := http.Get(fmt.Sprintf("http://servicos.cptec.inpe.br/XML/cidade/%s/dia/%s/ondas.xml", cityID, day))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var forecast WaveForecast
	if err := decoder.Decode(&forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}
