# Transactions API

> Dokumentasi endpoint yang berhubungan dengan transaksi (checkout dan laporan).

## Endpoints

1. Checkout

- Method: `POST`
- Endpoint: `/api/checkout`
- Deskripsi: Membuat transaksi baru (checkout) dengan beberapa item sekaligus. Stok produk akan dikurangi.

Request body (JSON):

```json
{
  "items": [
    {"product_id": 1, "quantity": 2},
    {"product_id": 3, "quantity": 1}
  ]
}
```

Response (200): `Transaction` object

```json
{
  "id": 10,
  "total_amount": 75000,
  "created_at": "2026-02-07T10:12:00Z",
  "details": [
    {"id": 101, "transaction_id": 10, "product_id": 1, "product_name": "Kopi", "quantity": 2, "subtotal": 50000},
    {"id": 102, "transaction_id": 10, "product_id": 3, "product_name": "Gula", "quantity": 1, "subtotal": 25000}
  ]
}
```

Curl example:

```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{"items":[{"product_id":1,"quantity":2},{"product_id":3,"quantity":1}]}'
```

Errors:
- 400 Bad Request: request body tidak valid
- 500 Internal Server Error: error pada server/DB (mis. product id tidak ditemukan, stock tidak cukup, dsb.)

---

2. Laporan Hari Ini

- Method: `GET`
- Endpoint: `/api/report/hari-ini`
- Deskripsi: Mengembalikan ringkasan transaksi untuk hari ini.

Response (200): `Summary` object

```json
{
  "total_revenue": 250000,
  "total_transaksi": 12,
  "produk_terlaris": {"nama":"Kopi","qty_terjual":20}
}
```

Curl example:

```bash
curl http://localhost:8080/api/report/hari-ini
```

---

3. Laporan Periode

- Method: `GET`
- Endpoint: `/api/report`
- Query params: `startDate` dan `endDate` (format `YYYY-MM-DD`). Jika tidak diberikan, default adalah hari ini (startDate) dan hari berikutnya (endDate).
- Deskripsi: Mengembalikan ringkasan transaksi pada rentang tanggal yang diberikan.

Contoh:

```
GET /api/report?startDate=2026-02-01&endDate=2026-02-07
```

Response sama seperti `Summary` di atas.

---

## Model JSON (ringkasan)

- `CheckoutRequest`:

```json
{
  "items": [
    {"product_id": 1, "quantity": 2}
  ]
}
```

- `Transaction`:

```json
{
  "id": 1,
  "total_amount": 50000,
  "created_at": "2026-02-07T10:00:00Z",
  "details": [
    {"id": 1, "transaction_id": 1, "product_id": 1, "product_name": "Kopi", "quantity": 2, "subtotal": 50000}
  ]
}
```

- `Summary`:

```json
{
  "total_revenue": 500000,
  "total_transaksi": 20,
  "produk_terlaris": {"nama":"Kopi","qty_terjual":50}
}
```

## Catatan Implementasi

- Endpoint checkout men-trigger pengurangan `stock` pada tabel `products` dan menyimpan data ke tabel `transactions` dan `transaction_details` dalam satu transaksi database (ACID).
- `Summary` menghitung total transaksi dan revenue dengan meng-query tabel `transactions` lalu menentukan produk terlaris dari `transaction_details`.
- Format tanggal untuk `startDate`/`endDate` harus `YYYY-MM-DD`.
