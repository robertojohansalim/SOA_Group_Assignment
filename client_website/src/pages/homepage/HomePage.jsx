import { Container, Row } from "react-bootstrap"
import ProductPost from "../../components/ProductPost/ProductPost"
export default function HomePage(){

    // TODO: Fetch From Product Service
    var productList = [
        {
            "id":"product-id-1",
            "title":"Lorem ipsum 1",
            "description":"Lorem Ipsum Dolor Si Amet",
            "imageURL":"https://picsum.photos/200"
        },
        {
            "id":"product-id-2",
            "title":"Lorem ipsum 2",
            "description":"Lorem Ipsum Dolor Si Amet",
            "imageURL":"https://picsum.photos/200"
        },
        {
            "id":"product-id-3",
            "title":"Lorem ipsum 3",
            "description":"Lorem Ipsum Dolor Si Amet",
            "imageURL":"https://picsum.photos/200"
        }
    ]

    function addToCart(productId){
        console.log(productId)
        // TODO: Add To Cart
        
    }

    return (
        <Container>
            <Row>
                {productList.map((product, idx)=>(
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
    )
}