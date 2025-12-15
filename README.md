# Huawei CCE Kubernetes Bootcamp

Bu repo, **Huawei Cloud CCE Kubernetes Bootcamp** süresince oluşturulan tüm notları, YAML manifestlerini, uygulamaları ve final raporlarını içermektedir.

Amaç; Kubernetes temelinden ileri konulara kadar yapılan tüm çalışmaların **düzenli, tekrar edilebilir ve GitHub üzerinden paylaşılabilir** bir referans haline getirilmesidir.

---

## 📁 Klasör Yapısı

```
Huawei-CCE-Kubernetes-Bootcamp/
│
├── Hafta-1-Kubernetes-Temelleri/
│   ├── Ders-1-Docker-Kubernetes/
│   │   ├── docker.md
│   │   └── kubernetes-cluster-connection.md
│   │
│   └── Ders-2-Pod-Deployment-Service-Namespace/
│       ├── YAML/
│       │   ├── cluster-cronjob.yaml
│       │   ├── cluster-daemonset.yaml
│       │   ├── cluster-deployment.yaml
│       │   ├── cluster-job.yaml
│       │   ├── cluster-service.yaml
│       │   └── cluster-statefulset.yaml
│       └── Pod-Deployment-Service-Namespace.md
│
├── Hafta-2-Konfigurasyon-Ve-Veri-Yonetimi/
│   ├── Ders-3-ConfigMap-Secret-RBAC/
│   │   ├── RBAC-YAML/
│   │   │   ├── 0-rbac-tanimi.txt
│   │   │   ├── 1-rbac-service-account.yaml
│   │   │   ├── 2-rbac-role.yaml
│   │   │   ├── 3-rbac-rolebinding.yaml
│   │   │   ├── 4-rbac-clusterrole.yaml
│   │   │   └── 5-rbac-clusterrolebinding.yaml
│   │   └── ConfigMap-Secret-RBAC-Ve-Deployment-Senaryolari.md
│   │
│   └── Ders-4-Storage-Helm/
│       └── Ephemeral-Volume-PV-PVC-StorageClass-Helm.md
│
├── Hafta-3-Izleme-Ve-Container-Yonetimi/
│   ├── Ders-5-Container-Image-Ve-Build-Surecleri/
│   │   ├── Uygulama-1-Single-Stage/
│   │   │   ├── Dockerfile
│   │   │   ├── main.go
│   │   │   ├── 1-SingleStageRun.png
│   │   │   └── 2-SingleStageSize.png
│   │   │
│   │   ├── Uygulama-2-Multi-Stage/
│   │   │   ├── dockerfile
│   │   │   ├── main.go
│   │   │   ├── 3-MultiStageRun.png
│   │   │   └── 4-MultiStageSize.png
│   │   │
│   │   ├── Uygulama-3-SWR/
│   │   │   └── registry.txt
│   │   │
│   │   ├── 5-Containers.png
│   │   └── Container-Image-Single-Multi-Stage-SWR-Kubectl.md
│   │
│   └── Ders-6-AOM-Ve-Monitoring/
│       └── AOM-Monitoring.md
│
├── Hafta-4-Ag-Yonetimi-Ve-Sertifikasyon/
│   ├── Ders-7-Concept-Ve-Ingress-Controller/
│   │   ├── k8s-Ingress-Examples/
│   │   │   ├── 1-ingress-basic.yaml
│   │   │   ├── 2-ingress-multi-path.yaml
│   │   │   ├── 3-ingress-multi-host.yaml
│   │   │   ├── 4-ingress-tls.yaml
│   │   │   ├── 5-ingress-rewrite-path-manipulation.yaml
│   │   │   ├── 6a-ingress-stable.yaml
│   │   │   ├── 6b-ingress-canary.yaml
│   │   │   └── 7-ingress-security.yaml
│   │   └── Concept-Ingress-Controller.md
│   │
│   └── Ders-8-Ingress-Demo/
│       ├── YAMLs/
│       │   ├── api-deployment.yaml
│       │   ├── api-service.yaml
│       │   ├── hello-deployment.yaml
│       │   ├── hello-service.yaml
│       │   ├── web-deployment.yaml
│       │   ├── web-service.yaml
│       │   └── ingress.yaml
│       └── Ingress-Demo.md
│
└── Labs-Ve-Final-Project/
    ├── 1-Huawei-Kubernetes-CCE-Bootcamp-Lab-1.pdf
    ├── 2-Huawei-Kubernetes-CCE-Bootcamp-Lab-2.pdf
    └── 3-Final-Project-Ingress-Demo.pdf
```

---

## 🎯 Repo İçeriği

* Kubernetes temel objeleri (Pod, Deployment, Service, Namespace)
* ConfigMap, Secret ve RBAC senaryoları
* Ephemeral Volume, PV, PVC, StorageClass
* Single-stage & Multi-stage container build süreçleri
* Huawei SWR entegrasyonu
* AOM ve monitoring yaklaşımları
* Ingress Controller kavramları ve gerçek demo uygulamaları
* Bootcamp lab dokümanları ve final proje raporu

---



---

> 📌 Not: Bu repo, eğitim süresince birebir yapılan uygulamalara ve gerçek konfigürasyonlara dayanmaktadır.
