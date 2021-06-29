/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';

import { MarketPlaceNavbar } from '../MarketPlaceNavbar/MarketPlaceNavbar';
import { MarketPlaceFilterField } from '../MarketPlaceFilterField/MarketPlaceFilterField';
import { MarketPlaceCardsGroup } from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import './MarketPlace.scss';

export const MarketPlace = () => {
    const cards = useSelector(state => state.footballerCard);

    return (
        <section className="marketplace">
            <MarketPlaceNavbar />
            <MarketPlaceFilterField />
            <MarketPlaceCardsGroup
                cards={cards} />
        </section>
    );
};
