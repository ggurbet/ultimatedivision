// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { FilterField } from '@components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';
import { Paginator } from '@components/common/Paginator';
import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';

import { RootState } from '@/app/store';
import {
    listOfLots,
    createLotsQueryParameters,
} from '@/app/store/actions/marketplace';
import { CardsQueryParametersField } from '@/card';

import './index.scss';
import { useState } from 'react';

const MarketPlace: React.FC = () => {
    const dispatch = useDispatch();
    const { lots, page } = useSelector(
        (state: RootState) => state.marketplaceReducer.marketplacePage
    );

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** Submits search by lots query parameters. */
    const submitSearch = async(
        queryParameters: CardsQueryParametersField[]
    ) => {
        createLotsQueryParameters(queryParameters);
        await dispatch(listOfLots(DEFAULT_PAGE_INDEX));
    };

    return (
        <section className="marketplace">
            <h1 className="marketplace__title">MARKETPLACE</h1>
            <FilterField>
                <FilterByVersion
                    submitSearch={submitSearch}
                />
                <FilterByStats
                    submitSearch={submitSearch}
                />
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
