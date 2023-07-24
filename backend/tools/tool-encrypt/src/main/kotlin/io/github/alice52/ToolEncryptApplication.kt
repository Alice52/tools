package io.github.alice52

import common.redis.config.ExcludeRedisConfig
import common.swagger.annotation.EnableSwagger
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.autoconfigure.amqp.RabbitAutoConfiguration
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Import

@EnableSwagger
//@SimpleBootApplication
@Import(ExcludeRedisConfig::class)
@SpringBootApplication(exclude = [DataSourceAutoConfiguration::class, RabbitAutoConfiguration::class])
open class ToolEncryptApplication

fun main(vararg args: String) {

    runApplication<ToolEncryptApplication>(*args)
}

// logger()
inline fun <reified R : Any> R.logger(): Logger =
    LoggerFactory.getLogger(this::class.java.name.substringBefore("\$Companion"))
