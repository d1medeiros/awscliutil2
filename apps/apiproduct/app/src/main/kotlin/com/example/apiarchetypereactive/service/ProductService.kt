package com.example.apiarchetypereactive.service

import com.example.apiarchetypereactive.client.http.SKUClient
import com.example.apiarchetypereactive.client.http.SellerClient
import com.example.apiarchetypereactive.component.ProductComponent
import com.example.apiarchetypereactive.model.Product
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Service

@Service
class ProductService(
    private val productComponent: ProductComponent,
    private val skuClient: SKUClient,
    private val sellerClient: SellerClient
) {

    private val log: Logger = LoggerFactory.getLogger(javaClass)
    fun findById(id: String): Product {
        log.info("buscando por id: $id")
        return productComponent.findById(id).apply {
            this.skus = setOf(skuClient.findById(this.skuId))
//            this.sellers = this.sellers.map {
//                sellerClient.findById(it.id)
//            }.toSet()
        }
    }

    fun findAll(skuId: String?): Set<Product> {
        return productComponent.findAll(skuId)
    }

}
