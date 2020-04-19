import React, {Component} from 'react';
import {Route, HashRouter} from 'react-router-dom';
import TitleBar from '../TitleBar/TitleBar';
import HomePage from '../HomePage/HomePage'
import RecipePage from '../RecipePage/RecipePage';

class Main extends Component {
    render() {
        return (
            <HashRouter>
                <div>
                    <TitleBar />
                    <div className="content">
                        <Route exact path="/" component={HomePage} />
                        <Route path="/recipe/:id" component={RecipePage} />} />
                    </div>
                </div>
            </HashRouter>
        );
    }
}

export default Main;