# Sistem Manajemen Inventori dan Transaksi Penjualan

Proyek ini adalah sebuah **Sistem Manajemen Inventori dan Transaksi Penjualan** yang dibangun menggunakan **Golang** dengan framework **Gin**. Sistem ini mencakup pengelolaan produk, kategori, serta transaksi penjualan. API ini memungkinkan admin untuk menambah, mengupdate, dan menghapus data produk dan kategori, serta mencatat transaksi penjualan.


## Fitur Utama

1. **CRUD Produk**: Menambah, melihat, mengupdate, dan menghapus produk.
2. **CRUD Kategori**: Menambah dan melihat kategori produk.
3. **Transaksi Penjualan**: Mencatat penjualan, termasuk detail transaksi.
4. **Pembayaran**: Menangani pencatatan pembayaran untuk setiap transaksi.
5. **Autentikasi**:
   - Registrasi pengguna
   - Login pengguna
   - Logout pengguna
6. **Manajemen Pengguna**: Mengelola informasi akun pengguna.

## ERD 
![Image](https://github.com/user-attachments/assets/29aea898-0fe5-46b0-af00-fee02a9e5513)

## Struktur Proyek

```
ecommerce-api/
│── .github/
│   └── workflows/
│       └── ci.yml
│── config/
│── database/
│── handler/
│── middleware/
│── migration/
│── models/
│── repository/
│── router/
│── service/
│── utils/
│── docker-compose.yml
│── Dockerfile
│── go.mod
│── go.sum
│── main.go
```


## CRUD
https://documenter.getpostman.com/view/26920342/2sAYXBHftr


## Instalasi dan Menjalankan Proyek

1. Clone repository:
   ```sh
   git clone https://github.com/YusronAbi/ecommerce-api.git
   cd ecommerce-api
   ```
2. Jalankan database menggunakan Docker:
   ```sh
   docker-compose up -d
   ```
3. Instal dependensi:
   ```sh
   go mod tidy
   ```
4. Jalankan migrasi database:
   ```sh
   go run main.go migrate
   ```
5. Jalankan server:
   ```sh
   go run main.go
   ```

## API Endpoint

- **Autentikasi**
  - `POST /register` - Registrasi pengguna
  - `POST /login` - Login pengguna
  - `POST /logout` - Logout pengguna

- **Produk**
  - `GET /products` - Mendapatkan semua produk
  - `POST /products` - Menambah produk baru
  - `PUT /products/:id` - Mengupdate produk
  - `DELETE /products/:id` - Menghapus produk

- **Kategori**
  - `GET /categories` - Mendapatkan semua kategori
  - `POST /categories` - Menambah kategori baru

- **Transaksi**
  - `POST /transactions` - Membuat transaksi baru
  - `GET /transactions/:id` - Melihat detail transaksi

## Teknologi yang Digunakan

- Golang
- Gin Framework
- postgreSQL
- Docker
- JWT untuk autentikasi
- GORM untuk ORM database

