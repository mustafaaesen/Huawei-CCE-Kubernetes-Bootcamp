# Week 2 – Video 1  
## Kubernetes ConfigMap, Secret, RBAC ve Deployment Senaryoları  
---

# 1) Uygulamalar Konfigürasyon Bilgilerini Nasıl Alır?

- Konfigürasyon değerlerini doğrudan kod içine yazmak (hard-coded) **güvenli değildir**.
- CI/CD süreçlerinde değişiklik yapmak zordur.
- Deployment YAML içerisine hard-coded environment variable yazmak bakım açısından **kötü pratiktir**.
- Çözüm: **ConfigMap**
  - Uygulamanın soft (dinamik değişebilen) ayarlarını dışarıdan yönetmemizi sağlar.

---

# 2) ConfigMap Mantığı

- ConfigMap, uygulamaya dışarıdan konfigürasyon enjekte etmek için kullanılır.
- Örnek ayarlar: `DB_HOST`, `LOG_LEVEL`, `APP_COLOR`
- Kubernetes, ConfigMap içerisindeki değerleri Pod içinde environment variable olarak inject eder.

### Örnek Senaryo – Flask APP_COLOR
- APP_COLOR değeri kodda olursa değiştirmek için:
  - Image rebuild  
  - Repository'e push  
  - Yeni tag seçimi  
  - Rollout  
  gerekir.
- ConfigMap ile sadece **value değiştirilir → redeploy** yapılır → hızlı.

---

# 3) WebApp (Flask) ConfigMap Senaryosu

### Namespace Oluşturma
Console:
- `webapp-color` isminde namespace oluşturuldu.

Terminal:
```bash
kubectl create namespace webapp-color
```

### Workload Oluşturma
- Image seçildi.
- Environment Variable → Custom:
  ```
  APP_COLOR = darkblue
  ```

### ConfigMap Oluşturma

Console:
- Name: `webapp-color`
- Key: `appcolor`
- Value: `darkblue`

Terminal:
```bash
kubectl create configmap webapp-cm \
  --from-literal=APP_COLOR=darkblue \
  -n webapp-color
```

### Deployment YAML İncelemesi
- Hard-coded değer yok.
- ConfigMap referansı kullanılmış.

### Service
- ClusterIP: TCP 8080 → 8080
- Dış erişim için NodePort eklendi.

---

# 4) Erişim Testi

Uygulama tarayıcıdan erişildi:

```
http://NODE_PUBLIC_IP:8080
```

---

# 5) ConfigMap Güncelleme Senaryosu

Renk değiştirildi:
```
APP_COLOR: darkblue → red
```

Workload → More → Redeploy  
Yeni renk uygulandı.

---

# 6) Secret Mantığı

- Secret da environment variable inject eder.
- Farkı:
  - Veriler etcd’ye **şifrelenmiş (encrypted)** yazılır.
  - `kubectl describe` ile değerler **görülemez**.
- Kullanımı:
  - Şifre, token, DB kullanıcı adı, DB password gibi gizli bilgiler.

### Güvenlik Senaryosu
Birisi cluster’a erişse bile:
```
kubectl describe configmap ...
```
ile ConfigMap değerlerini görebilir.

Ama Secret değerlerini **göremez**.

RBAC ile sadece yetkili kullanıcılar Secret okuyabilir.

---

# 7) MySQL Bağlantılı Flask Uygulaması Senaryosu

### Adımlar

1) Yeni namespace oluşturuldu.  
2) Deployment’tan önce ConfigMap ve Secret’lar oluşturuldu.

### ConfigMap
- DB host bilgisi:
  ```
  DB_HOST = mysql-service.default.svc.cluster.local
  ```
- `db_database` ConfigMap’te tutuldu.

### Secret
- `DB_USER`, `DB_PASSWORD` Secret'ta tutuldu.

### MySQL Deployment
- Eğitim amacıyla StatefulSet yerine Deployment kullanıldı.
- Secret içindeki env var’lar otomatik inject edildi.

Service:
- ClusterIP: TCP 3306 → 3306  
- Dış erişim için NodePort **eklenmedi**.

### Flask Uygulaması Deployment
- ConfigMap + Secret oluşturuldu.
- 4 env:
  ```
  DB_DATABASE
  DB_HOST
  DB_PASSWORD
  DB_USER
  ```

Service:
- 8080 → 9090  
- NodePort ile dışarıdan erişildi.

---

# 8) RBAC Mantığı

- RBAC, Kubernetes’te kullanıcı veya uygulamanın **hangi kaynaklarda hangi işlemleri** yapabileceğini belirler.
- Videoda ServiceAccount, Role, RoleBinding, ClusterRole ve ClusterRoleBinding örnekleri uygulandı.

---




