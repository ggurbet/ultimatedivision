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
    createLotsQueryParameters,
    getCurrentLotsQueryParameters,
    listOfLots,
} from '@/app/store/actions/marketplace';
import { Card, CardsQueryParametersField } from '@/card';
import { Lot } from '@/marketplace';

import './index.scss';

const DEFAULT_VALUE = 1;
const MarketPlace: React.FC = () => {
    const dispatch = useDispatch();
    const { lots, page } = useSelector(
        (state: RootState) => state.marketplaceReducer.marketplacePage
    );

    /** TODO: delete this after setting real data. */
    const MOCK_CARD = new Card({
        id: 'b00c00c8-4311-4461-a245-751089516853',
        playerName: 'Miquel Deloach',
        quality: 'silver',
        pictureType: 0,
        height: 176.62,
        weight: 67.5,
        skinColor: 0,
        hairStyle: 0,
        hairColor: 0,
        accessories: [],
        dominantFoot: 'right',
        isTattoos: false,
        status: 0,
        type: 'won',
        userId: 'c4b97f28-d314-4b60-a2dd-26a9f73ce66e',
        tactics: 49,
        positioning: 41,
        composure: 56,
        aggression: 52,
        vision: 58,
        awareness: 54,
        crosses: 42,
        physique: 7,
        acceleration: 11,
        runningSpeed: 1,
        reactionSpeed: 12,
        agility: 1,
        stamina: 16,
        strength: 8,
        jumping: 7,
        balance: 13,
        technique: 51,
        dribbling: 59,
        ballControl: 59,
        weakFoot: 47,
        skillMoves: 41,
        finesse: 57,
        curve: 44,
        volleys: 45,
        shortPassing: 49,
        longPassing: 45,
        forwardPass: 60,
        offense: 0,
        finishingAbility: 12,
        shotPower: 10,
        accuracy: 2,
        distance: 1,
        penalty: 7,
        freeKicks: 1,
        corners: 3,
        headingAccuracy: 1,
        defence: 57,
        offsideTrap: 58,
        sliding: 62,
        tackles: 64,
        ballFocus: 61,
        interceptions: 57,
        vigilance: 58,
        goalkeeping: 52,
        reflexes: 50,
        diving: 49,
        handling: 42,
        sweeping: 47,
        throwing: 56,
    });

    const MOCK_LOT = {
        id: '23056596-a25b-4580-a719-dd7ac13b79bb',
        itemId: '27056596-a25b-4580-a719-dd7ac13b79cc',
        type: '',
        userId: '27056596-a25b-4580-a719-dd7ac13b79bb',
        shopperId: '27056596-a25b-8880-a719-dd7ac13b79bb',
        status: '',
        startPrice: DEFAULT_VALUE,
        maxPrice: DEFAULT_VALUE,
        currentPrice: DEFAULT_VALUE,
        startTime: '',
        endTime: '',
        period: DEFAULT_VALUE,
        card: MOCK_CARD,
    };

    const MOCK_LOTS = [new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT), new Lot(MOCK_LOT)];

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    const lotsQueryParameters = getCurrentLotsQueryParameters();

    /** Submits search by lots query parameters. */
    const submitSearch = async(
        queryParameters: CardsQueryParametersField[]
    ) => {
        createLotsQueryParameters(queryParameters);
        await dispatch(listOfLots(DEFAULT_PAGE_INDEX));
    };

    return (
        <section className="marketplace">
            <FilterField>
                <FilterByVersion
                    cardsQueryParameters={lotsQueryParameters}
                    submitSearch={submitSearch}
                />
                <FilterByStats
                    cardsQueryParameters={lotsQueryParameters}
                    submitSearch={submitSearch}
                />
                <FilterByPrice />
                <FilterByStatus />
            </FilterField>
            <MarketPlaceCardsGroup lots={MOCK_LOTS} />
            <Paginator
                getCardsOnPage={listOfLots}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default MarketPlace;
