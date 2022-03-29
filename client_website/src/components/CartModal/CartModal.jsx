import React from "react"
import { Modal, Button, Row, Col } from "react-bootstrap"
import RemoveIcon from '@mui/icons-material/Remove';

function CartModal(props) {
    return (
        <>
            <Modal show={props.show} onHide={() => props.onHide()}>
                <Modal.Header closeButton>
                    <Modal.Title>Cart</Modal.Title>
                </Modal.Header>
                {props.cart.length ===0 ? <Modal.Body>"Cart Is Empty"</Modal.Body>:null}
                {props.cart.map((cartItem, idx) => {
                    // TODO STYLING: pake Bootstrap
                    return <Modal.Body>
                        <Row>
                            <Col sm={2} >
                                <Button
                                    variant="outline-danger"
                                    key={idx}
                                    onClick={() => props.removeFromCart(idx)}
                                >
                                    <RemoveIcon />
                                </Button>
                            </Col>
                            <Col sm={10}>
                                <p key={idx}>{cartItem}    </p>
                            </Col>
                        </Row>
                    </Modal.Body>
                })}
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => props.onHide()}>
                        Close
                    </Button>
                    <Button variant="primary" onClick={() => props.onHide()}>
                        Checkout
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}


export default CartModal;