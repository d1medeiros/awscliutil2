package com.example.apiarchetypereactive

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.context.properties.EnableConfigurationProperties
import org.springframework.boot.runApplication


@EnableConfigurationProperties
@SpringBootApplication
class SKUAPP

fun main(args: Array<String>) {

    runApplication<SKUAPP>(*args)
}
