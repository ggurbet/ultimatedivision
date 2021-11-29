// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';

import { Card } from '@/card';
import { CardEditIdentificators } from '@/api/club';
import { deleteCard } from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';

import './index.scss';

export const FootballerCard: React.FC<{ card: Card; index?: number; place?: string }> = ({ card }) => {
    const dispatch = useDispatch();
    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);
    const [isVisibile, setIsVisibile] = useState(false);
    const visibilityBlock = isVisibile ? '-active' : '-inactive';

    /** show/hide delete block, preventing scroll to cardSelection */
    function handleVisibility(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        setIsVisibile((prevVisability) => !prevVisability);
    }
    /** remove player card implementation */
    function handleDeletion(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        e.preventDefault();
        dispatch(deleteCard(new CardEditIdentificators(squad.clubId, squad.id, card.id)));
    }

    return (
        <div onClick={handleVisibility} className="footballer-card">
            <div
                className={`football-field-card__wrapper${visibilityBlock}`}
            ></div>
            <PlayerCard card={card} parentClassName="footballer-card" />
            <div
                onClick={handleDeletion}
                className={`footballer-card__control${visibilityBlock}`}
            >
                &#10006; delete a player
            </div>
        </div>
    );
};
