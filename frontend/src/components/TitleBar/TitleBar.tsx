import React, { Component } from 'react';
import './TitleBar.css';
import { Link } from 'react-router-dom';

interface Props {
}

interface State {
}

class TitleBar extends Component<Props, State> {
  render() {
    return (
      <>
      <Link to="/" className="title">RecipeBook</Link>

      <Link to={"/recipe/3"}>Recipe 3</Link>

      <Link to={"/addrecipe"}>Add Recipe</Link>
      </>
    );
  }
}

export default TitleBar;
