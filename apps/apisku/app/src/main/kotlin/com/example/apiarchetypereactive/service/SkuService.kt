package com.example.apiarchetypereactive.service

import com.example.apiarchetypereactive.component.SkuComponent
import com.example.apiarchetypereactive.model.dto.SkuDTO
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Service

@Service
class SkuService(
    private val skuComponent: SkuComponent,
) {

    private val log: Logger = LoggerFactory.getLogger(javaClass)
    fun findById(id: String): SkuDTO {
        log.info("buscando por id: $id")
        return skuComponent.findById(id)
    }


}
