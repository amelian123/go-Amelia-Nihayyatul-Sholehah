## Section 21_ORM And Code Structure (MVC)

## ORM (Object Relational Mapping)
ORM dalam ilmu komputer adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek.

>> GoRM yaitu library fantastis yang digunakan untuk golang

>> Keuntungan ORM
   Adapun keuntungan ORM yaitu : 
   1. Query yang kurang berulang 
   2. Mengambil data secara otomatis ke objek siap pakai 
   3. Cara sederhana jika ingin menyaring data sebelum menyimpannya di database 
   4. Beberapa memiliki fitur cache query

>> Kerugian ORM 
   Adapun kerugian ORM yaitu : 
   1. Menambahkan lapisan dalam kode dan membebani proses overhead 
   2. Memuat data hubungan yang tidak diperlukan Query mentah yang kompleks dapat memakan waktu lama untuk menulis dengan alterr ORM (>10 table joins) 
   3. Fungsi SQL khusus yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi khusus (MySQL INDEKS KEKUATAN)

## Database Migration 
   Database migration (migrasi pada basis data) merupakan suatu cara untuk mengupdate versi database agar sesuai dengan perubahan versi aplikasi. Perubahan dapat berupa upgrade ke versi terbaru atau rollback ke versi sebelumnya.

>> Alasan Penggunaan Database Migration
   Berikut ini alasan penggunaan database migration yaitu :
   1. Kesederhanaan pembaruan database
   2. Kesederhanaan rollback database
   3. Melacak perubahan pada struktur database
   4. History struktur database ditulis pada kode 
   5. Selalu kompatibel dengan perubahan versi aplikasi

## Installasi GORM
   Langkahnya :
   1. GORM : Library ORM untuk Golang
      https://gorm.io/docs/
   2. GORM : Migrate
      https://gorm.io/docs/migration.html

## Installasi GORM dan MysSQL Driver 
   Langkahnya :
   a. go get -u gorm.io/gorm
   b. go get -u gorm.io/driver/mysql

## Cara Connect GO Ke Database Menggunakan GORM
    <username>:<password>@/<db_name>?charset=utf8&parseTime=True&loc=LocaL
>> Langkah selanjutnya agar connect ke database antara lain :
   1. Create InitDB() //untuk koneksi databse
   scriptnya yaitu :
   // database connection
func InitDB() {
  
  // declare struct config & variable connectionString

  var err error
  DB, err = gorm.Open("mysql", connectionString)
  if err != nil {
    panic(err)
  }
}
   2. Create User Schema dan InitialMigration()
   scriptnya :
   # src/main.go
... 
var (
  DB *gorm.DB
)

type User struct {
  gorm.Model
  Name     string `json:"name" form:"name"`
  Email    string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

func InitialMigration() {
  DB.AutoMigrate(&User{})
}
...
  3. Memanggil InitDB() dan InitialMigration()
  package main

// import library

func init() {
  InitDB()
  InitialMigration()
}

...
   4. Create GetUserController() -> untuk mendapatkan data users dari database (ORM) dan menampilkan data
   scriptnya yaitu :
   ...

// get all users
func GetUsersController(c echo.Context) error {
  var users []User

  if err := DB.Find(&users).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success get all users",
    "users":   users,
  })
}

..
   5. CreateUserController() -> untuk binding data users dari client dan menyimpan user ke database
   scriptnya yaitu :
...

// create new user
func CreateUserController(c echo.Context) error {
  user := User{}
  c.Bind(&user)

  if err := DB.Save(&user).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success create new user",
    "user":    user,
  })
}

...
  6. Routing -> mengidentifikasi routing API dengan memanggil controller
  scriptnya yaitu :
  ...

func main() {
  // create a new echo instance
  e := echo.New()

  // Route / to handler function
  e.GET("/users", GetUsersController)
  e.POST("/users", CreateUserController)

  // start the server, and log if it fails
  e.Start(":8000")
}

...
   
>> Melengkapi Semua Routes pada Rest API yaitu :
   e.GET("/users", GetUserController)
   e.GET("/users/:id", GetUserController)
   e.DELETE("/users/:id", DeleteUserController)
   e.PUT("/users/:id", UpdateUserController)
   e.POST("/users", CreateUserController)

