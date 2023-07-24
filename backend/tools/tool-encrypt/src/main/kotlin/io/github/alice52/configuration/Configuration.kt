package io.github.alice52.configuration

import common.core.configuration.RestTemplateConfig
import org.springframework.beans.factory.annotation.Value
import org.springframework.boot.autoconfigure.EnableAutoConfiguration
import org.springframework.boot.web.client.RestTemplateBuilder
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.MediaType
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter
import org.springframework.web.client.RestTemplate

@Configuration
@EnableAutoConfiguration(exclude = [RestTemplateConfig::class])
open class Configuration {

    /**
     * this value cannot set default value in here, must be in yml.
     */
    @Value("\${vault.config.token}")
    private lateinit var token: String

    @Value("\${vault.config.username:zack}")
    private lateinit var username: String

    /**
     * 会导致此 @Bean 不会执行<fucking issue></fucking>
     *
     * @see RestTemplateConfig.mappingJackson2HttpMessageConverter
     *
     * @return
     */
    @Bean
    open fun mappingJackson2HttpMessageConverter(): MappingJackson2HttpMessageConverter {
        val converter = MappingJackson2HttpMessageConverter()
        converter.supportedMediaTypes = listOf(MediaType.ALL)
        return converter
    }

    @Bean
    open fun restTemplate(converter: MappingJackson2HttpMessageConverter?): RestTemplate {
        return RestTemplateBuilder().additionalMessageConverters(converter).build()
    }
}
