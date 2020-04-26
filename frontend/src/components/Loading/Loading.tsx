import React, { Component } from 'react';
import './Loading.css';

interface Props {
}

interface State {
}

class Loading extends Component<Props, State> {
  render() {
    return (
      <div className="logo-container">
          <div className="la-ball-atom la-2x">
            <div></div>
            <div></div>
            <div></div>
            <div></div>
          </div>
      </div>
    );
  }
}

export default Loading;
