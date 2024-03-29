import React from 'react';
import './App.css';
import NavBar from './components/NavBar';
import Home from './components/Home';
import Login from './components/Login';
import Cart from './components/Cart';
import Account from './components/Account';

import Checkout from './components/Checkout';


import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
function App() {
  return (
    <>
    <Router>
    
   
    <Routes>
    <Route exact path="/" element={<Login/>}/>
    </Routes>
    
    <NavBar/>
    <Routes>
    
    <Route exact path="/home" element={<Home/>}/>
    <Route exact path="/cart" element={<Cart/>}/>
    <Route exact path="/checkout" element={<Checkout/>}/>
    <Route exact path="/account" element={<Account/>}/>
    </Routes>
    </Router>
      
  
    </>
  );
}

export default App;
