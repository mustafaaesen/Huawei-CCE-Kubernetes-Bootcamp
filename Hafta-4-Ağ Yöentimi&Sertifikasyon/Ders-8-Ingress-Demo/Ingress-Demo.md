# Huawei Cloud CCE Bootcamp  
## Hafta 4 – Ders 8: Örnek Deployment (Ingress + LoadBalancer + DNS + SSL)

Bu derste, bootcamp boyunca öğrenilen tüm temel kavramlar **tek bir örnek deployment senaryosu** üzerinden pratik edilmiştir.  
Amaç: Kubernetes üzerinde çalışan birden fazla uygulamaya **domain + ingress + load balancer** üzerinden erişim sağlamak.

---

## 🎯 Amaç
- Kubernetes Deployment ve Service yapılarını pekiştirmek
- LoadBalancer ve NAT Gateway kavramlarını pratikte görmek
- Ingress ile **çoklu path yönlendirme** yapmak
- Domain üzerinden uygulamalara erişmek
- HTTPS / SSL kavramını mimari seviyede anlamak

---

## 📌 Kullanılan Başlıklar

- NAT Gateway (SNAT / DNAT)
- LoadBalancer (Layer 4 / Layer 7)
- HTTP vs HTTPS – SSL Sertifikası
- Huawei Cloud DNS Service
- Ingress & Ingress Controller
- Huawei Cloud CCE Autopilot Cluster

---

## 1️⃣ NAT Gateway ve LoadBalancer Kavramları

### 🔹 NAT Gateway
NAT Gateway, **VPC içindeki kaynakların internete çıkmasını** sağlar.

**Kullanım Amaçları:**
- Pod’ların Docker Hub, SWR gibi dış registry’lerden image çekebilmesi
- İnternete kapalı subnet’lerde çalışan kaynakların dış dünyaya erişmesi

**Türleri:**
- **SNAT:** İçeriden dışarıya çıkış (Pod → Internet)
- **DNAT:** Dışarıdan içeriye erişim (Internet → Resource)

Bu demoda:
- Pod’ların image çekebilmesi için **SNAT aktif edildi**

---

### 🔹 LoadBalancer
LoadBalancer:
- Gelen trafiği **birden fazla pod / node** arasında dağıtır
- Kubernetes Service veya Ingress için **dış dünyaya açılma noktasıdır**

**Faydaları:**
- Yük dağılımı (Load balancing)
- Yüksek erişilebilirlik
- Dış IP (Public IP) sağlar

---

## 2️⃣ HTTP – HTTPS – SSL Kavramı

### 🔹 HTTP
- Veriler **şifrelenmeden** iletilir
- Güvenli değildir
- Trafik dinlenebilir (Man-in-the-middle risk)

### 🔹 HTTPS + SSL
- Veri **şifrelenmiş** olarak iletilir
- SSL Sertifikası ile güvenli bağlantı kurulur
- Kimlik doğrulama + veri bütünlüğü sağlar

> Not: Bu derste HTTPS denemesi **ücretli SSL sertifikası** nedeniyle sadece izlenmiştir.

---

## 3️⃣ DNS Records (Domain Yönetimi)

DNS:
- Domain adını IP adresine çevirir

Bu senaryoda:
- Huawei Cloud DNS Service kullanılabilir
- Ya da farklı bir domain sağlayıcıdan:
  - **A Record** ile LoadBalancer Public IP’sine yönlendirme yapılır

Domain:
- Ingress demo için gereklidir
- Path bazlı yönlendirme domain üzerinden test edilir

---

## 4️⃣ Ingress Nedir?

Ingress:
- Kubernetes’te **çoklu URL / path yönlendirme kurallarını** tanımlar
- Tek bir LoadBalancer üzerinden birden fazla servise erişim sağlar

Örnek:
mustafaesen.com.tr/ → web-service
mustafaesen.com.tr/api → api-service
mustafaesen.com.tr/hello → hello-service


Ingress:
- Tek başına çalışmaz
- **Ingress Controller** gerektirir (örn: Nginx Ingress)

---

## 5️⃣ Huawei Cloud CCE Cluster Kurulumu

