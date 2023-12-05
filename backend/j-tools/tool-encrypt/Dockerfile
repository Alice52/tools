FROM openjdk:17
# FROM java:openjdk-8u91-jdk
# FROM openjdk:8u121-jre-alpine
# FROM hengyunabc/arthas:latest
VOLUME /app

ARG JASYPT_ENCRYPTOR_PASSWORD=""
ARG COMMON_SERVICE_IP=""

# RUN wget https://arthas.aliyun.com/arthas-boot.jar
COPY --from=hengyunabc/arthas:latest /opt/arthas /opt/arthas

RUN sh -c 'ls -la'

# otherwise target is not exist
ADD ./ ./
ADD target/**.jar /app.jar
ENV JAVA_OPTS=""
EXPOSE  8080

# Build a shell script because the ENTRYPOINT command doesn't like using ENV
RUN sh -c 'echo -e "#!/bin/sh\n  java ${JAVA_OPTS} -Djava.awt.headless=true -Djava.security.egd=file:/dev/./urandom -DJASYPT_ENCRYPTOR_PASSWORD=${JASYPT_ENCRYPTOR_PASSWORD} -Dsun.net.httpserver.nodelay=true -jar /app.jar" > /entrypoint.sh'
RUN chmod +x /entrypoint.sh
RUN cat ./entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
