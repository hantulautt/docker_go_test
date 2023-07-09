### Framework & Library
- [GoFiber v2](https://gofiber.io/)
- [Gorm v1](https://gorm.io/)

### Database
```
Menggunakan database Mysql
```

### Architecture
**Controller -> Service -> Repository**

### Instalasi Project

- Docker
```
docker-compose up -d
```

### Endpoint Api
| Method  | Endpoint                 | Keterangan                            |
|---------|--------------------------|---------------------------------------|
| **GET** | insert-data              | Insert data pada database             |
| **GET** | api/whois/xxx.xxx.xxx/xx | Menampilkan data whois berdasarkan IP |


### Keterangan

- Apabila cron job tidak berfungsi silahkan hit endpoit **insert-data** untuk melakukan insert dari file apnic.db.inetnum kedalam database (Proses ini berjalan secara running background menggunakan go routine)
```
http://localhost:8081/insert-data
```
- Untuk menampilkan data gunakan contoh dibawah ini (untuk ip bisa dilihat pada database apabila telah berhasil insert data)
```
http://localhost:8081/api/whois/202.14.146.0/24
```
- Untuk mengakses database menggunakan adminer :
  - User : user
  - Password : password
```
http://localhost:8282/
```