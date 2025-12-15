
#----------- Stage 1: Build (Go compiler ortamı)--------------------------
FROM golang:1.22-alpine AS builder
# Go uygulamasını derlemek için compiler içeren bir base image 

WORKDIR /app
# Derleme işlemleri için çalışma dizini oluşturma.

COPY main.go .
# Go kaynak kodunu builder container içine kopyalama

RUN CGO_ENABLED=0 go build -o go_app main.go
# CGO kapalı şekilde uygulamanın build dosyasını alıyoruz daha düşük bianry çıkar
# Çıktı: /app/go_app dosyası
#------------------------------------------------------------------------------------------


# -----------------Stage 2: Runtime (Sadece çalıştırma için minimal ortam)------------------

FROM alpine:latest
# Artık compiler’a gerek yok bu yüzden küçük runtime image yeterli

RUN apk add --no-cache tini
# PID-1 sorunlarına karşın hafif init olan Tini ekleyerek çalıştırma

WORKDIR /app
# Çalışma ortamını oluşturma

COPY --from=builder /app/go_app .
# Builder stage’de üretilen derlenmiş binary’yi kopyalama

EXPOSE 8080
# Container'ın dışarıya açacağı port.

ENTRYPOINT ["/sbin/tini", "--", "/app/go_app"]
# Container başladığında uygulamanın tini üzerinde çalışması için girilen komut.


#Multi stage Dockerfile adımları 

#Stage1(builder) --------> Uygulama derleyicisi kodu build eder sonucunda büyük boyutlu image oluşur

#Stage2(runtime) ---------> SAdece binary gerekli gördüğü ortam image'ına koyar ve sadece image'ın gerekli olan araçlarını alır

#bu adımlar sonucunda tek bir dockerfile dosyası 2 ayrı ortamda derlenir sonunda daha düşük boyutlu bir image oluşur


#--->Docker önce uygulama kaynak dosyasını kopyalar builder image oluşturur(/app/go_app binary oluşur)
#--->Sonra runtime stage'de oluşan binary kopyalanarak sadece image için gerekli araçlar kullanıldığı bir binary oluşur
#--->Final image çok düşü boyutta oluşur 
#--->Container tini üzerinden güvenli şekilde başlar