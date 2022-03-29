import React, { useState } from "react"
import { Navbar as NavbarBS, Nav, Container } from "react-bootstrap"
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import "./Navbar.css"
import logo from '../../assets/Logo.png'


// React Bootstrap
// https://react-bootstrap.netlify.app/components/navbar/#navbars
function Navbar(props) {

  const [cartItemCount, setCartItemCount] = useState(0)

  return (
    <div>
      <NavbarBS collapseOnSelect expand="lg" bg="light" className="mb-4">
        <Container>
          <NavbarBS.Brand href="#home">
            <img alt="Logo" src={logo} style={{ "height": "3rem" }} />
            SOAmart
          </NavbarBS.Brand>
          <Nav>
            <Nav.Link href="" onClick={props.showCartModalHanlder}>
              <ShoppingCartIcon />
            </Nav.Link>
          </Nav>
{/*           
          <NavbarBS.Toggle aria-controls="responsive-navbar-nav" />
          <NavbarBS.Collapse id="responsive-navbar-nav" className="flex-row-reverse">
            <Nav>
              <Nav.Link href="" onClick={props.showCartModalHanlder}>
                <ShoppingCartIcon />
              </Nav.Link>
              <Nav.Link eventKey={2} href="#memes">
                Login
              </Nav.Link>
              <Nav.Link eventKey={2} href="#memes">
                Sign up
              </Nav.Link>
            </Nav>
          </NavbarBS.Collapse> */}
        </Container>
      </NavbarBS>


    </div>
  )
}

export default Navbar;