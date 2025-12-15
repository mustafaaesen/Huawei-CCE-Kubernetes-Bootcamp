# Hafta 2 – Video 2 Notları  
## Ephemeral Volume – PV – PVC – StorageClass – Helm  


Ephemeral volume: Pod/container yaşam döngüsüne bağlı olan ve pod yeniden oluşturulduğunda tamamen silinen geçici dosya sistemi. `/tmp` gibi yerlere yazılan veriler pod ölünce yok olur. Bu davranış Kubernetes’in varsayılan çalışma mantığıdır.

Deployment ile pod oluşturulduğunda içine girilip dosya yazılabilir:

kubectl create deployment test-nginx --image=nginx
kubectl get pods -n default
kubectl exec -it -n default <pod> -- sh
echo "Hello CCE" > /tmp/test.txt
cat /tmp/test.txt

Pod silindiğinde filesystem sıfırdan yaratılır:

kubectl delete pod <pod> -n default
kubectl get pods -n default
kubectl exec -it -n default <yeni-pod> -- sh
cat /tmp/test.txt

Dosya bulunmaz → ephemeral storage.

Ephemeral volume seçenekleri:
1) Pod-level mount → pod ölünce volume gider, yeniden yaratılınca boş gelir.  
2) Node-level mount → veri node’a bağlıdır; node ölürse/veri kaybolursa erişilemez; pod başka node’a schedule edilince kullanılamaz.  
Her iki yöntem de kalıcı veri saklama ihtiyacı olan uygulamalar için uygun değildir.

Kalıcı depolama için PV ve PVC kullanılır.

PV (Persistent Volume): Cluster içinde fiziksel veya sanal kalıcı disk kaynağının tanımı. (OBS / EVS / SSD / HDD vs.)

PVC (Persistent Volume Claim): Pod’un PV’den disk talep etmesi.  
PV + PVC birlikte kullanıldığında pod ölse bile veri korunur.

Büyük cluster’larda her PVC için PV yazmak zahmetlidir.  
StorageClass bu problemi çözer.

StorageClass: Farklı disk tiplerinin şablonlandığı yapı. PVC oluşturulduğunda otomatik PV üretir (dynamic provisioning).  
Böylece yüzlerce PV dosyası yazma ihtiyacı ortadan kalkar.

Stateful uygulamalar (MySQL, Redis vb.) için StatefulSet workload tipi kullanılır.  
MySQL’in veri klasörü `/var/lib/mysql` dizinidir.  
Bu path’e PVC mount edildiğinde veri kalıcı hale gelir.

WordPress + MySQL mimarisinde kullanılan yapılar:
- WordPress Deployment (stateless)  
- MySQL StatefulSet (stateful)  
- ConfigMap (gizli olmayan environment değerleri)  
- Secret (şifreler gibi hassas değerler)  
- PV/PVC (kalıcı MySQL verisi için)  
- StorageClass (otomatik PV oluşumu için)  
- Service (WordPress–MySQL bağlantısı)  
- Headless Service (StatefulSet için DNS çözümlemesi)  

Environment örnekleri:
MYSQL_DATABASE  
MYSQL_USER  
WP_DB_PASSWORD  
MYSQL_PASSWORD  

PVC MySQL’e şu path ile bağlanır:

/var/lib/mysql

Helm:  
Kubernetes manifest dosyalarının template haline getirilmiş paket yapısı.  
Tüm YAML dosyaları chart içinde bulunur.  
values.yaml → uygulamanın değişkenlerinin tutulduğu dosya.  
Büyük uygulamaların yönetimini kolaylaştırır, tek yerden güncellenebilir.

Genel kavram özeti:  
- Ephemeral volume → geçici  
- PV → kalıcı disk  
- PVC → diski talep eden istek  
- StorageClass → otomatik PV üretir  
- StatefulSet → stateful workload  
- ConfigMap/Secret → uygulama ayarları  
- Helm → çoklu YAML şablon yönetimi

