// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { MarketPlaceCardsGroup } from '@components/MarketPlacePage/MarketPlaceCardsGroup';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import { RootState } from '@/app/store';

import './index.scss';

const MarketPlace = () => {
    const cards = useSelector((state: RootState) => state.cardReducer);

    return (
        <section className="marketplace">
            <FilterField
                title="MARKETPLACE"
            />
            <MarketPlaceCardsGroup
                cards={cards}
            />
            <Paginator
                itemCount={cards.length} />
        </section>
    );
};

export default MarketPlace;
