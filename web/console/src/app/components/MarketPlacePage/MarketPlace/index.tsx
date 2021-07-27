/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useSelector } from 'react-redux';

import { MarketPlaceCardsGroup } from '@components/MarketPlacePage/MarketPlaceCardsGroup';
import { MarketPlaceFilterField } from '@components/MarketPlacePage/MarketPlaceFilterField';
import { MarketPlaceFootballerCard } from '@components/MarketPlacePage/MarketPlaceCardsGroup/MarketPlaceFootballerCard';
import { MyCard } from '@components/MarketPlacePage/MyCard';
import { Paginator } from '@components/Paginator';

import { RouteConfig } from '@/app/routes';
import { RootState } from '@/app/store';

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
