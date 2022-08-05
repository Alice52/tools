FROM openjdk:8-jre-slim-stretch
VOLUME /tmp

ARG JASYPT_ENCRYPTOR_PASSWORD=""

ADD target/**.jar app.jar
RUN sh -c 'touch /app.jar'
ENV JAVA_OPTS=""
EXPOSE  8080
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -DJASYPT_ENCRYPTOR_PASSWORD=${JASYPT_ENCRYPTOR_PASSWORD} -jar /app.jar"]
