// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';
import { RootState } from '@/app/store';

import { FilterField } from '@components/common/FilterField';
import { ClubCardsArea } from '@components/ClubPage/ClubCardsArea';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {
    const cards = useSelector((state: RootState) => state.cardReducer);

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
