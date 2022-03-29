import React, { useState } from "react"
import { Container, Row } from "react-bootstrap"
import ProductPost from "../../components/ProductPost/ProductPost"
import Navbar from "../../components/Navbar/Navbar"
import CartModal from "../../components/CartModal/CartModal"
import FooterPage from "../../components/Footer/Footer"

export default function HomePage() {
    const [cart, setCart] = useState([])
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
            "description": "Lorem Ipsum Dolor Si Amet",
            "imageURL": "https://picsum.photos/200"
        }
    ]

    function addToCart(productId) {
        setCart(prevItem => [
            ...prevItem,
            productId
        ])
    }

    function removeFromCart(removeIndex) {
        setCart(prevItem => prevItem.filter((_,idx)=> idx !== removeIndex))
    }

    function onHideCartModalHandler() {
        setShowCart(false)
    }

    function showCartModalHanlder() {
        setShowCart(true)
    }

    function checkoutHandler(cart){
        // TODO: Checkout Handler
        console.log("Checkout Cart:", cart)
    }

    return (
        <>
            <Navbar
                cart={cart}
                removeFromCart={removeFromCart}
                showCartModalHanlder={showCartModalHanlder}

            />
            <CartModal
                cart={cart}
                show={showCart}
                onHide={onHideCartModalHandler}
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