## Code Structure MVC (Model View Controller)
   MVC adalah sebuah pola arsitektur dalam membuat sebuah aplikasi dengan cara memisahkan kode menjadi tiga bagian yang terdiri dari model, view, dan controller.
   >> Alur MVC yaitu :
      1. Proses pertama adalah view akan meminta data untuk ditampilkan dalam bentuk grafis kepada pengguna.
      2. Permintaan tersebut diterima oleh controller dan diteruskan ke model untuk diproses.
      3. Model akan mencari dan mengolah data yang diminta di dalam database
      4. Setelah data ditemukan dan diolah, model akan mengirimkan data tersebut kepada controller untuk ditampilkan di view.
      5. Controller akan mengambil data hasil pengolahan model dan mengaturnya di bagian view untuk ditampilkan kepada pengguna.

>> Manfaat MVC antara lain :
   1. Proses pengembangan aplikasi menjadi lebih efisien
   2. Penulisan kode menjadi lebih rapi
   3. Dapat melakukan testing dengan lebih mudah
   4. Perbaikan bug atau error lebih cepat untuk diselesaikan
   5. Mempermudah pemeliharaan

## Alasan Memerlukan Structure
   Untuk mencapai aplikasi modular Terapkan pemisahan perhatian.  Lebih sedikit konflik pada pembuatan versi.

## Langkah-langkah dalam Code Structuring
   1. Create Database di MySQL
   2. Create folder structure seperti folder confing, folder controller, folder lib (kemudian didalam folder lib ada folder database), folder models, dan folder routes
   3. Create Configuration masing-masing folder structure
   // Langkahnya yaitu :
      a. Folder Config -> di dalamnya dibuat file config.go
         scriptnya yaitu :  
         import (
            "fmt"
            "project/models"
            "gorm.io/driver/mysql"
            "gorm.io/gorm"
            )

            var DB *gorm.DB

            func InitDB() {

            config := map[string]string{
                "DB_Username": "root",
                "DB_Password": "123ABC4d.",
                "DB_Port":     "3306",
                "DB_Host":     "127.0.0.1",
                "DB_Name":     "training",
            }

            connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
                config["DB_Username"],
                config["DB_Password"],
                config["DB_Host"],
                config["DB_Port"],
                config["DB_Name"])

            var e error
            DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
            if e != nil {
                panic(e)
            }
            InitMigrate()
            }

            func InitMigrate() {
            DB.AutoMigrate(&models.Users{})
            }
      b. Create Models -> di dalam folder models ada file user.go
         scriptnya yaitu :
         package models
         import (
            "github.com/jinzhu/gorm"
        )

        type User struct {
        gorm.Model
        Name     string `json:"name" form:"name"`
        Email    string `json:"email" form:"email"`
        Password string `json:"password" form:"password"`
        }
      c. Craete Lib Database -> di dalam folder database diisi file user,go
         scriptnya yaitu :
         package database
        import (
            "project/config"
            "project/models"
        )
        func GetUsers() (interface{}, error) {
        var users []models.Users
        if e := config.DB.Find(&users).Error; e != nil {
            return nil, e
        }
        return users, nil
        }
      d. Create Controller -> di dalam folder controller ada file userController.go
         scriptnya yaitu : 
            package controllers
            import (
                "net/http"
                "project/lib/database"
                "project/models"

                "github.com/labstack/echo"
            )
            func GetUserControllers(c echo.Context) error {
            users, e := database.GetUsers()
            if e != nil {
                return echo.NewHTTPError(http.StatusBadRequest, e.Error())
            }
            return c.JSON(http.StatusOK, map[string]interface{}{
                "status": "success",
                "users":  users,
            })
            }
            . . .
      e. Create Router -> di dalam folder router ada file router.go
         Scriptnya yaitu :
         package routes
         import (
            "project/controllers"

            "github.com/labstack/echo"
         )
         func New() *echo.Echo {
            e := echo.New()
            e.GET("/users", controllers.GetUserControllers)
            return e
        }
       f. Create folder Main -> di dalam folder main ada file main.go
          scriptnya yaitu : 
          package main
          import (
            "project/config"
            "project/routes"
          )
          func main() {
            config.InitDB()
            e := routes.New()
            e.Logger.Fatal(e.Start(":8000"))
        }
    4. Jalankan file main.go => go run main.go
       Maka ketika run file main.go di terminal akan muncul tulisan echo di terminal, kemudian buka postman dan jakankan url nya.





        