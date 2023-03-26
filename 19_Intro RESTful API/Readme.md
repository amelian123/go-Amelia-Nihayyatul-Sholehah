## Section 19_INTRO RESTFUL API

## API => Application Programming Interface
>> API merupakan sekumpulan fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data dari sistem operasi, aplikasi, atau layanan lainnya.
>> API merupakan antarmuka yang berfungsi sebagai penghubung antara sebuah aplikasi dan aplikasi lainnya, atau antara klien dan server, untuk memungkinkan integrasi fitur tanpa harus menambahkan data secara manual.
>>Manfaat API antara lain :
  1. Mempermudah pembuatan aplikasi fungsional. contohnya Pada aplikasi transportasi ojek online yang memberikan banyak layanan, pasti akan sangat membutuhkan fungsionalitas maps. Nah, di sini, developer aplikasi transportasi/ojek online tidak perlu membuat aplikasi maps sendiri, cukup mengambil data dari Google Maps.
  2. Efisiensi pengembangan aplikasi 
  3. Meringankan beban server : semua data yang dibutuhkan di server tidak perlu disimpan secara keseluruhan karena sudah ada API. Cukup minta API untuk memperoleh data terbaru dari server aplikasi sumber. Kemudahan ini tentu akan sangat membantu karena server tidak akan terbebani dan meminimalkan down-time website.
>>Cara Kerja API yaitu :
  1. Aplikasi mengakses API. Pertama, API akan memulai pekerjaannya saat pengguna membuka aplikasi. Misalnya, pengguna membuka aplikasi pemesanan tiket online dan ingin mengakses tujuan tertentu. Di sini, aplikasi akan mengakses API maskapai penerbangan yang sudah dihubungkan.
  2. API membuat permintaan ke server. Setelah aplikasi berhasil mengakses alamat API, permintan akan diteruskan ke server maskapai. Jadi, API akan menyampaikan bahwa aplikasi membutuhkan data untuk tanggal dan tujuan penerbangan yang diminta. 
  3. Server merespons API. Setelah data ditemukan sesuai permintaan, server akan kembali ke API, lalu memberikan data berupa ketersediaan tempat duduk, waktu keberangkatan, dan lainnya.
  4. API memberikan hasil ke Aplikasi. Terakhir, informasi akan diberikan ke aplikasi yang diakses pengguna. Proses ini terjadi bersama permintaan ke maskapai yang lain; jadi, terkadang aplikasi bisa menampilkan jadwal dari berbagai maskapai sekaligus dalam satu kali permintaan.
>> Jenis-Jenis API antara lain : 
   1. Public API : boleh digunakan oleh siapa saja di berbagai platform. Public API juga sering disebut sebagai Open API, dan merupakan yang paling sering digunakan.
   2. Private API : tidak boleh digunakan secara umum. Jenis API ini biasanya digunakan untuk keperluan pribadi atau internal dalam pengembangan aplikasi tertentu. Sebagai contoh yaitu API dari backend yang digunakan mengakses frontend suatu website.
   3. Partner API : boleh digunakan secara umum, tapi hanya bagi pihak yang bekerja sama dan memilik izin untuk menggunakannya.

## REST (Representational State Transfer)
>> Format request dan respon yaitu JSON, XML, SOAP
>> Metode HTTP request yaitu GET, POST, PUT, DELETE, HEAD, OPTION, dan PATCH.
   ....GET : untuk mengambil data
   ....POST dan PUT : untuk menambahkan data. POST untuk insert data, sedangkan PUT untuk update data

## JSON (JavaScript Object Notation)
>> HTTP Response Code antara lain :
   a. status 200
      200 => berjalan dengan sesuai ekspektasi
      201 => server sudah membuat data baru dari inputan post
   b. status 400 (adanya eror dari sisi web browser)
      404 bad request => eror dari browser kita sendiri
      401 unauthorized => browser tidak bisa membuktikan kalau kita mempunyai hak akses yang sah untuk mengakses server tersebut
   c. status 500 (biasanya ada eror dari server)
      501 => fungsi server tidak mendukung permintaan dari browser
      500 => ada kesalahan dari sisi server atau internal server eror
      502 => ketika gateaway/proxy menerima respon yang tidak valid dari server

>> Tools API yaitu katalon, apache JMeter, postman, SoapUI

## Postman 
merupakan klien HTTP untuk menguji layanan web.Postman memudahkan untuk menguji, mengembangkan, dan mendokumentasikan API dengan memungkinkan pengguna dengan cepat mengumpulkan permintaan HTTP yang sederhana dan kompleks.

## Package GO, net/http, encoding/json
   A. Package net/http -> menyediakan implementasi klien http dan server
   B. Package encoding/json -> berisi banyak fungsi untuk operasi json yaitu :
   1. decode json ke struktur objek
   2. decode json ke map[string]interface{} dan interface{}
   3. decode array json ke array objek
   4. encode objek ke json string

## Library pada Golang antara lain :
   a. echo 
   b. Go Kit
   c. Logrus
   d. gorm.io
   e. cobra

>> Echo yaitu performa yang tinggi, ekstensible,dan minimalis web framework pada golang

>> Keunggulan menggunakan echo yaitu :
   1. Optimized Router -> dalam pemakaian router cukup powerful karena adanya optimasi penggunaan memori yang bernilai 0
   2. Data Rendering -> dalam memberikan respon pada sebuah request cukup mudah
   3. Data Binding
   4. Middleware
   5. Scalable -> cukup mudah bila dikembangkan dalam restAPI

>> Template echo bisa didapatkan dalam https://echo.labstack.com/guide/templates
>> Setting echo dengan perintah -> go get github.com/labstack/echo/v4

>>Retrieve Data antara lain :
  1. URL Params
  2. Query Params
  3. Form Value