import React, { useState, useEffect } from "react"
import { Container, Row } from "react-bootstrap"
import ProductPost from "../../components/ProductPost/ProductPost"
import Navbar from "../../components/Navbar/Navbar"
import CartModal from "../../components/CartModal/CartModal"
import FooterPage from "../../components/Footer/Footer"

import {upsertCartService, getCartService, placeOrderService, getProductListService, getProductService} from "./helper"

export default function HomePage() {
    const [productList, setProductList] = useState([])
    const [cart, setCart] = useState({
        ID: "",
        totalPrice: "",
        paymentMethod: "",
        lineItems: []
    })
    const [showCart, setShowCart] = useState(false)


    //* Only Fetch product List when it is empty
    if (productList.length === 0){
        getProductListService().then(products=>{
            console.log(products)
            setProductList(products)
        })
    }

    async function addToCart(productId) {
        //TODO: Get Product Detail from Product / Get Product entirely
        const product = await getProductService(productId)

        console.log("Add TO Cart:", productId)
        console.log(product)
        console.log( {
            "description": product.description,
            "price": product.price,
            "quantity": 1,
            "title": product.title,
            "product_id": product.id
        })

        var newLineItems = cart.lineItems
        newLineItems.push(
            {
                "description": product.description,
                "price": product.price,
                "quantity": 1,
                "title": product.title,
                "product_id": product.id
            }
        )

        console.log( newLineItems)

        const updatedCart = await upsertCartService(cart, newLineItems, "")

        console.log("updatedCart:", updatedCart)

        // Add Item to Cart
        setCart({
            ...updatedCart
        })
        // Show Cart when Add Item
        setShowCart(true)
    }

    // TODO: REMAKE THIS FUNCTION 
    function removeFromCart(removeIndex) {
        setCart(prevItem => prevItem.filter((_, idx) => idx !== removeIndex))
    }

    function onHideCartModalHandler() {
        setShowCart(false)
    }

    async function showCartModalHanlder() {
        const newCart = await getCartService(cart.ID)
        if (newCart !== null) {
            setCart(newCart)
        }
        setShowCart(true)
    }

    async function addPaymentMethod(paymentMethod) {
        const updatedCart = await upsertCartService(cart, null, paymentMethod)
        setCart({
            ...updatedCart
        })
    }

    async function checkoutHandler(cartID) {
        // TODO: Checkout Handler
        console.log("Checkout Cart:", cartID)

        const checkoutResponse = await placeOrderService(cartID)
        console.log(checkoutResponse)

    }

    return (
        <>
            <Navbar
                removeFromCart={removeFromCart}
                showCartModalHanlder={showCartModalHanlder}

            />
            <CartModal
                cart={cart}
                show={showCart}
                onHide={onHideCartModalHandler}
                addPaymentMethod={addPaymentMethod}
                removeFromCart={removeFromCart}
                checkout={checkoutHandler}
            />
            <Container>
                <Row>
                    {productList.map((product, idx) => {
                        console.log("productList", idx)
                        return (
                        <ProductPost
                            key={idx}
                            id={product.id}
                            title={product.title}
                            description={product.description}
                            imageURL={product.imageURL}
                            addToCart={addToCart}
                        />)
                        })}
                </Row>
            </Container>
            <FooterPage />
        </>
    )
}