import React from 'react';

import './App.css';
import NavBar from './components/NavBar';
import Home from './components/Home';
import Login from './components/Login';
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
    </Routes>
    </Router>
      
  
    </>
  );
}

export default App;
