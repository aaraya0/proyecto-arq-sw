import React, {useEffect, useState}  from "react";

import Cookies from 'universal-cookie';
import "./Account.css";
const cookies = new Cookies();
async function getUserByID(username) {
    return fetch('http://127.0.0.1:8090/user/' + username)
    .then(data => data.json())
  
  }
  async function getAddressesByUID(id) {
    return fetch('http://127.0.0.1:8090/address/' + id)
    .then(data => data.json())
  
  }

  async function getOrdersByUID(id){
    return fetch('http://127.0.0.1:8090/order/' + id)
    .then(data => data.json())
  }
  function gopath(path){
    window.location = window.location.origin + path
    }
async function getUser(){

    return await getUserByID(cookies.get("user"))

}





const GetOrders=()=>{
    const [ordenes,setOrdenes]=useState([]);
    useEffect( () => { async function getOrd(){
        let orders= await getOrdersByUID(cookies.get("user_id"))
    setOrdenes(orders)}
    getOrd();
    
},[]);



return (ordenes.map((order)=>
<div id="ordenes">
<div className="or"> ID de la orden: {order.order_id}</div>
<div className="or"> Total: ${order.total_amount}</div>
<div className="or"> Fecha: {order.date}</div>
<details>{order.orderDetail.map((detalle)=> <div><div>Libro: {detalle.name}</div><div>Cantidad: {detalle.quantity}</div></div>)}</details>
</div>
)
)}

const GetAddresses=()=>{
    const [address,setAddress]=useState([[]]);
    useEffect( () => { async function getAddr(){
        let addresses= await getAddressesByUID(cookies.get("user_id"))
    setAddress(addresses)}
    getAddr();

},[]);


return (address.map((ad)=>
<div id="direcciones">
<div className="direc">{ad.street}</div>
<div className="direc">{ad.number}, </div>
<div className="direc">{ad.city}, </div>
<div className="direc">{ad.country}, </div>
<div>{ad.cod_postal}</div>
</div>
)
)



}

const ShowUser = () => {
    const [user, setUser] = useState([]);
    
    useEffect( () => { 
        async function fetchData() {
            try {
                const res = await getUser()
                setUser(res);
            } catch (err) {
                console.log(err);
            }
        }
        fetchData();
    }, []);
    return (<div><div>{user.name}</div>
    <div id="userName">@{user.user_name}</div></div>
    )
}


function Account (){
	if(cookies.get("user_id")===undefined){
		alert("Debe iniciar sesion antes de continuar")
		gopath("/")
	}



 return (

<div className="usuario">

<div id="img"><img src="img/user.png" id="imguser" ></img></div>
<div id="user">{ShowUser()}</div>
<div id="direcc">Direcciones de envio registradas</div>
<div id="direcc2">{GetAddresses()}</div>

<div id="compras">Historial de compras</div>

<div id="compras2">{GetOrders()}</div>


</div>

 )

}
export default Account;