import React from "react"
import { Col, Card, Button } from "react-bootstrap"

// React Bootstrap
// https://react-bootstrap.netlify.app/components/navbar/#navbars
export default function ProductPost(props) {

    function cardOnClickHandler() {
        return
    }

    return (
        <Col xs={4} lg={3} xl={2}>
            <Card style={{height:"100%"}}
                onClick={cardOnClickHandler}
            >
                <Card.Img variant="top" src={props.imageURL} />
                <Card.Body>
                    <Card.Title>{props.title}</Card.Title>
                    <Card.Text>
                        {props.description}
                    </Card.Text>
                </Card.Body>
                    <Button onClick={()=>props.addToCart(props.id)} style={{"borderRadius": "0 !important"}}>Add To Cart</Button>
            </Card>
        </Col>
    )
}

