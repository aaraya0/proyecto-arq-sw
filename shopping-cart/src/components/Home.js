

import React from "react";
import './Home.css';


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
		if (!DataisLoaded) return <div>
			<h1> Please wait... </h1> </div> ;

		return (
		
	

		<div className = "App">
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
		

				
				
			
	);
}

}

export default Home;
