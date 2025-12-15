# 📘 Hafta 4 – Ders 7  
## **Concept & Ingress Controller**

Bu derste Kubernetes ortamında autoscaling, OSI katmanları, DNS-domain ilişkisi, Ingress Controller ve Ingress kaynak tipleri detaylı şekilde incelendi. Ayrıca pratikte node/pod bazlı autoscaling gözlemlendi ve ingress yönlendirme mantığı ele alındı.

---

# 🟦 1. AUTOSCALING

Autoscaling, yük arttığında **otomatik genişleme (scale-out)**, yük azaldığında **otomatik daralma (scale-in)** sağlayan bir mekanizmadır.  
Manuel müdahaleye gerek kalmadan belirlenen kurallara göre pod veya node ekleyip çıkarır.

### **Temel Mantık**
- Yük genellikle **CPU** üzerinden ölçülür (en yaygın metrik).
- **Scale-out:** Pod sayısını artırır  
- **Scale-in:** Pod sayısını azaltır  
- Kubernetes bunu **HPA (Horizontal Pod Autoscaler)** veya **Node Autoscaler** ile gerçekleştirir.

---

## 🟦 Huawei Console Autoscaling

Cluster üzerinden:

### **Pod Bazlı Autoscaling**
- Menü: `Cluster → Autoscaling → Pod-based Autoscaling (HPA)`

### **Node Bazlı Autoscaling**
- Menü: `CCE Advanced HDR → Node-based Autoscaling`
- Node pool oluşturuldu
- Node eklendi
- `More → Enable Scale-in Protection` özelliği tanıtıldı  
  Bu özellik, belirli node’ların scale-in sırasında silinmesini engellemek için kullanılır.

### **Autoscaling Rule Oluşturma**
Node pool → Autoscaling → **Add Rule**

- **Trigger:** CPU % threshold (ör. %70 üzeri → scale-out)  
- **Action:** Eklenecek node sayısı  
    - İstersen manuel, istersen otomatik hesaplatılabilir  
- Özellikle **gün içinde belirli saatlerde trafiği artan uygulamalar** için ideal bir senaryo.

### **Pratik Gözlem**
- 20 pod içeren bir Deployment oluşturuldu.
- İlk durumda yeterli node olmadığı için pod’lar **NotReady** oldu.
- Node autoscaler devreye girince yeni node’lar yaratıldı.
- Node’lar geldikten sonra pod’lar sırayla **Running** durumuna geçti.
- Böylece autoscaling davranışı canlı olarak gözlemlendi.

---

# 🟦 2. OSI LAYERS (Kubernetes ile ilişkisi)

Servislerin dış dünyaya açılmasında NodePort ve LoadBalancer yöntemleri kullanılır.  
Bu mekanizmaları doğru anlayabilmek için OSI katmanları önemlidir.

---

## **OSI Modeli ve Kubernetes Bağlantısı**

| Katman | Açıklama | Kubernetes İlgisi |
|-------|----------|-------------------|
| **Layer 7 – Application** | Trafiğin türünü ve içeriğini anlar. | **Ingress Controller** burada çalışır. URL, host, path bazlı yönlendirme yapılır. |
| **Layer 6 – Presentation** | Veri formatı, şifreleme. | - |
| **Layer 5 – Session** | İletişim oturumları. | - |
| **Layer 4 – Transport** | TCP/UDP, portlar. | **NodePort / LoadBalancer** burada çalışır. |
| **Layer 3 – Network** | IP yönlendirmesi. | Servisler IP’leri bilir ama yol detayını bilmez. |
| **Layer 2 – Data Link** | Paket iletimi. | - |
| **Layer 1 – Physical** | Donanım. | - |

### Kubernetes’in doğrudan ilişkili olduğu katmanlar:
- **Layer 3 → IP düzeyi**  
- **Layer 4 → Port düzeyi**  
- **Layer 7 → Uygulama düzeyi (Ingress Controller)**

### Neden Ingress Controller Layer 7’de?
- Tek domain → birden çok uygulama olabilir  
- Örneğin e-ticaret site yapısı:
  - /seller → Satıcı paneli  
  - /customer → Müşteri paneli  
  - /admin → Yönetim paneli  
- Bu yolları anlamak **layer 7 load balancer** gerektirir → Ingress Controller devreye girer.

