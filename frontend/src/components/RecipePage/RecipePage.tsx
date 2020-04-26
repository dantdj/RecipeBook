import React, { Component } from 'react';
import axios from "axios";
import "./RecipePage.css"
import { RouteComponentProps } from 'react-router-dom';
import Loading from '../Loading/Loading';

interface HomeRouterProps {
    id: string
}

interface Props extends RouteComponentProps<HomeRouterProps> {
}

interface State {
  recipe: Recipe
  loaded: boolean
}

interface Recipe {
  name?: string
  ingredients?: string
  method?: string
  created?: string
}

class RecipePage extends Component<Props, State> {

  constructor(props: any) {
    super(props)
    this.state = {recipe: {}, loaded: false}
  }

  componentDidMount() {
    axios.get("http://localhost:4000/recipe?id=" + this.props.match.params.id, 
    ).then(response => {
      this.setState({recipe: response.data, loaded: true})
    }).catch(error => {
      console.log("Failed to call API: " + error)
    })
  }

  render() {
    if (!this.state.loaded) {
      return <Loading />
    }
    return (
      <div className="App">
        <header className="App-header">
          Name: {this.state.recipe.name} <br />
          Ingredients: {this.state.recipe.ingredients} <br />
          Method: {this.state.recipe.method} <br />
          Created: {this.state.recipe.created} <br />
        </header>
      </div>
    );
  }
}

export default RecipePage;
