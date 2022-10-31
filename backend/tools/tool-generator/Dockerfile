FROM openjdk:8u121-jre-alpine
VOLUME /tmp
ARG MODULE_NAME=""
ARG JASYPT_ENCRYPTOR_PASSWORD=""
ENV JAVA_OPTS=""
ADD target/**.jar app.jar

# RUN wget https://arthas.aliyun.com/arthas-boot.jar
COPY --from=hengyunabc/arthas:latest /opt/arthas /opt/arthas

EXPOSE 8080
RUN echo -e "#!/bin/sh\njava ${JAVA_OPTS} -Djava.awt.headless=true -Djava.security.egd=file:/dev/./urandom -DJASYPT_ENCRYPTOR_PASSWORD=${JASYPT_ENCRYPTOR_PASSWORD} -Dsun.net.httpserver.nodelay=true -jar /app.jar" > ./entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
