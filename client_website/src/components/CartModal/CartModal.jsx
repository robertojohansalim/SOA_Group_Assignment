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
                Cart ID:{props.cart.ID}
                {props.cart.lineItems.length ===0 ? <Modal.Body>"Cart Is Empty"</Modal.Body>:null}
                {props.cart.lineItems.map((cartItem, idx) => {
                    // TODO STYLING: pake Bootstrap
                    return <Modal.Body key={idx}>
                        <Row>
                            <Col sm={2} >
                                <Button
                                    variant="outline-danger"
                                    onClick={() => props.removeFromCart(idx)}
                                >
                                    <RemoveIcon />
                                </Button>
                            </Col>
                            <Col sm={10}>
                                <p>{cartItem.title}</p>
                            </Col>
                        </Row>
                    </Modal.Body>
                })}
                <Modal.Footer>
                <p>
                    Choose Payment Method

                </p>
                    <Button 
                    variant={"outline-primary " + (props.cart.paymentMethod === "BANK_TRANSFER" ? "active" : "")} 
                    onClick={() => {
                            props.addPaymentMethod("BANK_TRANSFER")
                            }
                        }>
                        BANK TRANSFER
                    </Button>
                    <Button variant={"outline-primary " + (props.cart.paymentMethod === "BCA_VA" ? "active" : "")}
                     onClick={() => {
                            props.addPaymentMethod("BCA_VA")
                            }
                        }>
                        BCA
                    </Button>
                </Modal.Footer>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => props.onHide()}>
                        Close
                    </Button>
                    <Button variant="primary"  onClick={() => {
                            props.onHide()
                            props.checkout(props.cart.ID)
                            }
                        }>
                        Checkout
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}


export default CartModal;