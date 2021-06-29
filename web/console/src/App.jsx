import React from 'react';
import { Route, Switch } from 'react-router-dom';

import { MarketPlace } from './components/MarketPlacePage/MarketPlace/MarketPlace';

import './App.scss';


export function App() {
    return (
        <>
            <Switch>
                <Route exact path="/ud/marketplace/">
                    <MarketPlace />
                </Route>
            </Switch>
        </>
    );
};

export default App;
