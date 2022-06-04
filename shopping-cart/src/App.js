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
    <NavBar/>
   
    <Routes>
    <Route exact path="/" element={<Home/>}/>
    </Routes>
    <Routes>
    <Route exact path="/login" element={<Login/>}/>
    </Routes>
    </Router>
      
  
    </>
  );
}

export default App;
