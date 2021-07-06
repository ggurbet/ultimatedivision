import React from 'react';
import { Route, Switch } from 'react-router-dom';

import { MarketPlace }
    from './components/MarketPlacePage/MarketPlace/MarketPlace';

import { FootballerCard }
    from './components/FootballerCardPage/FootballerCard/FootballerCard';

import { FootballField }
    from './components/FootballFieldPage/FootballField/FootballField';

import './App.scss';

export function App() {
    return (
        <>
            <Switch>
                <Route exact path="/ud/marketplace/">
                    <MarketPlace />
                </Route>
                <Route exact path="/ud/marketplace/card">
                    <FootballerCard/>
                </Route>
                <Route exact path="/ud/footballField">
                    <FootballField />
                </Route>
            </Switch>
        </>
    );
}

export default App;
