package io.github.alice52.controller

import org.jasypt.encryption.StringEncryptor
import org.springframework.context.ApplicationContext
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
class EncryptController(val appCtx: ApplicationContext, val encryptBean: StringEncryptor) {

    @GetMapping("/encrypt")
    fun doEncrypt(@RequestParam("encryptWord") encryptWord: String) = encrypt(encryptWord)

    @GetMapping("/decrypt")
    fun doDecrypt(@RequestParam("decryWord") decryWord: String) = decrypt(decryWord)

    private fun encrypt(encryptWord: String) = encryptBean.encrypt(encryptWord);
    private fun decrypt(encryptedPassword: String) = encryptBean.decrypt(encryptedPassword)

}