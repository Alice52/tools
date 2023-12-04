FROM openjdk:17-jre-alpine
# FROM java:openjdk-8u91-jdk
# FROM openjdk:8u121-jre-alpine
VOLUME /tmp
ARG MODULE_NAME=""
ARG JASYPT_ENCRYPTOR_PASSWORD=""
ENV JAVA_OPTS=""
ADD **.jar app.jar

# RUN wget https://arthas.aliyun.com/arthas-boot.jar
COPY --from=hengyunabc/arthas:latest /opt/arthas /opt/arthas

EXPOSE 8080
RUN sh -c 'echo "#!/bin/sh\n  java ${JAVA_OPTS} -Djava.awt.headless=true -Djava.security.egd=file:/dev/./urandom -DJASYPT_ENCRYPTOR_PASSWORD=${JASYPT_ENCRYPTOR_PASSWORD} -Dsun.net.httpserver.nodelay=true -jar /app.jar" > ./entrypoint.sh'
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
