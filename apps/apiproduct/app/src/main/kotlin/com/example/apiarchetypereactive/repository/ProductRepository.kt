package com.example.apiarchetypereactive.repository

import com.example.apiarchetypereactive.model.Product
import com.example.apiarchetypereactive.model.dto.SellerDTO
import com.example.apiarchetypereactive.model.dto.SkuDTO
import org.springframework.stereotype.Repository

@Repository
class ProductRepository {

    private val productList = setOf(
        Product(
            id = "1",
            name = "notebook a",
            skus = emptySet(),
            skuId = "1",
            sellers = setOf(
                SellerDTO(
                    id = "1"
                )
            )
        ),
        Product(
            id = "2",
            name = "iphone 11",
            skus = emptySet(),
            skuId = "2",
            sellers = setOf(
                SellerDTO(
                    id = "1"
                )
            )
        ),
        Product(
            id = "3",
            name = "tv sony",
            skus = setOf(
                SkuDTO(
                    id = "4"
                )
            ),
            skuId = "3",
            sellers = setOf(
                SellerDTO(
                    id = "2"
                )
            )
        ),
        Product(
            id = "4",
            name = "notebook",
            skus = emptySet(),
            skuId = "1",
            sellers = setOf(
                SellerDTO(
                    id = "3"
                )
            )
        ),
        Product(
            id = "5",
            name = "iphone 15",
            skus = emptySet(),
            skuId = "2",
            sellers = setOf(
                SellerDTO(
                    id = "3"
                )
            )
        ),
    )

    fun getById(id: String): Product? {
        return productList.firstOrNull { it.id == id }
    }

    fun getAll(skuId: String?): Set<Product> {
        return productList.filter {
            it.skus.any {
                when {
                    skuId != null -> it.id == skuId
                    else -> true
                }
            }
        }.toSet()
    }
}