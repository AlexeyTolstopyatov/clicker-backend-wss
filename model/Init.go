package model

// MIND: Десериализация запроса будет успешна
// только если все поля объекта, в который будет переводиться
// тело запроса, будут идентичны с образцом тела запроса
//
// Чтобы служебные слова не совпали с тегами в описательной модели
// запроса, здесь используется (`json:"property"`)
// type Init struct {
//	Type string `json:"type"`
//  Team string `json:"team"`
//}

type Init struct {
	Type            string `json:"type"`
	BatteryCapacity int    `json:"battery_capacity"`
	BatteryCharge   int    `json:"battery_charge"`
	Clicks          int    `json:"clicks"`
	Team            string `json:"team"`
}
