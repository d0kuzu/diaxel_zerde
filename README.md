#Diaxel-analog

## Что в репозитории

- **Frontend**: Next.js + Tailwind, страницы — лендинг, фичи, аналитика с графиками, прайсинг, блог, контакты.
- **Backend**: микросервисы на Go — API Gateway, Auth Service, AI Service, Database Service, Telegram Service.
- **Протоколы**: gRPC/protobuf для связи сервисов.
- **Docker**: `docker-compose.yml` для запуска всего стенда.

## Как запустить (после клонирования)

### 1) Фронтенд
```sh
cd frontend
cp .env.example .env.local
# при желании: NEXT_PUBLIC_API_BASE_URL=http://localhost:8081
npm install
npm run dev
```
Откроется `http://localhost:3000`.

### 2) Бэкенд (чтобы аналитика работала с реальными данными)

#### а) Установить protoc и плагины (Windows)
```powershell
cd database-service/proto
.\install-protoc.ps1
```
Или скачай protoc с [GitHub releases](https://github.com/protocolbuffers/protobuf/releases) и добавь `bin` в PATH.

#### б) Сгенерировать protobuf
```powershell
cd database-service
protoc --plugin=protoc-gen-go=C:\Users\user\go\bin\protoc-gen-go.exe --plugin=protoc-gen-go-grpc=C:\Users\user\go\bin\protoc-gen-go-grpc.exe --go_out=. --go-grpc_out=. proto/database.proto
```
Или:
```powershell
.\proto\generate.ps1
```

#### в) Подтянуть Go‑зависимости
```sh
cd database-service
go mod tidy
```

#### г) Запустить всё через Docker
```sh
docker-compose up --build
```
API Gateway будет на `http://localhost:8081`.

### 3) Как проверить аналитику
- Если `NEXT_PUBLIC_API_BASE_URL` не задан — дашборд покажет мок‑данные.
- Если задан `http://localhost:8081` и бэкенд запущен — пойдут реальные запросы к `/api/analytics/*`.

## Структура

```
frontend/          # Next.js приложение
database-service/  # Go сервис + protobuf
api-gateway/       # Проксирование + JWT
ai-service/        # Логика AI и Telegram
auth-service/      # JWT‑авторизация
telegram-service/  # Вебхуки Telegram
docker-compose.yml # Все сервисы
```

## Дальше (что можно улучшить)

- Реальная авторизация (логин → JWT → хранение в cookie/localStorage).
- Эндпоинты `/api/analytics/*` в API Gateway (прокси в database-service gRPC).
- Тесты, CI, прод‑деплой.

---

**Если что‑то не запускается** — проверь, что:
- Node.js 20+ и Go 1.22+ установлены.
- `protoc` и плагины доступны в PATH.
- `.env.local` создан (можно пустой).