---

# 🟦 3. DOMAIN & DNS

### **Domain → İnsanlar için adres**
### **IP → Makineler için adres**

DNS çözümlemesi:
- Kullanıcı `example.com` yazar  
- DNS bu domainin IP adresini bulur  
- Trafik Kubernetes cluster’a yönlenir

Cluster’a gelen trafik:
1. LoadBalancer → cluster’a giriş  
2. Ingress Controller → URL/host/path kontrolü  
3. Belirlenen servise yönlendirme

**DNS sadece IP çevirir, yönlendirmeyi Ingress Controller yapar.**

---

# 🟦 4. INGRESS CONTROLLER

Kubernetes’te Service (NodePort / LoadBalancer) kullanarak dış dünyaya açılmak mümkündür.  
Ancak:

- Birden fazla uygulama varsa  
- Her biri için ayrı LoadBalancer almak  
- Ayrı servisler oluşturmak  

**maliyet ve yönetim açısından zahmetlidir.**

### Ingress Controller ne sağlar?

- Layer 7 seviyesinde çalışır  
- Host ve path’i okuyup anlar  
- Tek bir LoadBalancer ile çok sayıda uygulamayı dış dünyaya açabilir
- Ekonomiktir, yönetimi kolaydır

### E-Ticaret örneği:
```
/customer -> customer-service
/seller   -> seller-service
/admin    -> admin-service
```

Ingress Controller URL'yi anlar ve doğru servise yollar.

---

# 🟦 5. INGRESS

Ingress, yönlendirme kurallarını içeren **YAML tabanlı bir kaynak tanımıdır**.

Ingress Controller → Yürütücü  
Ingress → Kuralların yazıldığı tanım dosyası

---

# 🟩 **6. INGRESS TÜRLERİ (Kısa Açıklamalar)**

### **1) Basic Ingress**  
Tek domain → Tek service yönlendirmesi.

### **2) Multi-Path Ingress**  
Aynı domain altında farklı path’ler → farklı servislere yönlendirilir.

### **3) Multi-Host Ingress**  
Bir Ingress içinde birden fazla domain/subdomain yönetme.

### **4) TLS Ingress**  
HTTPS terminasyonu için TLS sertifikası kullanılır.

### **5) Rewrite / Path Manipulation Ingress**  
Gelen URL path'ini backend’e farklı şekilde iletir.

### **6) Canary Ingress**  
Trafiğin belirli bir yüzdesini yeni versiyon servise yönlendirir (ör. %20).

### **7) Security / Rate Limit Ingress**  
IP kısıtlamaları, istek limitleri, timeout ve body limitleri.

---

# 🟦 7. Ders İçinde Denenen Uygulamalar

- LoadBalancer oluşturuldu  
- Service ayarları yapıldı  
- Ingress Controller etkinleştirildi  
- Ingress oluşturuldu (UI üzerinden)  
- HTTP vs HTTPS farkı anlatıldı  
- Node seçimi ve LB bağlantısı gösterildi  

---

# 🟩 **8. Ek Bilgi: Ders boyunca oluşturulan tüm Ingress türlerinin YAML'ları ayrı ayrı hazırlanmış ve doğrulanmıştır.**  
Bu set aşağıdaki türleri içeriyor:

- Basic Ingress  
- Multi-Path Ingress  
- Multi-Host Ingress  
- TLS Ingress  
- Rewrite / Path Manipulation  
- Canary Ingress (Stable + Canary)  
- Security / Rate Limit Ingress  

➡️ Tümü başarıyla test edilmiş ve dökümantasyona eklenmiştir.

---

# ✔️ Ders 7 – ÖZET

Bu ders, Kubernetes ortamında trafiği yönetmenin tüm yöntemlerini anlamak için temel teori + pratik içeren kritik bir bölümdü.  
Ingress Controller’ın **neden** var olduğu, **hangi problemi çözdüğü** ve **nasıl yapılandırıldığı** hem teorik hem pratik olarak gösterildi.  


- Autoscaling mantığı  
- OSI katmanlarının Kubernetes ile ilişkisi  
- DNS’ten Ingress Controller’a trafik akışı  
- Tek LB ile çoklu uygulama yönetme  
- Tüm ingress türleri  

Kavrandı


