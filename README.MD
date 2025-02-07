# 🚀 URL Shortener Service

A lightweight and efficient URL shortener built using Go, PostgreSQL, and Redis.
## Features
✅ Generate short URLs for long links  
✅ Retrieve original URLs from short links  
✅ Expiry time for links (optional)  
✅ Rate limiting to prevent abuse  
✅ Dockerized for easy deployment

---
## 🔧 Installation & Setup

### 1️⃣ Clone the repository
```sh
git clone https://github.com/ilmedova/url-shortener-go.git
cd url-shortener-go
docker compose up
```

## 📖 API Documentation

| Endpoint                  | Description              |
|---------------------------|--------------------------|
| POST /shorten             | 	Shorten a URL           |
| GET /:short_code          | 	Redirect to long URL    |
| GET  /stats/:short_code   | Get URL usage statistics |


## 📜 License
This project is licensed under the MIT License.
