## Section 22_Middleware

// Pengertian Middleware
Middleware adalah entitas yang menghubungkan ke pemrosesan permintaan / respon server. Middleware digunakan untuk menerapkan berbagai fungsi sekitar permintaan http masuk ke server user dan respon keluar.

//Contoh Bagian Middleware
1. Negroni
2. Echo
3. Interpose
4. Alice

## Setup Middleware Echo
   //a. Echo#Pre() => mengeksekusi sebelum router memproses request. Yang termasuk Echo#Pre() antara lain :
   1. HTTPSRedirect
   2. HTTPSWWWRedirect
   3. WWWRedirect
   4. AddTrailingSlash
   5. MethodOverride
   6. Rewrite
   //b. Echo#Use() => mengeksekusi setelah router memproses request dan memiliki akses penuh ke echo.Context API. Yang termasuk Echo#Use() antara lain :
   1. BodyLimit
   2. Logger
   3. Gzip
   4. Recover
   5. BasicAuth
   6. JWTAuth
   7. Secure
   8. CORS
   9. Static

## HTTPS Redirect
HTTPS redirect middleware redirects
Contoh dari https request to https yaitu http://labstack.com akan redirect ke https://labstack.com

## Log Middleware
//kegunaannya antara lain :
1. Logger middleware mencatat request HTTP.  
2. Sebagai jejak kaki.
3. Sumber data untuk analisa

## Cara Membuat Logger Middleware
Sebenarnya hampir sama dengan pembuatan MVC pada tugas sebelumnya yaitu config, controllers, lib, models, routes, dan kali ini ditambahkan folder middleware.
Cara membuat middleware yaitu :
1. Membuat folder middleware
2. Membuat file logMiddlewares.go di dalam folder middleware
3. Pada main.go yang di routes perlu ditambahkan kodingan github.com/nama_file/middleware
4. Kemudian run file. Maka akan muncul tulisan echo di tewrminal vscode


## Auth Middleware
// Kegunaan Autentikasi yaitu :
   1. sebagai identifikasi pengguna
   2. untuk keamanan data dan informasi

// Yang digunakan untuk autentikasi yaitu :
   1. basic autentikasi
   2. JSON Web Token

>> Basic Authentication
Basic Autentikasi adalah permintaan teknik otentikasi http, metode ini membutuhkan informasi nama pengguna dan kata sandi untuk dimasukkan ke dalam header permintaan.

Format Encode Informasi : 
'Otorisasi : Basic + base64encode (' username : password ') 
Header Generate : Otorisasi : Basic dXN1cm5hbWU6cGFzc3dvcmQ =

>> Cara Membuat Encode
   1. Pada folder middleware, buatlah 3 file yaitu authDBMiddleware.go , authMiddleware.go , logMiddleware.go

## JWT Middleware
JWT merupakan salah satu standar JSON (RFC 7519) untuk keperluan akses token. Token dibentuk dari kombinasi beberapa informasi yang di-encode dan di-enkripsi. Informasi yang dimaksud adalah header, payload, dan signature.

// Skema JWT 
   Berikut ini skema JWT antara lain :
   1. Header, isinya adalah jenis algoritma yang digunakan untuk generate signature.
   2. Payload, isinya adalah data penting untuk keperluan otentikasi, seperti grant, group, kapan login terjadi, dan atau lainnya. Data ini dalam konteks JWT biasa disebut dengan CLAIMS.
   3. Signature, isinya adalah hasil dari enkripsi data menggunakan algoritma kriptografi. Data yang dimaksud adalah gabungan dari (encoded) header, (encoded) payload, dan secret key.

>> Pada aplikasi yang menerapkan JWT, yang terjadi sedikit berbeda. Token adalah hasil dari proses kombinasi, encoding, dan enkripsi terhadap beberapa informasi. Nantinya pada sebuah http call, pengecekan token tidak dilakukan dengan membandingkan token yang ada di request vs token yang tersimpan di database, karena memang token pada JWT tidak disimpan di database. Pengecekan token dilakukan dengan cara mengecek hasil decode dan decrypt token yang ditautkan dalam request.

>> Handler ini bertugas untuk meng-otentikasi client/consumer. Data username dan password dikirimkan ke endpoint dalam bentuk HTTP Basic Auth. Data tersebut kemudian disisipkan dalam pemanggilan fungsi otentikasi authenticateUser()

>> Objek claims bisa dibuat dari tipe map dengan cara membungkusnya dalam fungsi jwt.MapClaims(); atau dengan meng-embed interface jwt.StandardClaims pada struct baru, dan cara inilah yang akan dipakai.

>> Claims isinya adalah data-data untuk keperluan otentikasi. Dalam prakteknya, claims merupakan sebuah objek yang memilik banyak property atau fields.Objek claims harus memiliki fields yang termasuk di dalam list JWT Standard Fields. Dengan meng-embed interface jwt.StandardClaims, maka fields pada struct dianggap sudah terwakili.

>> Pada aplikasi yang sedang dikembangkan, claims selain menampung standard fields, juga menampung beberapa informasi lain, oleh karena itu memerlukan struct baru yang meng-embed jwt.StandardClaims.
type MyClaims struct {
    jwt.StandardClaims
    Username string `json:"Username"`
    Email    string `json:"Email"`
    Group    string `json:"Group"`
}