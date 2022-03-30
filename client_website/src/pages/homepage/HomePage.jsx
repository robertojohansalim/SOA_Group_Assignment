import React, { useState } from "react"
import { Container, Row } from "react-bootstrap"
import ProductPost from "../../components/ProductPost/ProductPost"
import Navbar from "../../components/Navbar/Navbar"
import CartModal from "../../components/CartModal/CartModal"
import FooterPage from "../../components/Footer/Footer"
import axios from "axios"

async function upsertCartService(cart, newLineItems, paymentMethod) {
    const host = process.env.REACT_APP_CART_SERVICE_HOST
    const url = host + "/api/upsert_cart"

    if (newLineItems === null) {
        newLineItems = cart.lineItems
    }

    var data = {
        ...cart,
        "lineItems": newLineItems
    }

    data["paymentMethod"] = paymentMethod

    // // TODO::REMOVE kalau sudah support pilih Payment? (atau buang di payment service)
    // data["paymentMethod"] = "BANK_TRANSFER"
    console.log("Update Cart", data)
    var axiosResponse = await axios({
        method: 'POST',
        url: url,
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
        data: data
    })
    return axiosResponse.data
}

async function getCartService(cartID) {
    if (cartID === "") {
        return null
    }
    const host = process.env.REACT_APP_CART_SERVICE_HOST
    const url = host + `/api/get_cart/${cartID}`

    console.log({
        "ID": cartID,
    })

    console.log("axiosResponse", url)
    var axiosResponse = await axios({
        method: 'GET',
        url: url,
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
    })

    console.log("axiosResponse", axiosResponse)

    var cartResponse = axiosResponse.data.cart

    // If Return null set as Zeroed Value
    if (!cartResponse) {
        cartResponse = {
            ID: "",
            totalPrice: "",
            paymentMethod: "",
            lineItems: []
        }
    }

    return cartResponse
}

async function placeOrderService(cartID) {
    const host = process.env.REACT_APP_CART_SERVICE_HOST
    const url = host + "/api/place_order"

    console.log({
        "ID": cartID,
        "action": "CHECKOUT"
    })

    var axiosResponse = await axios({
        method: 'POST',
        url: url,
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
        data: {
            "ID": cartID,
            "action": "CHECKOUT"
        }
    })
    return axiosResponse.data
}

export default function HomePage() {
    // const [cart, setCart] = useState([])
    const [cart, setCart] = useState({
        ID: "",
        totalPrice: "",
        paymentMethod: "",
        lineItems: []
    })
    const [showCart, setShowCart] = useState(false)

    // TODO: Fetch From Product Service
    var productList = [
        {
            "id": "product-id-1",
            "title": "Lorem ipsum 1",
            "description": "Lorem Ipsum Dolor Si Amet",
            "imageURL": "https://picsum.photos/200"
        },
        {
            "id": "product-id-2",
            "title": "Lorem ipsum 2",
            "description": "Lorem Ipsum Dolor Si Amet",
            "imageURL": "https://picsum.photos/200"
        },
        {
            "id": "product-id-3",
            "title": "Lorem ipsum 3",
            "description": "Lorem Ipsum Dolor Si Amet",
            "imageURL": "https://picsum.photos/200"
        },
        {
            "id": "product-id-4",
            "title": "Lorem ipsum 4",
            "description": "Lorem Ipsum Dolor Si Amet test",
            "imageURL": "https://picsum.photos/200"
        }
    ]

    async function addToCart(productId) {
        //TODO: Get Product Detail from Product / Get Product entirely

        var newLineItems = cart.lineItems
        newLineItems.push(
            {
                "description": "Item 1 Very Long Description",
                "price": 15000,
                "quantity": 1,
                "title": productId,
                "product_id": productId
            }
        )

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
                    {productList.map((product, idx) => (
                        <ProductPost
                            key={idx}
                            id={product.id}
                            title={product.title}
                            description={product.description}
                            imageURL={product.imageURL}
                            addToCart={addToCart}
                        />))}
                </Row>
            </Container>
            <FooterPage />
        </>
    )
}