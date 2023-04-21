## Section 27_COMPUTE SERVICES
Compute Services mengacu pada jenis layanan komputasi awan yang menyediakan akses kepada pelanggan untuk menggunakan sumber daya komputasi, seperti mesin virtual (virtual machine/VM) atau container, dengan pembayaran berbasis penggunaan. Layanan compute merupakan bagian fundamental dari infrastruktur komputasi awan dan digunakan untuk menjalankan aplikasi, melakukan pemrosesan data, dan menyimpan data.

// Infrastructure and Deployment
   1. Docker
   2. System and Software Deployment
   3. CI/CD
   4. Kubernetes

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

## Ada beberapa jenis layanan compute, di antaranya:
   1. Mesin Virtual (VM): Mesin virtual adalah emulasi perangkat lunak dari komputer fisik. VM biasanya digunakan untuk menjalankan aplikasi legacy atau aplikasi yang membutuhkan sistem operasi atau konfigurasi tertentu.
   2. Container Instance: Container instance adalah lingkungan yang ringan dan portabel yang dapat digunakan untuk menjalankan aplikasi secara terisolasi. Container menjadi alternatif populer untuk mesin virtual karena lebih cepat untuk memulai dan membutuhkan lebih sedikit sumber daya sistem.
   3. Fungsi Serverless: Komputasi serverless adalah model di mana penyedia awan mengelola infrastruktur yang diperlukan untuk menjalankan aplikasi, memungkinkan pengembang untuk fokus pada menulis kode. Fungsi serverless didorong oleh peristiwa, artinya hanya dieksekusi ketika dipicu oleh suatu peristiwa.

// Compute services biasanya disediakan oleh penyedia cloud seperti Amazon Web Services (AWS), Microsoft Azure, dan Google Cloud Platform (GCP). Para penyedia ini menawarkan berbagai jenis layanan compute yang dapat disesuaikan dengan kebutuhan pelanggan individu, mulai dari startup kecil hingga perusahaan besar. Pelanggan dapat memilih untuk membayar layanan compute berdasarkan per jam, per menit, atau per detik, sehingga mudah untuk meningkatkan atau menurunkan sumber daya komputasi sesuai dengan fluktuasi permintaan.

// Compute Services juga menawarkan beberapa keuntungan dibandingkan komputasi on-premises tradisional. Misalnya, layanan ini dapat membantu mengurangi biaya dengan menghilangkan kebutuhan untuk berinvestasi dalam perangkat keras dan infrastruktur yang mahal. Layanan compute juga sangat skalabel, memungkinkan pelanggan dengan cepat dan mudah meningkatkan atau menurunkan sumber daya komputasi sesuai kebutuhan. Secara keseluruhan, layanan compute merupakan komponen penting dari infrastruktur komputasi awan, memberikan akses kepada pelanggan untuk sumber daya komputasi dengan pembayaran berbasis penggunaan. Layanan ini menawarkan beberapa keuntungan dibandingkan dengan komputasi on-premises tradisional, termasuk penghematan biaya, skalabilitas, dan fleksibilitas.

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

## RDS (Relational Database Service)
RDS (Relational Database Service) dan EC2 (Elastic Compute Cloud) adalah dua layanan yang disediakan oleh Amazon Web Services (AWS) untuk mempermudah pengguna dalam mengelola infrastruktur aplikasi dan database. RDS adalah layanan manajemen database yang disediakan oleh AWS. Layanan ini memungkinkan pengguna untuk dengan mudah membuat, mengoperasikan, dan mengelola database relasional dalam cloud. Dalam layanan ini, AWS menyediakan manajemen konfigurasi database, pemeliharaan, dan tugas-tugas administratif lainnya, sehingga pengguna dapat fokus pada aplikasi mereka. RDS mendukung beberapa jenis database, seperti MySQL, PostgreSQL, Oracle, Microsoft SQL Server, dan Amazon Aurora. AWS juga menyediakan berbagai opsi skalabilitas dan keamanan database yang memungkinkan pengguna untuk mengatur kapasitas dan konfigurasi database mereka sesuai kebutuhan.

## EC 2 (Elastic Compute Cloud)
EC2 adalah layanan komputasi awan yang disediakan oleh AWS yang memungkinkan pengguna untuk menyewa mesin virtual (virtual machine/VM) dalam cloud. Layanan ini memungkinkan pengguna untuk mengakses dan mengelola sumber daya komputasi sesuai dengan kebutuhan aplikasi mereka.EC2 mendukung berbagai sistem operasi, termasuk Linux, Windows, Unix, dan FreeBSD. AWS juga menyediakan berbagai jenis instance EC2 yang dapat disesuaikan dengan kapasitas, jenis pemrosesan, dan jenis penyimpanan.Selain itu, EC2 juga memungkinkan pengguna untuk menggunakan beberapa fitur tambahan seperti Auto Scaling, Load Balancing, dan Amazon Elastic Block Store (EBS). Fitur-fitur ini membantu pengguna dalam meningkatkan kapasitas dan kinerja aplikasi mereka dengan mudah. Dalam keseluruhan, RDS dan EC2 adalah layanan yang sangat berguna dalam mengelola aplikasi dan database dalam lingkungan cloud. Kedua layanan ini membantu pengguna dalam mengurangi biaya infrastruktur, meningkatkan fleksibilitas dan skalabilitas, serta memungkinkan pengguna untuk fokus pada aplikasi mereka tanpa harus khawatir tentang manajemen infrastruktur.

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
