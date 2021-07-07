import React from 'react';
import { BrowserRouter } from 'react-router-dom';

import { UltimateDivisionNavbar }
    from './components/UltimateDivisionNavbar/UltimateDivisionNavbar';
import { Routes } from './routes/index'

import { FootballField }
    from './components/FootballFieldPage/FootballField/FootballField';

import './App.scss';


export function App() {
    return (
        <BrowserRouter>
            <UltimateDivisionNavbar />
            <Routes />
        </BrowserRouter>
    );
}

export default App;
