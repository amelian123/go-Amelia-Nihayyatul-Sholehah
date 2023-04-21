## Section 28_CI/CD
CI/CD (Continuous Integration/Continuous Deployment) adalah praktik pengembangan perangkat lunak yang bertujuan untuk meningkatkan kualitas, kecepatan, dan keandalan dalam pengiriman aplikasi. CI/CD melibatkan penggunaan otomatisasi dan integrasi terus-menerus dalam pengembangan perangkat lunak, dari pengujian hingga pengiriman.
// CI (Continuous Integration) adalah praktik pengembangan perangkat lunak di mana setiap perubahan kode yang dibuat oleh pengembang diuji secara otomatis dan diintegrasikan dengan kode yang sudah ada sebelumnya secara berkala dan sering. Dalam CI, kode yang dihasilkan akan terus diuji untuk memastikan tidak terjadi error dan kode yang sudah diintegrasikan masih dapat berjalan dengan baik.
// CD (Continuous Deployment) adalah praktik pengembangan perangkat lunak di mana setiap perubahan kode yang sudah melewati tahap CI akan diunggah secara otomatis ke lingkungan produksi. Dalam CD, perubahan kode yang sudah diuji akan segera diterapkan ke aplikasi yang sudah berjalan tanpa harus menunggu waktu lama.

// Adapun manfaat dari CI/CD adalah:
   1. Meningkatkan kualitas kode: Dalam CI/CD, setiap perubahan kode akan diuji secara otomatis, sehingga mengurangi peluang adanya error dan memastikan bahwa kode yang dihasilkan berkualitas.
   2. Meningkatkan kecepatan pengiriman aplikasi: Dalam CI/CD, setiap perubahan kode yang sudah diuji akan segera diterapkan ke aplikasi yang sudah berjalan, sehingga mengurangi waktu yang diperlukan untuk pengiriman aplikasi.
   3. Mengurangi biaya dan risiko pengiriman aplikasi: Dalam CI/CD, setiap perubahan kode akan diuji secara otomatis, sehingga mengurangi peluang adanya error yang dapat mempengaruhi biaya dan risiko dalam pengiriman aplikasi.
  4. Memudahkan manajemen versi: Dalam CI/CD, setiap perubahan kode yang sudah diuji akan diberikan nomor versi secara otomatis, sehingga memudahkan manajemen versi aplikasi.
  5. Meningkatkan kolaborasi antar tim: Dalam CI/CD, setiap perubahan kode yang sudah diuji akan diterapkan secara otomatis, sehingga memudahkan tim dalam berkolaborasi dan mengurangi kesalahan yang dapat terjadi dalam pengiriman aplikasi.

## Dalam keseluruhan, CI/CD adalah praktik pengembangan perangkat lunak yang sangat penting dalam meningkatkan kualitas, kecepatan, dan keandalan dalam pengiriman aplikasi. CI/CD dapat membantu pengembang dalam mengurangi biaya dan risiko pengiriman aplikasi, meningkatkan kolaborasi antar tim, serta memudahkan manajemen versi aplikasi.

// Infrastructure and Deployment
   1. Docker
   2. System and Software Deployment
   3. CI/CD
   4. Kubernetes

// Application Deployment History antara lain monolithic apps on physical, virtual machine abstraction, stateless and horizontal scalable apps, dan microservices and containers

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

## Running CI/CD on Kubernetes
   1. Aktifkan repo FE pada Travis
   2. Tambahkan .travis.yml ke dalam folder FE
   3. Tambahkan deploy.sh ke dalam folder FE
   4. Konfigurasi beberapa variabel seperti :
      a. DOCKER_USERNAME = username docker hub
      b. DOCKER_PASSWORD = password login docker hub
      c. GCLOUD_SERVICE_KEY_PRD = next slide
      d. CLUSTER_NAME_PRD = alterra-5-gke
      e. CLOUDSDK_COMPUTE_REGION = asia-southeast1-b
      f. PROJECT_NAME_PRD = alterra-academy
