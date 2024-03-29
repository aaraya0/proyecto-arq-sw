import React, { useState } from "react";

import "./Login.css";

async function getUserByID(id) {
  return fetch('http://127.0.0.1:8090/user/' + id)
  .then(data => data.json())

}

async function setCookie (cname, cvalue){

  document.cookie = cname+"="+cvalue+";"+"path=/";
}

function gopath(path){
  window.location = window.location.origin + path
}

function Login (){

const [errorMessages, setErrorMessages] = useState({});
const [isSubmitted, setIsSubmitted] = useState(false);
const errors = {
uname: "invalid username",
pass: "invalid password" };

const handleSubmit = async event => {
//Prevent page reload
event.preventDefault();

var { uname, pass } = document.forms[0];

const user = await getUserByID(uname.value);

// Compare user info
  if (user) {
  if (user.password != pass.value) {
  // Invalid password
  setErrorMessages({ name: "pass", message: errors.pass });
  } else {
  setIsSubmitted(true);

  setCookie("name", user.name)
  setCookie ("user_id", user.id)
  setCookie("user", uname.value)

}
} else {

  // Username not found
  setErrorMessages({ name: "uname", message: errors.uname });

}
};


const renderErrorMessage = (name) =>
name === errorMessages.name && (
<div className="error">{errorMessages.message}</div>
);


const renderForm = (
    <div className="form2">
    <form onSubmit={handleSubmit}>
    <div className="input-container">
    <label>Usuario </label>
    <input type="text" name="uname" required />
    {renderErrorMessage("uname")}
    </div>
    <div className="input-container">
    <label>Contraseña </label>
    <input type="password" name="pass" required />
    {renderErrorMessage("pass")}
    </div>
    <div className="button-container">
    <input type="submit" value="Iniciar Sesión"/>
    </div>
    </form>
    </div>
);

      return (
      <div className="app">
      <div className="login-form">

      {isSubmitted ? gopath('/home')
      : renderForm}

      </div>
      </div>
      );
}
export default Login;
