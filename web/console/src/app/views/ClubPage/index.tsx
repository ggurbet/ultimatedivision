// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import { RootState } from '@/app/store';
import { listOfCards } from '@/app/store/actions/cards';

import './index.scss';

const Club: React.FC = () => {
    const { page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);

    return (
        <section className="club">
            <h1 className="club__title">
                MY CARDS
            </h1>
            <FilterField />
            <ClubCardsArea />
            <Paginator
                getCardsOnPage={listOfCards}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default Club;
