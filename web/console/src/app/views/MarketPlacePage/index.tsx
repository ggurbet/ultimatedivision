// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '@/app/store';
import { filteredLots } from '@/app/store/actions/marketplace';

import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const MarketPlace: React.FC = () => {
    const { lots } = useSelector((state: RootState) => state.marketplaceReducer.marketplace);

    return (
        <section className="marketplace">
            <FilterField
                title="MARKETPLACE"
                thunk={filteredLots}
            />
            <MarketPlaceCardsGroup
                lots={lots}
            />
            <Paginator
                itemCount={lots.length}
            />
        </section>
    );
};

export default MarketPlace;
