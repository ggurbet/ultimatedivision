/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import { MarketPlaceNavbar } from '../MarketPlaceNavbar/MarketPlaceNavbar';
import { MarketPlaceFilterField }
    from '../MarketPlaceFilterField/MarketPlaceFilterField';
import { MarketPlaceCardsGroup }
    from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import { UltimateDivisionPaginator }
    from '../../UltimateDivisionPaginator/UltimateDivisionPaginator';

import './MarketPlace.scss';

export const MarketPlace = () => {
    const cards = useSelector((state: RootState) => state.cardReducer);

    return (
        <section className="marketplace">
            <MarketPlaceNavbar />
            <MarketPlaceFilterField />
            <MarketPlaceCardsGroup
                cards={cards} />
            <UltimateDivisionPaginator
                itemCount={cards.length} />
        </section>
    );
};
