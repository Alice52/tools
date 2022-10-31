FROM java:openjdk-8u91-jdk
# FROM hengyunabc/arthas:latest
VOLUME /tmp

ARG JASYPT_ENCRYPTOR_PASSWORD=""
ARG COMMON_SERVICE_IP=""

# RUN wget https://arthas.aliyun.com/arthas-boot.jar
## COPY --from=hengyunabc/arthas:latest /opt/arthas /opt/arthas

ADD target/**.jar app.jar
RUN sh -c 'touch /app.jar'
ENV JAVA_OPTS=""
EXPOSE  8080
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -DCOMMON_SERVICE_IP=${COMMON_SERVICE_IP} -DJASYPT_ENCRYPTOR_PASSWORD=${JASYPT_ENCRYPTOR_PASSWORD} -jar /app.jar"]
