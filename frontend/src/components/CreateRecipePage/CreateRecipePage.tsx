import React, { Component } from 'react';
import './CreateRecipePage.css';
import axios from "axios";
import { Redirect } from 'react-router-dom';

interface IProps {
}

interface IState {
    name: string
    ingredients: string
    method: string
    redirect: string
}

class CreateRecipePage extends Component<IProps, IState> {

    constructor(props: any) {
        super(props);
        this.state = {
            name: "default",
            ingredients: "default",
            method: "default",
            redirect: ""
        }
    }

    addRecipe(event: React.FormEvent) {
        axios.post("http://localhost:4000/recipe/add", JSON.stringify({
            name: this.state.name,
            ingredients: this.state.ingredients,
            method: this.state.method
        })).then(response => {
            console.log(response.data);
            this.setState({redirect: "/recipe/" + response.data.id})
        }).catch(error => {
            console.log(error);
        })
    }

  render() {
      if (this.state.redirect) {
          return <Redirect to={this.state.redirect} />
      }
    return (
      <div className="App">
        <header className="App-header">
            <form>
                <label>
                    Title:
                    <input type="text" name="name" onChange={e => this.setState({name: e.target.value})} />
                </label>
                <label>
                    Ingredients:
                    <input type="text" name="ingredients" onChange={e => this.setState({ingredients: e.target.value})} />
                </label>
                <label>
                    Method:
                    <input type="text" name="method" onChange={e => this.setState({method: e.target.value})} />
                </label>
                <input type="button" value="Submit" onClick={this.addRecipe.bind(this)}/>
            </form>
        </header>
      </div>
    );
  }
}

export default CreateRecipePage;
