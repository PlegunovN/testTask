#!/bin/bash
# Скрипт для нагрузочного тестирования gRPC сервиса с использованием ghz.
# Использование:
#   ./grpc_load_test.sh <target_address:port> <call_method> <number_of_requests> <concurrency> [payload_file]
#
# Пример:
#   ./grpc_load_test.sh localhost:50051 pb.FileService/UploadFile 10000 50 payload.json
#
# Для работы скрипта необходимо, чтобы ghz был установлен и доступен в PATH.
# Можно установить ghz, следуя инструкциям в https://github.com/bojand/ghz

TARGET=${1:-"localhost:50051"}
CALL=${2:-"pb.FileService/UploadFile"}
NUMBER=${3:-10000}
CONCURRENCY=${4:-50}
PAYLOAD_FILE=${5:-"payload.json"}

if [ ! -f "$PAYLOAD_FILE" ]; then
  echo "Файл с полезной нагрузкой '$PAYLOAD_FILE' не найден!"
  exit 1
fi

echo "Запуск нагрузочного тестирования gRPC сервиса:"
echo "  Target:      $TARGET"
echo "  Метод:       $CALL"
echo "  Количество запросов:    $NUMBER"
echo "  Параллелизм: $CONCURRENCY"
echo "  Файл payload: $PAYLOAD_FILE"

ghz --insecure \
    --proto=pkg/pb/file.proto \
    --call="$CALL" \
    -n "$NUMBER" \
    -c "$CONCURRENCY" \
    -d @"$PAYLOAD_FILE" \
    "$TARGET"

