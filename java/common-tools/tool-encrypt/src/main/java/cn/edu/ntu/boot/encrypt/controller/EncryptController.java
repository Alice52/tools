package cn.edu.ntu.boot.encrypt.controller;

import org.jasypt.encryption.StringEncryptor;
import org.springframework.context.ApplicationContext;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

/**
 * @author zack <br>
 * @create 2020-05-03 22:04 <br>
 */
@RestController
public class EncryptController {

    @Resource private ApplicationContext appCtx;

    @Resource private StringEncryptor encryptBean;

    @GetMapping("/encrypt")
    public String doEncrypt(@RequestParam("encryptWord") String encryptWord) {

        return encrypt(encryptWord);
    }

    @GetMapping("/decrypt")
    public String doDecrypt(@RequestParam("decryWord") String decryWord) {

        return decrypt(decryWord);
    }

    private String encrypt(String originPassword) {
        return encryptBean.encrypt(originPassword);
    }

    private String decrypt(String encryptedPassword) {
        return encryptBean.decrypt(encryptedPassword);
    }
}
