var svc_product = 'http://svr-product.user.cloudjkt01.com';
// var svc_cart = 'http://77a4-118-99-110-227.ngrok.io/api';
// const svc_product = 'http://127.0.0.1:8000';
const svc_cart = 'http://127.0.0.1:5000/api';
const cart_id = "cart-id-1"

// Helper Function
function getCart(cart_id) {
    return $.ajax({
        url: svc_cart + `/get_cart/${cart_id}`,
        type: 'get',
        dataType: 'json',
        data: {}
    }
    )
}

function getProductDetail(product_id) {
    return $.ajax({
        url: svc_product + '/productdetail',
        type: 'get',
        dataType: 'json',
        data: {
            'id': product_id
        },
    })
}

function showAllProduct() {
    $('#product-list').html('');

    console.log("get All Product")
    $.ajax({
        url: svc_product + '/getproductlist',
        type: 'get',
        dataType: 'json',
        data: {

        },
        success: function (productList) {
            console.log("Get Success")
            console.log(productList);
            if (productList) {
                let products = productList;

                $.each(products, function (i, data) {
                    // title = data.title.substr(0,37)
                    $('#product-list').append(`
                        <a id="card-thumb" class="product-card" href="product_page.html?${data.id}">
                            <div class="thumb flex-col p-1">
                                <div class="thumb-img">
                                    <img src="`+ (data.product_image || "https://dummyimage.com/300") + `" alt="">
                                </div>
                                <div class="thumb-content flex-col">
                                    <p class="thumb-title mb-1">`+ data.product_name + `</p>
                                    <p class="thumb-price">`+ "$" + `${data.product_price}</p>
                                </div>
                            </div>
                        </a>
                    `);
                });

            } else {
                $('#product-list').html(`
                    <div>
                        <h1>No Data :(</h1>
                    </div>
                `)
            }
        }
    });
}

function addToCart(product_id, quantity = 1) {
    $.when(
        getProductDetail(product_id),
        getCart(cart_id)
    ).done(([productDetail], [{ cart: cart }]) => {
        var isProductExistsInCart = false
        const prevLineItems = cart?.lineItems?.map(lineItem => {
            var qty = lineItem.quantity
            if (lineItem.product_id == product_id) {
                isProductExistsInCart = true
                qty++;
            }

            return {
                "product_id": lineItem.product_id,
                "title": lineItem.title || "",
                "description": lineItem.description || "",
                "quantity": qty,
                "price": lineItem.price || 0
            }
        }) || []

        if (!isProductExistsInCart) {
            prevLineItems.push({
                "product_id": productDetail.id,
                "title": productDetail.product_name || "",
                "description": productDetail.product_desc || "",
                "quantity": 1 || 0,
                "price": productDetail.product_price || 0
            })
        }

        var requestData = {
            "ID": cart_id,
            "lineItems": prevLineItems,
        }
        console.log(requestData)
        return $.ajax({
            url: svc_cart + '/upsert_cart',
            type: 'post',
            dataType: 'json',
            data: JSON.stringify(requestData),
            success: function (data) {
                console.log('add to cart ok');
                console.log(data)
            }
        });
    })
    return
}

function showProductDetails() {
    console.log("Show Product Details")
    $('#product-details').html('');
    var productId = location.search.substring(1);

    $.ajax({
        url: svc_product + '/productdetail',
        type: 'get',
        dataType: 'json',
        data: {
            'id': productId
        },
        success: function (data) {
            console.log("SUCCESS Show Product Details")
            if (data) {
                $('#product-details').append(`
                    <div class="details-container flex-row">
                        <div class="product-img">
                            <img src="` + (data.product_image || "https://dummyimage.com/300") + `" alt="">
                        </div>
                        <div class="product-details flex-col">
                            <h1>${data.product_name}</h1><br>
                            <p>${data.product_desc}</p><br><br>
                            <div class="price-details flex-row">
                                <h2 id="prod-price">`+ "$" + `${data.product_price}</h2>
                                <button onclick="addToCart(${data.id})">
                                    ADD TO CART
                                </button>
                            </div>
                        </div>
                    </div>
                `);
            } else {
                $('#product-details').html(`
                    <div>
                        <h1>No Data :(</h1>
                    </div>
                `)
            }
            console.log("SUCCESS Show Product Details")
        }
    });
    console.log("END Show Product Details")
}

