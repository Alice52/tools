package cn.edu.ntu.boot.encrypt;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

/**
 * @author zack <br>
 * @create 2020-05-03 21:23 <br>
 */
@EnableSwagger2
@SpringBootApplication
public class ToolEncryptApplication {

    public static void main(String[] args) {
        SpringApplication.run(ToolEncryptApplication.class, args);
    }
}
