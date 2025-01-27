package services

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"weatherAPI/pkg/model"
)

type WeatherForecast struct {
	XMLName  xml.Name `xml:"cidade"`
	Previsao []struct {
		Dia    string `xml:"dia"`
		Tempo  string `xml:"tempo"`
		Maxima string `xml:"maxima"`
		Minima string `xml:"minima"`
		IUV    string `xml:"iuv"`
	} `xml:"previsao"`
}

type WaveForecast struct {
	XMLName  xml.Name `xml:"cidade"`
	Previsao []struct {
		Dia     string `xml:"dia"`
		Altura  string `xml:"altura"`
		Direcao string `xml:"direcao"`
	} `xml:"previsao"`
}

func GetWeatherForecast(cityID string) (*WeatherForecast, error) {
	resp, err := http.Get(fmt.Sprintf("http://servicos.cptec.inpe.br/XML/cidade/%s/previsao.xml", cityID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var forecast WeatherForecast
	if err := xml.Unmarshal(body, &forecast); err != nil {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var forecast WaveForecast
	if err := xml.Unmarshal(body, &forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}

func SendNotification(user model.User, forecast *WeatherForecast, waveForecast *WaveForecast) {
	// Implementar la lógica para enviar la notificación al usuario
	// Esto puede ser un correo electrónico, una notificación push, etc.
	fmt.Printf("Enviando notificación a %s sobre el clima: %+v y olas: %+v\n", user.Email, forecast, waveForecast)
}

func ScheduleNotifications(db *sql.DB) {
	for {
		users, err := model.GetUsers(db)
		if err != nil {
			fmt.Println("Error obteniendo usuarios:", err)
			continue
		}

		for _, user := range users {
			if user.OptOut {
				continue
			}

			now := time.Now().Format("15:04")
			if now == user.NotificationTime {
				forecast, err := GetWeatherForecast("1234") // Reemplazar con el ID de la ciudad del usuario
				if err != nil {
					fmt.Println("Error obteniendo previsión climática:", err)
					continue
				}

				var waveForecast *WaveForecast
				if isCoastalCity("1234") { // Reemplazar con la lógica para determinar si es una ciudad costera
					waveForecast, err = GetWaveForecast("1234", "0")
					if err != nil {
						fmt.Println("Error obteniendo previsión de olas:", err)
						continue
					}
				}

				SendNotification(user, forecast, waveForecast)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

func isCoastalCity(cityID string) bool {
	// Implementar la lógica para determinar si una ciudad es costera
	return true
}
