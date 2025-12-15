# Go tabanlı hafif bir Linux imajı kullanıyoruz.
# Alpine küçük boyutlu olduğu için prod ortamlarda tercih edilir.
FROM golang:1.22-alpine

# 'tini' adlı küçük init process'i yükler.
# Kubernetes ortamında container'ın PID1 olarak düzgün çalışmasını sağlar.
# Sinyal (SIGTERM, SIGKILL) yönetimini düzeltir, zombie process oluşmasını engeller.
RUN apk add --no-cache tini

# Container içinde çalışma dizini oluştur
WORKDIR /app

# Lokal makinedeki main.go dosyasını container içindeki /app dizinine kopyala
COPY main.go .

# CGO_ENABLED=0  statik binary oluşturur (çok taşınabilir olur).

# go build -o go_app main.go  uygulamasını go_app adında çalıştırılabilir hale getirir.
RUN CGO_ENABLED=0 go build -o go_app main.go

# Uygulamanın container içinde dinleyeceği portu ayarla
EXPOSE 8080

# Container başlatıldığında tini çalışır ve uygulamayı kendi üzerinden başlatır.
# Bu, Kubernetes pod'larında graceful shutdown için gereklidir.
ENTRYPOINT ["/sbin/tini", "--", "/app/go_app"]


#container başladığında çalışacak komut

#Single Stage Dockerfile adımları

#--->Base image seç (uygulamanın sistemi lazım örneğin go)
#--->Container içerisinde çalışma klasörü belirle
#--->Kaynak kodu içine aktar
#--->Koddan çalıştırılabilir binary(build) dosyası üret
#--->Dinlenecek portu ayarla
#--->Container başladığında çalışacak komutu belirle
