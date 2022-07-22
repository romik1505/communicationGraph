**Установка:**
Для запуска сервиса потребуется docker 

**Первичный запуск:**
Скомпилируем сервис `make build`
Загрузим утилиту для миграций `make bin-deps`
Собираем docker образы `docker compose build --no-cache`
Запускаем контейнеры `docker compose up`
Прогоняем миграции `make db:up`
Сервис готов к работе.

Rest маршруты: 
1. Добавление записи о  коммуникации в журнал
	Метод  _Post_  localhost:8080/record
	
    Формат _входных данных:_  `{"target_id": "007f3b6b-38d6-4624-87f0-5d879a0f503e", "source_id": "3faef0ae-6aad-4038-ba7d-415ee1bde66c", "time": "2020-01-01T14:30:12Z" }`

	target_id, source_id: формат uuid v4
	
    time: RFC3999

2. Вывести граф коммуникаций пользователей.

Метод *Get* localhost:8080/graph 
