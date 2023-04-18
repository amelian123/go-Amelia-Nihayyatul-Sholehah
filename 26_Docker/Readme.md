## Section 26_Docker

// Infrastructure and Deployment
   1. Docker
   2. System and Software Deployment
   3. CI/CD
   4. Kubernetes

// Mengapa docker dipelajari? 
   - kontainer bekerja 
   - untuk meluncurkan kontainer menggunakan Docker 
   - untuk membangun dan meluncurkan gambar kontainer milik sendiri
   - untuk menerapkan aplikasi owner sebagai kontainer

// Application Deployment History antara lain monolithic apps on physical, virtual machine abstraction, stateless and horizontal scalable apps, dan microservices and containers

## Container
Kontainer bukanlah mesin virtual.  Wadah adalah sebuah proses ... dengan isolasi sistem file. Semua yang ada di Linux adalah file.  
/dev/sda=hard disk 
/dev/proc=proses 
/dev/usb 
/dev/cpu 
/dev/std (in/out)
/bin/bash... hanya file biner

## Container vs Virtual Machines
>> Container
   1. Abstraksi pada lapisan aplikasi
   2. Kontainer memakan lebih sedikit ruang daripada VM (gambar kontainer biasanya berukuran puluhan MB)
   3. Menangani lebih banyak aplikasi dan memerlukan lebih sedikit VM dan sistem Operasi.
>> Virtual Machines
   1. Abstraksi perangkat keras fisik.  
   2. Setiap VM menyertakan salinan lengkap dari sistem operasi 
   3. Juga menjadi lambat untuk boot.

// Docker Basics 
   Berikut ini docker basics antara lain :
   1. Image
   2. Container
   3. Engine
   4. Registry
   5. Control Plane

// Yang dapat dilakukan dengan docker yaitu :
   1. Membangun docker images
   2. Upload dan publish images
   3. Download dan run

// Macam Syntax dalam Docker yaitu :
   1. From -> mendapatkan image dari docker registry
   2. Run -> eksekusi comand bash ketika container dibangun
   3. env -> set variabel di dalam container
   4. add -> copy file dengan proses lainnya
   5. workdir -> set pengerjaan file direktori 
   6. entrypoint -> eksekusi command ketika selesai membangun container
   7. cmd -> eksekusi comand tetapi, yang kelebihan penulisan

## PART 2
1. Definisi dan strategi deployment
2. Deployment cycle
3. Proses Deployment
4. GIT sebagai repository
5. AWS
6. Linux server (remote server, transfer file)
7. Setup dan instalasi webserver
8. setup domain HTTPS

// System and Sostware Deployment
Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara menjadi permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web dan api ke server sedangkan aplikasi mobile ke Playstore/Appstore.

// Strategi Deployment
   1. Big-Bang Deployment Strategy atau sering disebut Replace 
   >> Kelebihan Big-Bang Deployment Strategy yaitu mudah diimplementasikan, perubahan kepada sistem langsung 100% secara instan
   >> Kekurangan Big-Bang Deployment Strategy yaitu terlalu beresiko, rata-rata downtime cukup lama
   2. Deployment Strategy Rollout
   >> Kelebihan Rollout Deployment Strategy yaitu lebih aman dan less downtime dari versi sebelumnya
   >> Kekurangan Rollout Deployment Strategy yaitu tidak ada kontrol request, akan ada 2 versi berjalan secara bersamaan sampai semua server terdeploy
   3. Deployment Strategy Blue/Green
   >> Kelebihan Blue/Green Deployment Strategy yaitu perubahan sangat cepat, sekali switch service langsung berubah 100%, tidak ada isu beda versi pada service
   >> Kekurangan Blue/Grren Deployment Strategy yaitu resource yang dibutuhkan lebih banyak, testing harus diprioritaskan sebelum di switch
   4. Deployment Strategy Canary
   >> Kelebihan Canary Deployment Strategy yaitu cukup aman, mudah untuk rollback jika ada eror tanpa ada imbas ke semua user
   >> Kekurangan Canary Deployment Strategy yaitu untuk mencapai 100% cukup lama 

## Simple Dev-Ops Cycle 
   1. Plan
   2. Code
   3. Build
   4. Test
   5. Release
   6. Deploy
   7. Operate
   8. Monitor

## PART 3
   Continuous : integration, delivery, deployment

// Continuous Integration
   => adalah proses otomatis. Hal ini dilakukan untuk mengintegrasikan berbagai kode dari berbagai sumber potensial untuk membangun atau mengujinya.

// Continuous Delivery Deployment
   => proses penyebaran terus menerus berjalan satu langkah lebih jauh dari proses integrasi dan pengiriman. Jalur penerapan berkelanjutan secara otomatis menyebarkan setiap bangunan yang telah diverifikasi.

## PART 4_ KUBERNETES
Mengapa membutuhkan sistem orkestrasi kontainer? Kubernetes (K8s) adalah sistem sumber terbuka untuk mengotomatiskan penyebaran, penskalaan, dan pengelolaan aplikasi kemas. 

// Yang dapat dilakukan yaitu :
- Penemuan layanan dan penyeimbangan muatan 
- Penskalaan horizontal 
- Peluncuran dan pengembalian otomatis.  
- Manajemen rahasia dan konfigurasi

// Apa yang bukan termasuk kubernetes?  
- Tidak membatasi jenis aplikasi yang didukung 
- Tidak menyebarkan kode sumber dan tidak membangun aplikasi pengguna
- Tidak menyediakan layanan tingkat aplikasi, seperti middleware.  
- Tidak mendikte solusi logging, monitoring, atau alerting. 

// Kubernetes Workloads
- Deployment adalah pengawas untuk pod , memberi Anda kendali penuh atas bagaimana dan kapan versi pod baru diluncurkan serta dikembalikan ke keadaan sebelumnya.
- Service adalah abstraksi untuk pod, menyediakan alamat IP virtual (VIP) yang stabil.
- Ingress objek API yang mengelola akses eksternal ke layanan di cluster, biasanya HTTP.

## DNS (Domain Name Server)
DNS adalah sebuah sistem yang menghubungkan Uniform Resource Locator (URL) dengan internet protocol address (IP Address)

## MATERI DOCKER VOLUME
Docker volume dianggap sebagai storage/penyimpanan data di container.

// Running Container With Volume
   Ketika container terhapus, data yang ada pada container tetap tersimpan dan tidak akan ikut terhapus.

// Basic Command 
   a. Create Volume : docker create volume <name_volume>
   b. Remove Volume : docker volume rm <name_volume>
   c. Add Container To Volume : 
   docker container --name <name_container> \ 
   --mount "type=volume, source=<name_volume>, dst=<folder_destionation>"

// Docker Network
- Defaultnya container poda docker akan saling tensolasi satu sama lain Kita tidak dapat melakukan request api (misal) dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container di network yang sama.
- Secara default, container tidak bisa saling berkomunikasi satu sama lain. Oleh karena itu, perlu docker network agar bisa saling berkomunikasi antar docker.
