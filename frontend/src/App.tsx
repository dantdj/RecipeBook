import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from "axios";

interface Props {

}

interface State {
  recipe: Recipe
}

interface Recipe {
  name?: string
  ingredients?: string
  method?: string
  created?: string
}

class App extends Component<Props, State> {

  constructor(props: any) {
    super(props)
    this.state = {recipe: {name: "default"}}
  }

  componentDidMount() {
    axios.get("http://localhost:4000/recipe?id=3", 
    ).then(response => {
      console.log(response.data)
      this.setState({recipe: response.data})
    }).catch(error => {
      console.log("Failed to call API: " + error)
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          Name: {this.state.recipe.name} <br />
          Ingredients: {this.state.recipe.ingredients} <br />
          Method: {this.state.recipe.method} <br />
          Created: {this.state.recipe.created} <br />
        </header>
      </div>
    );
  }
}

export default App;
