import React, { useState } from "react";
import {Link} from 'react-router-dom'

import "./Login.css";
async function getUserByID(id) {
  return fetch('http://127.0.0.1:8090/user/' + id)
    .then(data => data.json())
    
 }

function Login (){
   

  const [errorMessages, setErrorMessages] = useState({});
  const [isSubmitted, setIsSubmitted] = useState(false);
  //const [userData, setUserData] = useState({});
  
  const errors = {
    uname: "invalid username",
    pass: "invalid password"
  };

  const handleSubmit = async event => {
    //Prevent page reload
    event.preventDefault();

    var { uname, pass } = document.forms[0];

    // Find user login info
    const user = await getUserByID(uname.value);
    
    //setUserData(user);
    
    alert(user.password);

    // Compare user info
    if (user) {
      if (user.password != pass.value) {
        // Invalid password
        setErrorMessages({ name: "pass", message: errors.pass });
      } else {
        setIsSubmitted(true);
      }
    } else {
      // Username not found
      setErrorMessages({ name: "uname", message: errors.uname });
    }
  };

  // Generate JSX code for error message
  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

  // JSX code for login form
  const renderForm = (
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Username </label>
          <input type="text" name="uname" required />
          {renderErrorMessage("uname")}
        </div>
        <div className="input-container">
          <label>Password </label>
          <input type="password" name="pass" required />
          {renderErrorMessage("pass")}
        </div>
        <div className="button-container">
          <input type="submit" />
        </div>
      </form>
    </div>
  );

  return (
    <div className="app">
      <div className="login-form">
       
        {isSubmitted ? <Link to="/home" className="home">Home</Link> 
        : renderForm}
        
      </div>
    
     
    </div>
  );
}

export default Login;
