import axios from "axios"

async function upsertCartService(cart, newLineItems, paymentMethod) {
    const host = process.env.REACT_APP_ORCHESTRATOR_SERVICE_HOST
    const url = host + "/api/upsert_cart"

    if (newLineItems === null) {
        newLineItems = cart.lineItems
    }

    var data = {
        ...cart,
        "lineItems": newLineItems
        
    }

    if (paymentMethod !== ""){
        data["paymentMethod"]= paymentMethod
    }

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
    const host = process.env.REACT_APP_ORCHESTRATOR_SERVICE_HOST
    const url = host + `/api/get_cart`

    console.log({
        "ID": cartID,
    })

    console.log("axiosResponse", url)
    var axiosResponse = await axios({
        method: 'POST',
        url: url,
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
        data:{
            "ID":cartID
        }
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
    const host = process.env.REACT_APP_ORCHESTRATOR_SERVICE_HOST
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

async function getProductListService() {
    const host = process.env.REACT_APP_ORCHESTRATOR_SERVICE_HOST
    const url = host + "/api/get_product_list"

    var axiosResponse = await axios({
        mode: 'no-cors',
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json"
        },
        method: 'POST',
        url: url,
        data: {
            "enableOnly": true
        }
    })

    console.log("Client Response",axiosResponse.data)
    return axiosResponse.data.map(product=>({
        id: product["id"],
        title: product["title"],
        description: product["stock"],
        price: product["price"],
        imageURL: "https://picsum.photos/200"
    }))

}

async function getProductService(productId) {
    const host = process.env.REACT_APP_ORCHESTRATOR_SERVICE_HOST
    const url = host + "/api/get_product"

    var axiosResponse = await axios({
        mode: 'no-cors',
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json"
        },
        method: 'POST',
        url: url,
        data: {
            "ID": productId
        }
    })

    console.log("Client Response Get product ", productId,axiosResponse.data)
    const product = axiosResponse.data
    return {
        id: product["id"],
        title: product["title"],
        description: product["stock"],
        price: product["price"],
        imageURL: "https://picsum.photos/200"
    }

}

export {upsertCartService, getCartService, placeOrderService, getProductListService, getProductService}