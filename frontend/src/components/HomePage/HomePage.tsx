import React, { Component } from 'react';
import './HomePage.css';

interface Props {
}

interface State {
}

class HomePage extends Component<Props, State> {

  constructor(props: any) {
    super(props)
    this.state = {recipe: {name: "default"}}
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          Welcome to the home page!
        </header>
      </div>
    );
  }
}

export default HomePage;
