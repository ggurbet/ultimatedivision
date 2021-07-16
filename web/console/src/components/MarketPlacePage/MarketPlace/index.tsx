/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import { MarketPlaceFilterField } from '../MarketPlaceFilterField';
import { MarketPlaceCardsGroup } from '../MarketPlaceCardsGroup';
import { Paginator } from '../../Paginator';
import { MarketPlaceFootballerCard } from '../MarketPlaceCardsGroup/MarketPlaceFootballerCard';
import { MyCard } from '../MyCard';

import { RouteConfig } from '../../../routes';

import './index.scss';

export const MarketPlace = ({ ...children }) => {
    const cards = useSelector((state: RootState) => state.cardReducer);

    let Component = MarketPlaceFootballerCard;
    let title = 'MARKETPLACE';
    if (children.path === RouteConfig.MyCards.path) {
        Component = MyCard;
        title = 'MY CARDS';
    };

    return (
        <section className="marketplace">
            <MarketPlaceFilterField
                title={title}
            />
            <MarketPlaceCardsGroup
                cards={cards}
                Component={Component}
            />
            <Paginator
                itemCount={cards.length} />
        </section>
    );
};
