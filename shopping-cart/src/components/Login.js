
import './Login.css'

import React, { useState } from "react";


async function getUserByID(id) {
 return fetch('http://127.0.0.1:8090/user/' + id, {
   method: 'GET',
   headers: {
     'Content-Type': 'application/json'
   }
 })
   .then(data => data.json())
}


function Login() {

  const [userData, setUserData] = useState({});
  const [isUser, setIsUSer] = useState(false);
  
  const handleSubmit = async event => {
    //Prevent page reload
    event.preventDefault();

    // Obtenemos Textos
    var {uid, upass}  = document.forms[0];

    alert(upass.value);
    // Find user login info
    const user = await getUserByID(uid.value);
    
    setUserData(user);
    setIsUSer(true);
    
  };

  
  const renderForm = (
      <div className="back">
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Usuario </label>
          <input type="text" name="uid" required />
          <label>Contraseña </label>
          <input type="password" name="upass" required />
        </div>
        <div className="button-container">
          <input type="submit" value="Log In" />
        </div>
      </form>
    </div>
    </div>
  );

  return (
    <div className="app">
      <div className="login-form">
        <div className="title">Login{userData.name}</div>
          {isUser? <div>Usuario: {userData.name} </div> : renderForm}
      </div>
    </div>
  );
}

export default Login;