import { Route, Routes } from "react-router-dom"
import Navbar from "./components/Navbar/Navbar"
import CartModal from "./components/CartModal/CartModal"
import { HomePage, } from "./pages"
// import './App.css';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<HomePage />} />
      </Routes>
    </div>
  );
}

export default App;
