# Docker Pratik Notları  
Huawei CCE Kubernetes Eğitimi – Docker Bölümü  
Bu doküman eğitim sürecinde terminalde uyguladığım Docker komutlarının sade bir özeti ve açıklamasıdır.

---

## 1. Sistem Güncelleme ve Docker Kurulumu

### 1.1 Sistem güncelleme
```
sudo apt update && sudo apt upgrade -y
```
Sistemdeki paketleri günceller, güvenlik yamalarını yükler.

### 1.2 Docker kurulum
```
sudo apt install docker.io -y
```
Ubuntu’nun resmi reposundan Docker Engine paketini kurar.

### 1.3 Docker servisini başlatma
```
sudo systemctl start docker
sudo systemctl enable docker
sudo systemctl status docker
```
Docker servisinin çalışmasını sağlar ve her açılışta otomatik başlaması için aktif eder.

### 1.4 Docker versiyon kontrol
```
docker --version
```
Kurulu Docker sürümünü gösterir.

---

## 2. Docker Image Komutları

### 2.1 Kurulu image’leri listeleme
```
docker images
```
Localde bulunan tüm Docker image’lerini listeler.

### 2.2 Nginx container çalıştırıldığında image otomatik iner
```
docker run -d -p 8080:80 nginx
```
Nginx image yoksa Docker Hub’dan indirir ve container’ı arka planda çalıştırır.

### 2.3 Exam için önemli: curl ile public IP öğrenme
```
curl icanhazip.com
```

---

## 3. Docker Container Komutları

### 3.1 Container listeleme
```
docker ps -a
```
Çalışan ve durmuş tüm container’ları listeler.

### 3.2 Container durdurma
```
docker stop <container_id>
```
Container’ı düzgün bir şekilde durdurur. (SIGTERM)

### 3.3 Container silme
```
docker rm <container_id>
```
Durdurulmuş container’ı sistemden siler.

---

## 4. Çalışan Container ile Etkileşim

### 4.1 Nginx container başlatma
```
docker run -d -p 8080:80 nginx
```
8080 portunu container içindeki 80 portuna yönlendirerek yeni bir nginx container’ı çalıştırır.

### 4.2 Container içine girme (exec)
```
docker exec -it <container_id> /bin/bash
```
Container içinde terminal açar.  
Çıkmak için: `exit`

### 4.3 Container loglarını görüntüleme
```
docker logs <container_id>
```
Başlangıç ve istek loglarını gösterir.

### 4.4 Canlı kaynak kullanımı izleme
```
docker stats
```
Container CPU, RAM, I/O, network gibi bilgileri canlı olarak gösterir.  
Çıkmak için: **CTRL + C**

---

## 5. Sık Yapılan Hatalar ve Çözümler (Notlar)

### 5.1 exec hatası
Hata:
```
docker exec -it <id>/bin/bash
```
Çözüm: ID ve komut arasında boşluk olmalı.
```
docker exec -it <id> /bin/bash
```

### 5.2 Nginx'e tarayıcıdan bağlanamama
Sebep: Huawei Cloud Security Group'ta port 8080 açılamamış olabilir.

Çözüm:
- Security Group → Inbound Rule →  
  - Protocol: TCP  
  - Port: 8080  
  - Source: 0.0.0.0/0  
  - Priority: 1  

### 5.3 favicon.ico 404 hatası
Nginx varsayılan kurulumda favicon bulunmadığı için bu hata normaldir.

---

## 6. Kullanılan Tüm Komutların Sade Listesi
(Eğitim boyunca terminalde kullanılan komutlar)

```
sudo apt update && sudo apt upgrade -y
sudo apt install docker.io -y
sudo systemctl start docker
sudo systemctl enable docker
sudo systemctl status docker
docker --version
docker run -d -p 8080:80 nginx
curl icanhazip.com
docker images
docker ps -a
docker stop <id>
docker rm <id>
docker stats
docker exec -it <id> /bin/bash
docker logs <id>
history
```

---

## 7. Docker Yaşam Döngüsü (Özet)
1. Image indir → `docker pull`
2. Container oluştur → `docker run`
3. Container kontrol → `docker ps`
4. Container içine gir → `docker exec`
5. Log incele → `docker logs`
6. Durdur → `docker stop`
7. Sil → `docker rm`

Bu temel akış Kubernetes’e geçişte yeterlidir.

---


