# Huawei Cloud CCE – AOM (Application Operation Management) & Monitoring  
## Bootcamp Week – Ders 6 Notları

Bu ders kapsamında CCE üzerinde uygulama operasyon yönetimi (AOM), monitoring eklentileri, node hazırlıkları, add-on kurulumu ve cluster gözlemleri incelenmiştir.  
Aşağıdaki notlar konsolda uygulamalı olarak yapılan işlemlerin düzenlenmiş hâlidir.

---

## 📌 1. Cluster Hazırlığı ve Node Oluşturma

Mevcut cluster üzerinde uygulama dağıtımı için hazırlık adımları tamamlandı.  
Yeni node ekleme aşamasında yapılan seçimler:

- **Container Engine:** containerd  
- **Operating System:** Ubuntu  
- **Disk:** Varsayılan  
- **Subnet:** default  
- **Node IP:** Automatic  
- **EIP:** Auto Assign  

### Node Configuration

Node üzerinde Kubernetes yönetişimi için şu yapılandırmalar tanımlandı:

- **Node Labels:**  
  Pod’ların belirli node’larda çalışması için affinity kurallarıyla birlikte kullanılır.

- **Taints:**  
  Toleransı olmayan pod’ların bu node’a schedule edilmesini engellemek için tanımlanır.  
  > Add-on bileşenlerinin çalışabilmesi için bazı node’ların taint’li olması gerekebilir.

- **Labels ile Taints Farkı:**  
  - *Taint:* Pod’u uzak tutar.  
  - *Label:* Pod’un gitmesini istediğimiz node’u işaretler.

Submit sonrası bazı durumlarda:
- Node çalışır görünse bile, pod planlanamayabilir.
- Bunun nedeni gerekli node sayısının yetersizliği veya scheduling kuralları olabilir.
- Gerekli node eklendiğinde pod çalışır hâle gelir.

---

## 📌 2. Add-On Yapısı ve Kurulumlar

Add-on’lar, CCE cluster’ına ek işlevler kazandıran Huawei Cloud servis bileşenleridir.

Aşağıdaki add-on’lar kuruldu ve incelendi:

---

### 🔹 **Cluster Autoscaler**

- Cluster’ın trafikteki değişime göre otomatik ölçeklenmesini sağlar.  
- Gerektiğinde yeni node açabilir veya node sayısını azaltabilir.

---

### 🔹 **Cluster Monitoring**

Cluster’ın CPU, RAM, disk ve metrik durumlarını takip eden add-on’dur.  
**Log sistemi ile karıştırılmamalıdır.**

- 2 Deployment  
- 1 StatefulSet  
- 1 DaemonSet  
olmak üzere toplam **4 bileşen** hâlinde çalışır.

Bu add-on kritik olduğundan, pod’ları kaybetmemek için genellikle **çoklu node’a dağılmış** şekilde çalıştırılır.

---

### 🔹 **Grafana**

- Monitoring add-on’un topladığı metrikleri dashboard üzerinde gösterir.
- Harici erişim için **EIP atanır**.
- Bu EIP üzerinden Grafana arayüzüne ulaşılır.

---

### 🔹 **Node Problem Detector**

Cluster içindeki node’larda oluşan hata ve sorunları algılar, AOM üzerinden görünür hâle getirir.

---

### 🔹 **Cloud Native Log Collection**

- Uygulamaların ürettiği logları toplar, filtreler ve ilgili servislere iletir.
- Kullanıcı hareketleri, hata logları vb. gözlemlenebilir.

---

### 🔹 **Kubernetes Metrics Server**

- Pod ve node kaynak metriklerini Kubernetes API’ye iletir.
- `kubectl top` komutlarının çalışmasını sağlar.

---

## 📌 3. Monitoring & Log İncelemeleri

### Monitoring Center / IAM User Center / Alarm Center

Ders kapsamında bu üç panel ayrı sekmelerde açılarak incelendi:
- **Monitoring Center:** Cluster ve kaynak düzeyindeki metrikler  
- **IAM User Center:** Kullanıcı ve erişim hareketleri  
- **Alarm Center:** Sistem uyarıları, eşik aşımı bildirimleri, SMS/e-mail alarm tanımları  

---

## 📌 4. Cluster Resource Yönetimi

Cluster Resource ekranı cluster’ın genel durumunu gösterir.

### CPU Limitleri
- Pod seviyesinde CPU limiti belirlenebilir.
- CPU limitleri aşılırsa pod yavaşlar ama kill edilmez.

### Memory Limitleri
- Memory limiti aşılırsa Kubernetes pod’u **kill eder ve yeniden başlatır**.  
- Bunun nedeni node’un *Out of Memory* olup tamamen düşmesini engellemektir.

---

## 📌 5. Log İnceleme

- Pod logları Cloud Native Log Collection üzerinden görülebilir.
- `curl` veya filtreleme seçenekleriyle loglar anlık olarak incelendi.

Grafana arayüzü açılarak dashboard metrikleri görüntülendi.

---

## 📌 6. Özet

Bu derste:

- Node ekleme ve Kubernetes label/taint mantığı
- Add-on kurulumu ve işlevleri  
- Monitoring, log ve alarm altyapısı  
- Pod scheduling davranışları  
- Cluster resource yönetimi (CPU & Memory limitlerinin etkisi)  
- Grafana ve metrics server ile metrik takibi  
- Log toplama ve hata izleme mekanizmaları  

gibi konular tamamlanmıştır.

---


