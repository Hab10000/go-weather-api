# go-weather-api # 🌤️ Go Weather API Wrapper

This is a simple weather API written in Go that fetches real-time weather data from OpenWeatherMap and returns it in JSON format. It supports Redis caching and `.env` configuration.

## 🚀 Features

- Fetch weather by city name
- Caches responses using Redis
- Loads configuration from `.env` file
- Returns clean JSON API response

## 📦 Requirements

- Go 1.18+
- Redis server
- `.env` file with:
  ```env
  WEATHER_API_KEY=your_api_key
  REDIS_ADDR=localhost:6379
  REDIS_PASSWORD=
  REDIS_DB=0
