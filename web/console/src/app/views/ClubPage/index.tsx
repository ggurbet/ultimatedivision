// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { filteredCards } from '@/app/store/actions/cards';
import { RootState } from '@/app/store';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {
    const { cards } = useSelector((state: RootState) => state.cardsReducer.cards);

    return (
        <section className="club">
            <FilterField
                title="My cards"
                thunk={filteredCards}
            />
            <ClubCardsArea
                cards={cards}
            />
            <Paginator
                itemCount={cards.length} />
        </section>
    );
};

export default Club;
