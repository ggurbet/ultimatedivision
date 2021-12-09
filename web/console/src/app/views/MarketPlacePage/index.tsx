// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { FilterField } from '@components/common/FilterField';
import { FilterByPrice } from '@/app/components/common/FilterField/FilterByPrice';
import { FilterByStatus } from '@/app/components/common/FilterField/FilterByStatus';
import { Paginator } from '@components/common/Paginator';
import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';

import { RootState } from '@/app/store';
import { listOfLots } from '@/app/store/actions/marketplace';

import './index.scss';

const MarketPlace: React.FC = () => {
    const { lots, page } = useSelector((state: RootState) => state.marketplaceReducer.marketplacePage);

    return (
        <section className="marketplace">
            <h1 className="marketplace__title">
                MARKETPLACE
            </h1>
            <FilterField>
                <FilterByPrice />
                <FilterByStatus />
            </FilterField>
            <MarketPlaceCardsGroup lots={lots} />
            <Paginator
                getCardsOnPage={listOfLots}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default MarketPlace;
