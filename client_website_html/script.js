var svc_product = 'http://svr-product.user.cloudjkt01.com';
var svc_cart = 'http://77a4-118-99-110-227.ngrok.io/api';


function showAllProduct() {
    $('#product-list').html('');

    $.ajax({
        url: svc_product + '/getproductlist',
        type: 'get',
        dataType: 'json',
        data: {

        },
        success: function (productList) {
            console.log(productList);
            if (productList) {
                let products = productList;

                $.each(products, function (i, data) {
                    // title = data.title.substr(0,37)
                    $('#product-list').append(`
                        <a id="card-thumb" class="product-card" href="product_page.html?${data.id}">
                            <div class="thumb flex-col p-1">
                                <div class="thumb-img">
                                    <img src="`+ data.product_image + `" alt="">
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

async function addToCart(id) {
    var productDetail = '';
    $.ajax({
        url: svc_product + '/productdetail',
        type: 'get',
        dataType: 'json',
        data: {
            'id': id
        },
        success: function (data) {
            if (data) {
                productDetail = data
                $.ajax({
                    url: svc_cart + '/upsert_cart',
                    type: 'post',
                    dataType: 'json',
                    data: JSON.stringify({
                        "ID": "cart-id",
                        "lineItems": [
                            {
                                "product_id": id,
                                "title": productDetail.product_name || "",
                                "description": productDetail.product_desc || "",
                                "quantity": 1 || 0,
                                "price": productDetail.product_price || 0
                            }
                        ],
                        "paymentMethod": "",
                        "totalPrice": productDetail.product_price * 1
                    }),
                    success: function (data) {
                        console.log('add to cart ok');
                        console.log(data)
                    }
                });
            } else {
                console.log(data)
                alert('error');
            }
        }
    }).then();
    
}

function showProductDetails() {
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
            if (data) {
                $('#product-details').append(`
                    <div class="details-container flex-row">
                        <div class="product-img">
                            <img src="${data.product_image}" alt="">
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
        }
    });
}

function showCart() {
    $('#cart-item-list').html('');
    var totalOrder = 0;
    var shippingCost = 2;
    totalOrder = totalOrder + shippingCost
    console.log(totalOrder);

    $.ajax({
        url: 'https://fakestoreapi.com/carts/1',
        type: 'get',
        dataType: 'json',
        data: {

        },
        success: function (itemList) {
            if (itemList) {
                let products = itemList.products;

                $.each(products, function (i, data) {
                    let qty = data.quantity;

                    $.ajax({
                        url: 'https://fakestoreapi.com/products/' + data.productId,
                        type: 'get',
                        dataType: 'json',
                        data: {

                        },
                        success: function (data) {

                            let product = data;
                            var totalPrice = product.price * qty;
                            // console.log(totalOrder);
                            totalOrder = totalOrder + totalPrice

                            $('#total-order').html('');
                            $('#total-order').append(`$` + (totalOrder - shippingCost));

                            $('#total-amount').html('');
                            $('#total-amount').append(`<strong>$` + totalOrder + `</strong>`);

                            $('#cart-item-list').append(`
                                <div class="cart-item flex-row">
                                    <a href="product_page.html?${data.id}" class="flex-row">
                                        <div class="cart-item-details flex-row">
                                            <div class="cart-item-img">
                                                <img src="`+ product.image + `" alt="">
                                            </div>
                                            <div class="cart-item-title">`+ product.title + `</div>
                                        </div>
                                    </a>
                                    <div class="cart-item-qty">
                                        <form type="submit">
                                            <input type="number" value="`+ qty + `"/>
                                        </form>
                                    </div>
                                    <div class="cart-item-price">$`+ product.price + `</div>
                                    <div class="cart-item-total">$`+ totalPrice + `</div>
                                </div>
                            `);
                        }

                    });
                });
            } else {
                $('#cart-item-list').html(`
                    <div>
                        <h1>No Data :(</h1>
                    </div>
                `)
            }


        }
    });
}
