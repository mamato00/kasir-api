# Kasir API

Kasir API adalah REST API yang dibangun menggunakan Go untuk penugasan project bootcamp Jago Golang Dasar - CodeWithUmam.

## Deployment Link : https://kasir-api-f.zeabur.app/

## Ada apa disini

- **CRUD Product**: CRUD operations untuk produk (Create, Read, Update, Delete)
- **CRUD Category**: CRUD operations untuk kategori (Create, Read, Update, Delete)
- **Checkout Transaction**: Checkout transaksi
- **Report Transaction**: Melihat ringkasan transaksi berdasarkan tanggal.
- **Health Check**: Endpoint untuk mengecek status API
- **Database Integration**: Terintegrasi dengan PostgreSQL

## Tech Stack

- **Language**: Go 1.25.6
- **Database**: PostgreSQL
- **Dependencies**:
  - `github.com/lib/pq` - PostgreSQL driver untuk Go
  - `github.com/spf13/viper` - Environment variable loader

## Instalasi & Setup

### Step-by-Step

1. **Clone repository**
```bash
git clone <repository-url>
cd kasir-api
```

2. **Setup environment variables**
Buat file `.env` di root directory:
```
DB_CONN=postgresql://username:password@host:port/database
PORT=8080
```

3. **Install dependencies**
```bash
go mod download
```

4. **Run server**
```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

### Build (Compile) untuk Linux & Windows

Anda bisa meng-compile aplikasi Go menjadi binary untuk platform target. Contoh-perintah di bawah menghasilkan binary statis untuk `linux` dan `windows`.

Build untuk Linux (amd64):
```bash
GOOS=linux GOARCH=amd64 go build -o bin/kasir-linux ./cmd/server
```

Build untuk Windows (amd64):
```bash
GOOS=windows GOARCH=amd64 go build -o bin/kasir-windows.exe ./cmd/server
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `DB_CONN`| PostgreSQL connection string |
| `PORT`   | Port server |

Sekian dan Terima Kasih.
