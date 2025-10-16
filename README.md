# Vietnam Location API

API RESTful cung cấp thông tin về các tỉnh thành và xã phường của Việt Nam.

## Cài đặt

```bash
go mod download
```

## Chạy ứng dụng

```bash
go run cmd/main.go
```

Server sẽ chạy trên port 8080.

## API Endpoints

### Lấy danh sách tỉnh thành

```
GET /provinces
```

Trả về danh sách tất cả các tỉnh thành của Việt Nam.

### Lấy danh sách xã phường theo mã tỉnh

```
GET /wards/?province_code={code}
```

Trả về danh sách xã phường của tỉnh thành có mã được chỉ định.

**Parameters:**
- `province_code` (required): Mã tỉnh thành

## Cấu trúc dữ liệu

### Province
```json
{
  "code": "01",
  "fullname": "Thành phố Hà Nội",
  "name": "Hà Nội", 
  "slug": "ha-noi",
  "type": "Thành phố"
}
```

### Ward
```json
{
  "code": "00001",
  "name": "Phúc Xá",
  "fullname": "Phường Phúc Xá", 
  "slug": "phuc-xa",
  "type": "Phường"
}
```

## Ví dụ sử dụng

```bash
# Lấy danh sách tỉnh thành
curl http://localhost:8080/provinces

# Lấy danh sách xã phường của Hà Nội (mã: 01)
curl http://localhost:8080/wards/?province_code=01
```

## Công nghệ sử dụng

- Go 1.24.2
- Gin Web Framework
- JSON data files