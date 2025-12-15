# Hafta 3 – Container Mantığı (Video 1)
Huawei CCE Bootcamp Ders Notları

Bu derste container kavramı, image yapısı, single-stage ve multi-stage build süreçleri, local registry simülasyonu, Huawei Cloud SWR kullanımı ve Kubernetes üzerinde image deploy işlemleri uygulanmıştır.

---

## 1. Container & Image Temelleri

### Docker Image Tanımı
- Uygulamanın bağımlılıkları ve çalıştırma ortamını içeren paketlenmiş yapıdır.
- Immutable katmanlardan oluşur.
- Container'lar bu imajdan üretilir.

### Container Image Tanımı
- Cloud ortamlarında kullanılan image kavramının karşılığıdır.
- Docker image yapısıyla aynıdır.

---

## 2. Dockerfile → Image → Container Süreci

Dockerfile  --->  Docker Image  --->  Docker Container

### Dockerfile Tanımı
Image’in nasıl oluşturulacağını tarif eden dosyadır (FROM, COPY, RUN, CMD).

### Katmanlı Mimari (Layered Architecture)
- Her komut bir katman oluşturur.
- Avantajları:
  - Daha hızlı build süreci
  - Güvenlik artışı
  - Cache mekanizması
  - Güncellemelerde düşük maliyet

---

## 3. Single-Stage Build

### Tanım
- Build ve runtime aynı imaj içinde bulunur.
- Build araçları final image içinde kalır → büyük boyut.

### Uygulama 1 – Single Stage (Go uygulaması)
Go uygulaması 8080 portunu dinler ve mesaj döndürür.

### Build
```bash
docker build -t uyg1:v1 .
docker image ls
```

### Çalıştırma
```bash
docker run -d --name uyg1con -p 80:8080 uyg1:v1
```

---

## 4. Multi-Stage Build

### Tanım
Build aşaması ve runtime aşaması farklı imajlarda tutulur.

### Avantajlar
- Daha küçük final image
- Build araçları runtime’a aktarılmaz
- Daha güvenli ve hızlıdır

### Uygulama 2 – Multi Stage (Go uygulaması)

### Build
```bash
docker build -t uyg2:v1 .
docker image ls
```

### Kod Değişikliği & Yeni Versiyon Build
```bash
docker build -t uyg2:v2 .
```

### Çalıştırma
```bash
docker run -d --name uyg2con -p 8080:8080 uyg2:v2
```

### Container Listeleme
```bash
docker ps
```

---

## 5. Docker Registry Mantığı

### Registry Tanımı
Container image depolama servisidir.

### Neden Kullanılır?
- CI/CD süreçleri
- Güvenli paylaşım ve dağıtım
- Versiyonlama ve yönetim

### Public & Private Registry
- DockerHub → Public registry
- SWR (Huawei Cloud) → Private registry

---

## 6. Local Docker Registry Simülasyonu

### Local Registry Oluşturma
```bash
docker run -d -p 5000:5000 --name local-registry registry:v2
```

### Image Tag Atma
```bash
docker tag uyg1:latest localhost:5000/my-single-stage-app:v1.0.0
```

### Push
```bash
docker push localhost:5000/my-single-stage-app:v1.0.0
```

### Pull
```bash
docker pull localhost:5000/my-single-stage-app:v1.0.0
```

### Registry İçeriğini Listeleme
```bash
curl -X GET http://localhost:5000/v2/_catalog
```

### Tag Listesi Görme
```bash
curl -X GET http://localhost:5000/v2/uyg1/my-single-stage-app/tag/list
```

---

## 7. Huawei Cloud – SWR (Software Repository)

### SWR Tanımı
Huawei Cloud üzerinde image saklama ve yönetme servisidir.

### Organizasyon Mantığı
- IAM kullanıcılarını izole etme
- Proje bazlı yetkilendirme
- Image yönetimini bölümlere ayırma

### Adımlar
1. Container engine indir  
2. Image build et  
3. Organization oluştur  
4. SWR login  
5. Tag → Push işlemi  

### Login
```bash
docker login -u <username> -p <token> <swr-endpoint>
```

### Tag & Push
```bash
docker tag uyg1:v1 <swr-endpoint>/<org>/uyg1:v1
docker push <swr-endpoint>/<org>/uyg1:v1
```

---

## 8. Kubernetes Üzerine Deployment

### Cluster Oluşturma
Huawei Cloud CCE üzerinde yeni cluster kuruldu.

### Workload (Deployment) Açma
- Kubernetes container değil pod yönetir.
- Pod: aynı amaç için çalışan container grubudur.

### Image Seçimi
SWR üzerinden push edilen image seçildi.

### Worker Node
Workload oluşturulduğunda worker node otomatik geldi.

### kubectl Yapılandırması
- Console → Configure üzerinden kubeconfig indirildi.
- Local makinaya eklendi.

### Bağlantı Testi
```bash
kubectl get nodes
```

### Nginx Deployment (Test)
```bash
kubectl create deployment nginx --image=nginx
kubectl expose deployment nginx --port=80 --type=LoadBalancer
```

---

## Özet
Bu videoda:
- Container & image temelleri  
- Single-stage / multi-stage build süreçleri  
- Local registry simülasyonu  
- Huawei SWR image push  
- Kubernetes cluster ve workload oluşturma  
uygulamalı olarak gerçekleştirilmiştir.

