package com.example.apiarchetypereactive.controller

import com.example.apiarchetypereactive.model.dto.SkuDTO
import com.example.apiarchetypereactive.service.SkuService
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("skus")
class SkuController(
    private val skuService: SkuService
) {
    val log: Logger = LoggerFactory.getLogger(javaClass)


    @GetMapping("/{id}")
    fun findById(@PathVariable("id") id: String): SkuDTO {
        log.info("buscando sku $id")
        return skuService.findById(id)
    }


}