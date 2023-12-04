package io.github.alice52

import common.swagger.annotation.EnableSwagger
import io.github.alice52.common.inject.annotation.SimpleBootApplication
import org.springframework.boot.runApplication

@EnableSwagger
@SimpleBootApplication
open class ToolEncryptApplication

fun main(vararg args: String) {

    runApplication<ToolEncryptApplication>(*args)
}