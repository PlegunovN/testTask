
# gRPC File Service

Этот проект представляет собой gRPC-сервис, реализованный на Go в стиле "Go Way". Сервис позволяет:
- **Загружать файлы (изображения):** Принимать бинарные файлы от клиента и сохранять их на диск.
- **Просматривать список файлов:** Возвращать список загруженных файлов с датами создания и обновления.
- **Скачивать файлы:** Отдавать файлы клиенту по запросу.
- **Ограничивать одновременные подключения:** 
  - 10 конкурентных запросов для операций загрузки/скачивания.
  - 100 конкурентных запросов для операций получения списка файлов.
- **Валидация запросов:** Используя [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate), входящие запросы проверяются согласно правилам, указанным в proto-файле.
- **Мониторинг и метрики:** Сбор метрик (общее число запросов, латентность, ошибки, количество параллельных запросов) с помощью Prometheus и визуализация их в Grafana.
- **Нагрузочное тестирование:** Скрипт для тестирования gRPC-сервиса с использованием [ghz](https://github.com/bojand/ghz) для проверки работы лоадбалансера.
- **Модульная архитектура:** Разделение логики обработки запросов на слой хэндлеров и бизнес-сервис, что позволит в будущем легко добавлять новые транспорты (например, HTTP/REST).

## Структура проекта

```
testTask
├─ .env                                 # config env
├─ .golangci.yml                        # linter
├─ Dockerfile
├─ README1.md
├─ api                                  # Протоколы с описанием сервиса и валидаторами (protoc-gen-validate)
│  ├─ file.proto
│  └─ validate
│     └─ validate.proto
├─ cmd
│  └─ server
│     ├─ main.go                        # Точка входа: запуск gRPC сервера, цепочка интерцепторов
│     └─ uploads                        # Папка куда падают/ читаются файлы
│        └─ 1.jpg
├─ docker-compose.yaml                  # Docker Compose: сервис, Prometheus, Grafana
├─ go.mod
├─ go.sum
├─ grafana                              # настройки grafana
│  └─ provisioning
│     ├─ dashboards
│     │  ├─ dashboards.yml
│     │  ├─ datasources
│     │  │  └─ prometheus.yaml
│     │  └─ grpc_dashboard.json
│     └─ datasources
├─ internal
│  ├─ configs                         # Обработка конфигов
│  │  └─ config.go
│  ├─ files
│  │  ├─ service.go                   # Бизнес-логика работы с файлами (загрузка, скачивание, список)
│  ├─ grpc
│  │  ├─ concurrency_Interceptor.go   # Интерцептор для ограничения одновременных запросов
│  │  ├─ metrics_interceptor.go       # Интерцептор для сбора метрик с использованием Prometheus
│  │  └─ validation_interceptor.go    # Интерцептор для валидации входящих запросов
│  ├─ handlers                        # Хэндлеры gRPC, делегирующие вызовы бизнес-сервису
│  │  ├─ download.go
│  │  ├─ handlers.go
│  │  ├─ list.go
│  │  └─ upload.go
│  └─ logger                          # ZAP logger
│     └─ logger.go
├─ makefile
├─ payload.json
│─ file.pb.go                         # Сгенерированные протофайлы
│─ file_grpc.pb.go
├─ prometheus                         # Настройки prometheus
│  └─ prometheus.yml
└─ test
   └─ grpc_load_test.sh               # Скрипт для нагрузочного тестирования gRPC сервиса с использованием ghz

   

```

## Требования

- **Go:** версия 1.18 или выше.
- **Protocol Buffers:** установлен protoc.
- **Плагины для protoc:** protoc-gen-go, protoc-gen-go-grpc, protoc-gen-validate.
- **Docker & Docker Compose:** для контейнеризации.
- **ghz:** для проведения нагрузочного тестирования gRPC сервиса.
- **zap:** для логирования.
- **Viper:** для обработки конфигов.

## Установка и запуск

### Генерация proto файлов

Proto-файл находится в каталоге `api/file.proto`. Для генерации Go-кода (с поддержкой валидации) выполните:

```bash
make proto
```

### Сборка проекта

Используйте Makefile для сборки:

```bash
make build
```

Бинарный файл будет создан в каталоге `bin/`.



### Запуск сервиса локально

После сборки запустите gRPC сервер:

```bash
./bin/server
```

Сервис будет доступен на порту **50051** (gRPC) и **2112** (метрики).

### Запуск через Docker Compose

Для развёртывания сервиса вместе с Prometheus и Grafana выполните:

```bash
docker-compose up -d
```

**Порты:**
- gRPC сервис: 50051
- Метрики сервиса: 2112
- Prometheus: 9090 (http://localhost:9090)
- Grafana: 3000 (http://localhost:3000; по умолчанию логин: admin/admin)

## Мониторинг и Метрики

Сервис использует интерцепторы для сбора метрик:
- **grpc_requests_total:** общее число запросов по методу.
- **grpc_requests_latency_seconds:** гистограмма латентности запросов.
- **grpc_requests_errors_total:** количество ошибок по методу и коду.
- **grpc_concurrent_requests:** текущее количество параллельных запросов.

Grafana настроена через provisioning:
- Конфигурация datasource (`grafana/provisioning/dashboards/datasources/prometheus.yaml`) указывает на Prometheus по адресу `http://prometheus:9090`.
- Готовый дашборд (`grafana/provisioning/dashboards/grpc_dashboard.json`) отображает метрики сервиса.



### Запуск нагрузочного теста

Выполните скрипт для нагрузочного тестирования:

```bash
./test/grpc_load_test.sh localhost:50051 pb.FileService/UploadFile 10000 50 payload.json 
```

Скрипт отправит 10,000 запросов с 50 параллельными соединениями на указанный метод.

## Dockerfile

Проект включает multi-stage Dockerfile для сборки и запуска сервиса в минимальном образе Alpine.  
См. файл [Dockerfile](./Dockerfile) для подробностей.


