package cn.edu.ntu.boot.encrypt;

import common.redis.config.ExcludeRedisConfig;
import common.swagger.annotation.EnableSwagger;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.amqp.RabbitAutoConfiguration;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.context.annotation.Import;

/**
 * @author zack <br>
 * @create 2020-05-03 21:23 <br>
 */
@EnableSwagger
@Import(ExcludeRedisConfig.class)
@SpringBootApplication(exclude = {DataSourceAutoConfiguration.class, RabbitAutoConfiguration.class})
public class ToolEncryptApplication {

    public static void main(String[] args) {
        SpringApplication.run(ToolEncryptApplication.class, args);
    }
}
