#  Hafta 1 – Ders 2  
## Kubernetes Temelleri  
### Pod, Deployment, Service, Namespace ve Genel Yapı

Bu derste Kubernetes’te pod, workload, service ve YAML yapılarının temelleri; panel üzerinden iş yükü oluşturma; pod yaşam döngüsü; workload türleri ve service tipleri teorik + uygulamalı olarak ele alınmıştır.

---

#  Pod, YAML ve Temel Kavramlar

##  Pod Nedir?
Pod, Kubernetes’in **en küçük deploy edilebilir birimidir**.  
Bir veya birden fazla container’ı aynı ağ & depolama alanını paylaşacak şekilde bir arada tutar.

- Her pod’un kendi **IP adresi** vardır.
- Pod’lar geçicidir; silinebilir, yeniden oluşturulabilir.

---

##  YAML Nedir?
YAML, Kubernetes nesnelerini tanımlamak için kullanılan **insan tarafından okunabilir** bir konfigurasyon dosyası formatıdır.  
Tüm Kubernetes kaynakları YAML üzerinden tarif edilir.

---

## Pod Yaşam Döngüsü (Lifecycle)
Bir pod oluşturulduğunda aşağıdaki aşamalardan geçer:

1. **Pending** – Pod planlandı ancak henüz çalışmıyor.  
2. **ContainerCreating** – Container image’ı çekiliyor ve pod hazırlanıyor.  
3. **Running** – Pod aktif olarak çalışıyor.  
4. **Succeeded / Failed** – Pod başarılı veya hatalı şekilde tamamlandı.  
5. **CrashLoopBackOff** – Uygulama sürekli çöküyor ve yeniden başlatılıyor.

Lifecycle ayarlarıyla:
- Başlatma komutları (postStart)  
- Durdurma adımları (preStop)  
- Restart politikaları  

belirlenebilir.

---

##  CI/CD Pipeline
CI/CD, uygulamaların **otomatik build, test ve deploy** sürecini tanımlar.  
Kubernetes tarafında çoğunlukla YAML dosyalarının otomatik olarak cluster’a uygulanmasıyla sonuçlanır.

---

#  Workload Türleri

##  Deployment
Stateless uygulamalar için kullanılır.  
Pod ölürse **ReplicaSet** otomatik olarak yeniden ayağa kaldırır.

##  StatefulSet
Veri tutarlılığı gereken durumlarda (ör. veritabanları).  
Pod isimleri ve storage **kalıcıdır**: `pod-0`, `pod-1` gibi.

##  Job
Tek seferlik bir görevi tamamlayana kadar çalışır.  
Görev bitince pod sonlanır.

##  CronJob
Job’un **zamanlanmış hali**dir.  
Belirlenen cron pattern’e göre çalışır:  
`0 */1 * * *` → Her saat başı.

---

# Service Türleri

Service, pod’lara **kalıcı erişim noktası** sağlar.  
Pod IP’leri değişse bile service IP hep sabit kalır.

##  ClusterIP
Varsayılan servistir.  
Sadece Kubernetes cluster içinden erişilir.

##  NodePort
Her node’un 30000–32767 arası portlarından dışarıya açılır.  
`<NodeIP>:NodePort` şeklinde erişilir.

##  LoadBalancer
Cloud ortamlarında dış dünyaya **public IP** ile erişim sağlar.

---

#  YAML Dosyasının Yapısı

Kubernetes YAML dosyaları üç ana bölüm içerir:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-app
spec:
  replicas: 3
  template:
    spec:
      containers:
        - name: example
          image: nginx
```

###  metadata
Kaynağın adı, etiketleri, açıklamaları.

###  kind
Nesnenin türü:  
`Pod`, `Deployment`, `Service`, `Job`, `CronJob` vb.

###  spec
Kaynağın davranışı, container ayarları, portlar, replika sayısı vb.

---

#  Dersin 2. Saati – Uygulamalı Çalışma  
## Panel Üzerinden Workload Oluşturma

1. **Workload seçilir**
   - Deployment  
   - StatefulSet  
   - DaemonSet  
   - Job  
   - CronJob  

2. **Adlandırma yapılır**

3. **Replica (pod sayısı) belirlenir**

4. **Runtime Container Ayarları**
   - container name  
   - container image  
   - env variables  
   - ports  
   - lifecycle seçenekleri

5. **Health Check (Sağlık Kontrolleri)**
   - Liveness probe  
   - Readiness probe

6. **Service Ayarları**
   - service type seçimi  
   - port mapping  
   - dış dünyaya erişim seçenekleri

7. **Update Strategy**
   - Rolling Update → kesintisiz güncelleme  
   - Recreate → tüm podlar durur, yeniden oluşturulur

---

#  Senaryo 1: 3 Pod’luk Deployment (80 → 80 Port)

Hedef:
- 3 adet pod ayağa kalkacak  
- Container 80 portundan çalışacak  
- Service ile dış dünyaya 80 portundan yayın yapılacak  

Önemli Not:  
Deployment, podlar silinse veya bozulsa bile **ReplicaSet** sayesinde hemen yeniden oluşturur.

---

#  Senaryo 2: Frontend–Backend İletişimi

Elimizde iki ayrı deployment olsun:

- frontend  
- backend  

Doğru mimari:

```
frontend → service → backend
             ↓
        backend pod IP
```

###  Neden Service?
- Pod IP’leri **kalıcı değildir**  
- Pod ölürse yeni pod **yeni IP** alır  
- Frontend → backend pod IP’si üzerinden erişim yapıyorsa bağlantı kopar  

Service, backend için **sabit bir sanal IP (ClusterIP)** sağlar.  
Bu sayede pod IP değişse bile iletişim hiç kopmaz.

---

# Örnek Çalışmalar (Ders İçinde Yapılan YAML’lar)

Aşağıdaki YAML türleri derste yazılmış ve panele uygulanmıştır:

- **Deployment**  
- **Service**  
- **DaemonSet**  
- **Job**  
- **StatefulSet**  
- **CronJob**

Örnek Deployment:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demo-deploy
spec:
  replicas: 3
  selector:
    matchLabels:
      app: demo
  template:
    metadata:
      labels:
        app: demo
    spec:
      containers:
        - name: demo
          image: nginx
          ports:
            - containerPort: 80
```

Örnek Service:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: demo-svc
spec:
  type: ClusterIP
  selector:
    app: demo
  ports:
    - port: 80
      targetPort: 80
```

---

#  Özet

Bu derste aşağıdaki beceriler kazanılmıştır:

- Pod, YAML, lifecycle, workload ve service temelleri  
- Kubernetes paneli üzerinden workload oluşturma  
- Deployment ile otomatik self-healing mantığı  
- Service ile uygulamalar arası kalıcı iletişim sağlama  
- Frontend–backend iletişim tasarımı  
- YAML dosyalarının metadata–kind–spec yapısının analizi  
- Çeşitli workload türleri ve kullanım alanları  
- Uygulamalı scenario çalışmaları ile teorinin pekiştirilmesi  




