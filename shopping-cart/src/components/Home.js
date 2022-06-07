

import React from "react";
import './Home.css';
import image from "./img/libro1.jpg";
class Home extends React.Component {

	// Constructor
	constructor(props) {
		super(props);

		this.state = {
			items: [],
			DataisLoaded: false
		};
	}

	// ComponentDidMount is used to
	// execute the code
	componentDidMount() {
		fetch(
"http://localhost:8090/product")
			.then((res) => res.json())
			.then((json) => {
				this.setState({
					items: json,
					DataisLoaded: true
				});
			})
	}
	
	render() {
		const { DataisLoaded, items } = this.state;
	//	const {query, setQuery}=this.state;
		if (!DataisLoaded) return <div>
			<h1> Please wait... </h1> </div> ;


const producto= items.map((item) => (
	
	<ol key = { item.id } className="item">
		
	<div id="titulo">{ item.title}</div>
	<div id="autor"> { item.author }</div>	
	 <div id="precio">${ item.base_price }</div> 
	 <img src={image}/>
	 <div><button class="button-50" role="button">Agregar 🛒</button></div>
		</ol>
		
))

		return (
		<div className="product">{producto}</div>
		)

		/*<div className = "App">
			<h1 id="prod">Listado de productos</h1> {
				items.map((item) => (
				<ol key = { item.id } className="item">
				<div id="titulo">{ item.title}</div>
				<div id="autor"> { item.author }</div>	
				 <div id="precio">${ item.base_price }</div> 
					</ol>
				))
			}
		</div>
		);*/



      /*  <div>
      <input placeholder="Enter Title" onChange={event => setQuery(event.target.value)} />
	
		{
			items.filter(item => {
			  if (query === '') {
				return items;
			  } else if (item.title.toLowerCase().includes(query.toLowerCase())) {
				return items;
			  }
			}).map((item) => (
			  <div className="box" key={item.id}>
				<p>{item.title}</p>
				<p>{item.author}</p>
			  </div>
			))
		  }

</div>		
				
			
		);*/
	/*	<div>
		<h1 id="prod">Listado de productos</h1>
<Card style={{width:'18rem'}}>
			 {
				items.map((item) => (
					
				<div key = { item.id } className="item">
				
				<Card.Body>
				
				<Card.Title>{ item.title}</Card.Title>
				<Card.Subtitle>{ item.author }</Card.Subtitle> 
				<Card.Text>${ item.base_price }</Card.Text>
		
				
				
					
					</Card.Body>
					</div>
				
				))
			}
		</Card>
		</div>
		)*/
}

}

export default Home;
