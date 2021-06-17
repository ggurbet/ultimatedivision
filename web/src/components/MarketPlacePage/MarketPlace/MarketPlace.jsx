import React from 'react';

import { MarketPlaceNavbar } from '../MarketPlaceNavbar/MarketPlaceNavbar';
import { MarketPlaceFilterField } from '../MarketPlaceFilterField/MarketPlaceFilterField';
import './MarketPlace.scss';


export const MarketPlace = () => {
    return (
        <section className="marketplace">
            <MarketPlaceNavbar />
            <MarketPlaceFilterField />
        </section>
    );
};
