# 🚀 Huawei Cloud CCE — Kubernetes Cluster Connection Guide

Bu doküman Huawei Cloud üzerinde **CCE (Cloud Container Engine)** kullanarak Kubernetes cluster oluşturma ve yerel makine üzerinden `kubectl` ile bağlantı kurma sürecini anlatır.

---

## 📌 1. Kubernetes Cluster Oluşturma (CCE)

Huawei Cloud panelinden:

1. **Cloud Container Engine (CCE)** servisine girilir.  
2. **Create Cluster** ile yeni Kubernetes cluster oluşturulur.  
3. CCE Standard seçilir, Kubernetes versiyonu ayarlanır.  
4. Network ve diğer ayarlar varsayılan bırakılabilir.

---

## 📌 2. Node Pool Oluşturma (AZ1 ve AZ2)

Cluster oluşturulduktan sonra:

1. Sol menüden **Node Pools** seçilir.  
2. İki adet node pool oluşturulur:
   - **AZ1**
   - **AZ2**
3. Her node pool’a **1 adet ECS (Pay-Per-Use)** atanır.
4. ECS sunucular Huawei tarafından **otomatik** oluşturulur.

---

## 📌 3. EIP Bağlama ve Kubeconfig İndirme

Kubectl’in cluster’a internet üzerinden bağlanabilmesi için:

1. CCE → **Clusters** üzerinden cluster seçilir.  
2. **Information → Access/Connection** bölümüne girilir.  
3. API Server’a **Elastic IP (EIP)** bağlanır.  
4. **kubeconfig YAML dosyası** indirilir.

Bu dosya yerel kubectl bağlantısı için gereklidir.

---

## 📌 4. Ubuntu Üzerinde Kubectl Kurulumu

### ✔ Kubectl kur

```bash
sudo snap install kubectl --classic
```

---

## 📌 5. Kubeconfig Dosyasını Yerleştirme

İndirilen dosyayı `.kube/config` konumuna taşı:

```bash
mkdir -p ~/.kube
mv ~/İndirilenler/test-cluster-kubeconfig.yaml ~/.kube/config
```

> Not: İngilizce Ubuntu kullananlar için klasör: `~/Downloads`

---

## 📌 6. Kubectl Context Kontrolü

Context listesini görüntüle:

```bash
kubectl config get-contexts
```

Örnek:

```
CURRENT   NAME                CLUSTER                    AUTHINFO
*         external            externalCluster            user
          externalTLSVerify   externalClusterTLSVerify   user
          internal            internalCluster            user
```

External erişim için:

```bash
kubectl config use-context external
```

---

## 📌 7. Cluster Bağlantısını Test Etme

### Kubernetes API Server bilgisi:

```bash
kubectl cluster-info
```

Örnek:

```
Kubernetes control plane is running at https://<EIP>:5443
CoreDNS is running at ...
```

### Node’ları görüntüle:

```bash
kubectl get nodes
```

Örnek:

```
NAME            STATUS   ROLES    AGE   VERSION
192.168.0.125   Ready    <none>   40m   v1.32.6
192.168.0.246   Ready    <none>   49m   v1.32.6
192.168.0.25    Ready    <none>   50m   v1.32.6
```

Node’lar **Ready** durumunda ise bağlantı başarılıdır.

---

## 🎉 Artık Hazırsınız

Artık cluster üzerinde Kubernetes kaynakları oluşturabilirsiniz:

```bash
kubectl get pods -A
kubectl create deployment nginx-deploy --image=nginx
kubectl expose deployment nginx-deploy --type=LoadBalancer --port=80
```

---



