package cn.edu.ntu.boot.encrypt;

import common.swagger.annotation.EnableSwagger;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * @author zack <br>
 * @create 2020-05-03 21:23 <br>
 */
@EnableSwagger
@SpringBootApplication
public class ToolEncryptApplication {

    public static void main(String[] args) {
        SpringApplication.run(ToolEncryptApplication.class, args);
    }
}
