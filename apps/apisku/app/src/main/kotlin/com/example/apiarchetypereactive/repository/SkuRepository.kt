package com.example.apiarchetypereactive.repository

import com.example.apiarchetypereactive.model.dto.SkuDTO
import org.springframework.stereotype.Repository

@Repository
class SkuRepository {

    private val skuList = setOf(
        SkuDTO(
            id = "1",
            productId = "1",
            type = "pc",
        ),
        SkuDTO(
            id = "2",
            productId = "2",
            type = "phone",
        ),
        SkuDTO(
            id = "3",
            productId = "3",
            type = "tv",
        ),
    )

    fun getById(id: String): SkuDTO? {
        return skuList.firstOrNull { it.id == id }
    }


}