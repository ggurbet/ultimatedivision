/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import { MarketPlaceFilterField }
    from '../MarketPlaceFilterField/MarketPlaceFilterField';
import { MarketPlaceCardsGroup }
    from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import { UltimateDivisionPaginator }
    from '../../UltimateDivisionPaginator/UltimateDivisionPaginator';
import { MarketPlaceFootballerCard }
    from '../MarketPlaceCardsGroup/MarketPlaceFootballerCard/MarketPlaceFootballerCard';

import './MarketPlace.scss';
import { MyCard } from '../MyCard/MyCard';
import { RouteConfig } from '../../../routes';


export const MarketPlace = ({ ...children }) => {
    const cards = useSelector((state: RootState) => state.cardReducer);

    //TODO: Route with SubRoutes
    let Component = MarketPlaceFootballerCard;
    let title = "MARKETPLACE";
    if (children.path === RouteConfig.MyCards.path) {
        Component = MyCard;
        title = "MY CARDS";
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
            <UltimateDivisionPaginator
                itemCount={cards.length} />
        </section>
    );
};
