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
    --proto=api/file.proto \
    --call="$CALL" \
    -n "$NUMBER" \
    -c "$CONCURRENCY" \
    -d '{"filename": "1.txt", "data": "AWtVsd0qx3vo1MVVFtwPXP4NIq2hEBJ4tY=="}' \
    "$TARGET"
