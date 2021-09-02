// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';
import { RootState } from '@/app/store';
import { useCards } from '@/app/hooks/cards';
import { userCards } from '@/app/store/actions/cards';

import { FilterField } from '@components/common/FilterField';
import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {
    /** TODO: decide use custom hook or directly dispatch thunk into useEffect*/
    useCards(userCards);

    const cards = useSelector((state: RootState) => state.cardsReducer.clubCards);

    return (
        <section className="club">
            <FilterField
                title="My cards"
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
