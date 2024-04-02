package com.example.apiarchetypereactive.component

import com.example.apiarchetypereactive.model.dto.SkuDTO
import com.example.apiarchetypereactive.repository.SkuRepository
import org.springframework.stereotype.Component

@Component
class SkuComponent(
    private val skuRepository: SkuRepository
) {
    fun findById(id: String): SkuDTO {
        return skuRepository.getById(id)
            ?: throw RuntimeException("not found sku $id")
    }


}