function showCart() {
    $('#cart-item-list').html('');
    var totalOrder = 0;
    var shippingCost = 2;
    totalOrder = totalOrder + shippingCost
    console.log(totalOrder);

    $.ajax({
        url: svc_cart + `/get_cart/${cart_id}`,
        type: 'get',
        dataType: 'json',
        data: {

        },
        success: function ({ cart }) {
            const products = cart?.lineItems || []
            $.each(products, function (i, product) {
                let qty = product.quantity;

                $.ajax({
                    url: svc_product + '/productdetail',
                    type: 'get',
                    dataType: 'json',
                    data: {
                        'id': product.product_id
                    },
                    success: function (productDetail) {
                        let product = productDetail;
                        console.log(product)
                        var totalPrice = product.product_price * qty;
                        console.log(totalPrice);
                        totalOrder = totalOrder + totalPrice

                        $('#total-order').html('');
                        $('#total-order').append(`$` + (totalOrder - shippingCost));

                        $('#total-amount').html('');
                        $('#total-amount').append(`<strong>$` + totalOrder + `</strong>`);

                        $('#cart-item-list').append(`
                            <div class="cart-item flex-row">
                                <a href="product_page.html?${product.id}" class="flex-row">
                                    <div class="cart-item-details flex-row">
                                        <div class="cart-item-img">
                                            <img src="`+ (product.product_image || "https://dummyimage.com/300") + `" alt="">
                                        </div>
                                        <div class="cart-item-title">`+ product.product_name + `</div>
                                    </div>
                                </a>
                                <div class="cart-item-qty">
                                    <form type="submit">
                                        <input type="number" value="`+ qty + `"/>
                                    </form>
                                </div>
                                <div class="cart-item-price">$`+ product.product_price + `</div>
                                <div class="cart-item-total">$`+ totalPrice + `</div>
                            </div>
                        `);
                    }

                });
            });
            return
        }
    });
}

function registerCheckoutButton() {
    $('#checkout-btn').click(function (event) {
        // Testing only, prevent refreshing page
        event.preventDefault();
        console.log("Checkout Cliecked")

        var method = $('#method :selected').text();
        console.log(method)

        getCart(cart_id).then(({ cart: cart })=>{
            console.log(cart)
            $.ajax({
                url: svc_cart + "/upsert_cart", // Add To Cart
                type: 'post',
                dataType: 'json',
                data: JSON.stringify({
                    ...cart,
                    "paymentMethod":method
                })
            })
        }
        )
        .then(_ =>
            $.ajax({
                url: svc_cart + "/place_order",
                type: 'post',
                dataType: 'json',
                data: JSON.stringify({
                    "ID": cart_id,
                    "action": "CHECKOUT"
                })
            })
        ).then(response2 => {
            window.location = '?paymentID=' + response2.payment_id;
        })
    });
}

function registerPayButton() {
    var urlParams = new URLSearchParams(window.location.search);
    //TODO: Dapetin ?paymentID di url param
    var paymentId = urlParams.get('paymentID');
    console.log("PAYMENTID:", paymentId);


    //TODO: masukin value kyk SC021 dari sini (pake value paymentID diatas)
    $('#id-order').html(`<strong>`+ (paymentId || "No Order Yet") +`</strong>`);


    $('#pay-btn').click(function (event) {
        // Testing only, prevent refreshing page
        event.preventDefault();
        // https://stackoverflow.com/questions/9870512/how-to-obtain-the-query-string-from-the-current-url-with-javascript  
        

        //TODO: Handle Payment
        console.log("Pay Cliecked")
        $.ajax({
            url: '',
            type: 'post',
            dataType: 'json',
            data: {

            },
            success: function (itemList) {
                alert("Payment Success!")
                window.location = '?paymentID=' + response2.payment_id;

                $('#ongoing-payment').html('There is no ongoing payment.');
            }
        });
    });
}


