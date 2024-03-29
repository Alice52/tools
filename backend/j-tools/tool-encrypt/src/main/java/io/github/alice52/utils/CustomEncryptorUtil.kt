package io.github.alice52.utils

import org.jasypt.encryption.StringEncryptor
import org.jasypt.encryption.pbe.PooledPBEStringEncryptor
import org.jasypt.encryption.pbe.config.SimpleStringPBEConfig

object CustomEncryptorUtil {

    fun stringEncryptor(): StringEncryptor {
        val encryptor = PooledPBEStringEncryptor()
        val config = SimpleStringPBEConfig()
        config.password = "CodeSheep"
        config.algorithm = "PBEWITHHMACSHA512ANDAES_256"
        config.setKeyObtentionIterations("1000")
        config.setPoolSize("1")
        config.providerName = "SunJCE"
        config.setSaltGeneratorClassName("org.jasypt.salt.RandomSaltGenerator")
        config.setIvGeneratorClassName("org.jasypt.iv.RandomIvGenerator")
        config.stringOutputType = "base64"
        encryptor.setConfig(config)
        return encryptor
    }
}