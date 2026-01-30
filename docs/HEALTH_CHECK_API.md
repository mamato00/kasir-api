# API Dokumentasi - Health Check Endpoint

Dokumentasi untuk Health Check endpoint yang digunakan untuk mengecek status API.

---

## Health Check

Endpoint ini digunakan untuk memverifikasi bahwa API sudah berjalan dan siap menerima requests.

**Request:**
```
GET /health
```

**Parameters:** None

**Headers:**
```
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "status": "OK",
  "message": "Api Running"
}
```

**Status Codes:**
- `200 OK` - API sedang berjalan dengan baik

**Example dengan curl:**
```bash
curl -X GET http://localhost:8080/health
```

**Output:**
```json
{"status":"OK","message":"Api Running"}
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/health')
  .then(response => response.json())
  .then(data => {
    console.log('Status:', data.status);
    console.log('Message:', data.message);
  })
  .catch(error => console.error('Error:', error));
```

**Example dengan axios:**
```javascript
const axios = require('axios');

axios.get('http://localhost:8080/health')
  .then(response => {
    console.log('Health Check Response:', response.data);
  })
  .catch(error => {
    console.error('Health Check Failed:', error.message);
  });
```

---

## Response Format

### Success Response

```json
{
  "status": "OK",
  "message": "Api Running"
}
```

**Field Descriptions:**

| Field | Type | Deskripsi |
|-------|------|-----------|
| status | string | Status API (OK = running) |
| message | string | Pesan deskriptif |

---

## Expected Behavior

| Scenario | Status | Response |
|----------|--------|----------|
| API berjalan normal | 200 | `{"status":"OK","message":"Api Running"}` |
| API tidak berjalan | Connection refused | N/A |
| Server error | 500 | Error message |

---

Last Updated: January 2026
