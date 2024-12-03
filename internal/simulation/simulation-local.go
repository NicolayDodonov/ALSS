package simulation

// Проверяет достиг ли мир возраста выхода из симуляции 3 раза подряд.
// Возвращаем true, если надо ещё продолжать обучение.
func check(count *int, age int) bool {
	if age < FinalAgeTrain {
		// Если не достигли, обнуляем счётчи
		*count = 0
		return true
	} else {
		*count++
		if *count >= 3 {
			return false
		}
		return true
	}
}
