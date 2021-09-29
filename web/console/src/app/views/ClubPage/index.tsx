// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { filteredCards, listOfCards } from '@/app/store/actions/cards';
import { RootState } from '@/app/store';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {
    const { page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);

    return (
        <section className="club">
            <FilterField
                title="My cards"
                thunk={filteredCards}
            />
            <ClubCardsArea />
            <Paginator
                getCardsOnPage={listOfCards}
                pagesCount={page.pageCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default Club;
