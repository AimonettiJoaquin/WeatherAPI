package services

import (
	"database/sql"
	"fmt"
	"weatherAPI/pkg/model"
)

func SendNotification(user model.User, forecast *model.WeatherForecast, waveForecast *model.WaveForecast) {
	// Implementar la lógica para enviar la notificación al usuario
	// Esto puede ser un correo electrónico, una notificación push, etc.
	fmt.Printf("Enviando notificación a %s sobre el clima: %+v y olas: %+v\n", user.Email, forecast, waveForecast)
}

func ScheduleNotifications(db *sql.DB) {
	// Implementar la lógica para programar notificaciones
	fmt.Printf("Programando notificaciones...\n")
}
