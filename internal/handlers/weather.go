package handlers

import (
	"encoding/json"
	"net/http"
	"weatherAPI/pkg/model"

	"github.com/gorilla/mux"
)

func WeatherRouterHandlers(router *mux.Router) {
	router.HandleFunc("/weather/{cityID}", getWeatherForecast).Methods("GET")
	router.HandleFunc("/waves/{cityID}/{day}", getWaveForecast).Methods("GET")

}

func getWeatherForecast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityID := vars["cityID"]

	forecast, err := model.GetWeatherForecast(cityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forecast)
}

func getWaveForecast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityID := vars["cityID"]
	day := vars["day"]

	forecast, err := model.GetWaveForecast(cityID, day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forecast)
}