### 🔹 Cluster Ayarları
- **CCE Autopilot** seçildi
- Worker node’lar otomatik oluşturuldu
- Alarm Center kullanılmadı

### 🔹 Network Yapısı
- VPC oluşturuldu
- Subnet’ler ayrıldı:
  - Node subnet
  - Pod subnet
  - Control Plane subnet

### 🔹 SNAT
- Pod’ların internet erişimi için **aktif edildi**

### 🔹 Add-ons
- Monitoring & Log Collection:
  - Demo seviyesinde gerekmediği için **devre dışı bırakıldı**

---

## 6️⃣ Deployment ve Service Dosyaları

### 🔹 API Uygulaması
Dosyalar:
- `api-deployment.yaml`
- `api-service.yaml`

Durum:
- Hocanın kullandığı image SWR kaynaklı pull hatası verdi
- Alternatif olarak:
image: kennethreitz/httpbin

kullanıldı

Service:
- Port: **80**

---

### 🔹 Web Uygulaması
Dosyalar:
- `web-deployment.yaml`
- `web-service.yaml`

Özellikler:
- Docker Hub’dan image çekildi
- **2 replica**
- Port: **80**

---

### 🔹 Hello Uygulaması
Dosyalar:
- `hello-deployment.yaml`
- `hello-service.yaml`

Özellikler:
- Basit mesaj dönen uygulama
- Docker Hub image
- **1 replica**
- Port: **80**

---

## 7️⃣ Ingress Tanımı

Dosya:
- `ingress.yaml`

Tanımlanan Path’ler:
- `/`        → web-service
- `/api`     → api-service
- `/hello`   → hello-service

Not:
- Pratik sırasında ingress:
  - Yaml ile
  - Daha sonra **Huawei Cloud Console UI** üzerinden de tanımlandı

---

## 8️⃣ Cluster’a Bağlantı ve Deploy İşlemleri

Adımlar:
1. Cluster oluşturuldu
2. Elastic IP (EIP) alındı ve bağlandı
3. `kubectl config` dosyası indirildi
4. Local terminal üzerinden cluster’a bağlanıldı
5. Deployment ve Service yaml dosyaları apply edildi
6. Console üzerinden:
   - Pod
   - Deployment
   - Service durumları kontrol edildi

API deployment’ta pull hatası alınınca:
- Image değiştirilerek sorun çözüldü

---

## 9️⃣ Ingress Controller Kurulumu

Ingress için:
- **Ingress Controller Add-on** gereklidir
- Bu add-on için **LoadBalancer zorunludur**

### 🔹 LoadBalancer Oluşturma
- Network + Application (Layer 4 + Layer 7)
- Traffic Mode:
  - Fixed / Elastic
- VPC:
  - Cluster kurulurken oluşturulan VPC
- Subnet:
  - Node subnet (control plane subnet)

---

### 🔹 Nginx Ingress Controller
- Add-ons bölümünden kuruldu

---

## 🔟 Ingress UI Üzerinden Tanımlama

Console üzerinden:
- Ingress oluşturuldu
- İsim: `multi-path-ingress`
- Domain girildi
- Path bazlı yönlendirmeler tanımlandı
- Service ve port eşleştirildi

Sonuç:
- Tarayıcıdan test edildi
- `/`, `/api`, `/hello` path’leri doğru servislere yönlendi

---

## 1️⃣1️⃣ HTTPS Denemesi

- Mevcut ingress ve add-on’lar silindi
- HTTPS olarak yeniden tanımlama denendi
- SSL sertifikası **ücretli** olduğu için:
  - Kurulum izlenerek geçildi
  - Mimari mantık anlatıldı

---

## ✅ Sonuç

Bu ders ile:
- Kubernetes temel objeleri tek senaryoda birleştirildi
- Ingress + LoadBalancer + DNS ilişkisi netleşti
- Huawei Cloud CCE üzerinde gerçekçi bir prod mimarisi görüldü
- HTTP / HTTPS farkı uygulamalı olarak incelendi

---

📌 **Not:**  
Bu yapı, gerçek hayatta:
- Microservice mimarileri
- API Gateway senaryoları
- Multi-domain / multi-path projeler için temel oluşturur.

