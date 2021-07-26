/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useSelector } from 'react-redux';

import { MarketPlaceCardsGroup } from '@marketPlace/MarketPlaceCardsGroup';
import { MarketPlaceFilterField } from '@marketPlace/MarketPlaceFilterField';
import { MarketPlaceFootballerCard } from '@marketPlace/MarketPlaceCardsGroup/MarketPlaceFootballerCard';
import { MyCard } from '@marketPlace/MyCard';
import { Paginator } from '@paginator';

import { RouteConfig } from '@routes';
import { RootState } from '@store';

import './index.scss';

const MarketPlace = ({ ...children }) => {
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

export default MarketPlace;
