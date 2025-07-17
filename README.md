# 🧩 Projekt: URL Shortener API (mit Go)

## 🎯 Ziel

Entwickle eine **REST-API in Go**, die **lange URLs in kurze Codes umwandelt** und bei Aufruf des Codes automatisch weiterleitet.
Optional: Admin-Interface und User-Auth.

---

## 🔧 Tech Stack

- **Sprache:** Go (Golang)
- **Web Framework:** `net/http`, optional `gorilla/mux` oder `fiber`
- **Datenbank:** In-Memory (für Start), optional SQLite / PostgreSQL
- **Persistenz:** optional JSON-Datei
- **(Optional) Auth:** JWT für Benutzer
- **(Optional) Frontend:** kleines Web-UI oder Swagger-Doku

---

## 📚 Anforderungen (MVP)

### 🔹 Endpunkt: `POST /api/shorten`

- **Beschreibung:** Erzeugt eine Kurz-URL
- **Request Body (JSON):**

```json
{
  "url": "https://example.com/super/lange/url"
}
```

- **Antwort (JSON):**

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

- **Verhalten:**
  - Die Kurz-URL wird als zufälliger 6–8-stelliger Code generiert.
  - Mehrfacheinsendungen derselben URL dürfen denselben oder neuen Code erhalten (je nach Modus).
  - Validierung: Nur gültige URLs akzeptieren.

### 🔹 Endpunkt: `GET /{shortcode}`

- **Beschreibung:** Leitet zur echten URL weiter (HTTP 301/302)
- **Beispiel:**
  `GET /abc123 → redirect → https://example.com/super/lange/url`

### 🔹 Endpunkt: `GET /api/info/{shortcode}`

- **Beschreibung:** Gibt Infos zur Short-URL zurück
- **Antwort:**

```json
{
  "original_url": "https://example.com/super/lange/url",
  "short_url": "http://localhost:8080/abc123",
  "created_at": "2025-07-15T12:34:56Z",
  "clicks": 42
}
```

### 🔹 Error Handling

- `404` bei nicht gefundenem Code
- `400` bei ungültiger URL
- Saubere JSON-Fehlermeldungen

---

## 🗃️ Speicheroptionen

### Stufe 1 (Pflicht):

- **In-Memory Map\[string]URLInfo**

```go
type URLInfo struct {
    OriginalURL string
    CreatedAt   time.Time
    Clicks      int
}
```

### Stufe 2 (optional):

- Persistenz in JSON-Datei beim Herunterfahren und Laden beim Start

### Stufe 3 (fortgeschritten):

- SQLite oder PostgreSQL mit `sqlx` oder `gorm`

---

## 🧪 Tests

- Schreibe Unit Tests für:
  - URL-Validierung
  - Shortcode-Generierung
  - API-Handler

- Nutze `go test`, z. B.:

```bash
go test ./...
```

---

## 🚀 Erweiterungen (für Bonuspunkte im Portfolio)

| Feature          | Beschreibung                                            |
| ---------------- | ------------------------------------------------------- |
| 🔒 Auth          | Registrieren + Einloggen (JWT), User hat eigene Links   |
| 🌐 Rate Limiting | Pro IP oder Benutzer nur X Shorten-Anfragen pro Minute  |
| 📈 Analytics     | Endpunkt mit Statistik (Clicks pro Tag, Top URLs, etc.) |
| 🧼 Vanity-URLs   | Wunsch-Codes (z. B. `myshop`) statt Zufalls-Strings     |
| 📁 Admin-Panel   | Web-GUI mit Go Template, React oder Svelte              |
| 📜 Swagger       | OpenAPI-Spezifikation für alle Endpunkte                |

---

## 🧱 Projektstruktur (Vorschlag)

```text
urlshortener/
├── main.go
├── handlers/
│   └── shortener.go
├── router/
│   └── routes.go
├── storage/
│   └── memory.go
├── utils/
│   └── validate.go
├── tests/
│   └── handler_test.go
├── go.mod
```

---

## 🧠 Präsentationsideen fürs Portfolio

- Zeige **Live-Demo** mit `ngrok` oder auf `fly.io`
- Dokumentiere mit `README.md` und Screenshots/Postman
- Baue eine kleine `curl`- oder Swagger-Doku ein
- Zeige gute **Projektstruktur + Testabdeckung**
- Link zu GitHub-Repo mit sauberem Commit-Verlauf
