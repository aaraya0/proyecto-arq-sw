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
<div>
<div>{order.order_id}</div>
<div>{order.total_amount}</div>
<div>{order.date}</div>
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
<div>
<div>{ad.street}</div>
<div>{ad.number}</div>
<div>{ad.city}</div>
<div>{ad.country}</div>
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
                setUser(res.name);
            } catch (err) {
                console.log(err);
            }
        }
        fetchData();
    }, []);
    return <div>{user}</div>
}


function Account (){
	if(cookies.get("user_id")===undefined){
		alert("Debe iniciar sesion antes de continuar")
		gopath("/")
	}



 return (

<div>
<div>Usuario</div>

<div>{ShowUser()}</div>
<div>{GetOrders()}</div>
<div>{GetAddresses()}</div>


</div>

 )

}
export default Account;