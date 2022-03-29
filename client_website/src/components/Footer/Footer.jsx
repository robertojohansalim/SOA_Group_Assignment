import React from "react";
import { Container } from "react-bootstrap";

 function FooterPage  ()  {
  return (
    <footer color="blue" className="mt-5">
      <div className="text-center py-3">
        <Container fluid>
          &copy; {new Date().getFullYear()} Copyright: <a href="#"> SOA Group Project 5 </a>
        </Container>
      </div>
    </footer>
  );
}

export default FooterPage;