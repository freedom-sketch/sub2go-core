package utils

import "time"

func Greeting() string {
	hour := time.Now().Hour()

	var timeOfDay string
	switch {
	case hour >= 5 && hour < 12:
		timeOfDay = "Доброе утро"
	case hour >= 12 && hour < 17:
		timeOfDay = "Добрый день"
	case hour >= 17 && hour < 22:
		timeOfDay = "Добрый вечер"
	default:
		timeOfDay = "Доброй ночи"
	}

	return timeOfDay
}
