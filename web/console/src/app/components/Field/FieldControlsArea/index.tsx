// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import Clubs from '@/app/components/Field/FieldControlsArea/Clubs';
import { FieldControl } from '@/app/components/Field/FieldControlsArea/FieldControl';

import {
    setCaptain,
    setFormation,
    setTactic,
    startSearchingMatch,
} from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';
import { Control } from '@/app/types/club';
import { SquadCard } from '@/club';

import './index.scss';

export const FieldControlsArea: React.FC = () => {
    const dispatch = useDispatch();
    const [isPossibleToStartMatch, setIsPossibleToStartMatch] =
        useState<boolean>(true);
    const squadCards = useSelector(
        (state: RootState) => state.clubsReducer.activeClub.squadCards
    );
    const EMPTY_CARD_ID = '00000000-0000-0000-0000-000000000000';

    useEffect(() => {
        /** Function checks field cards and compare it with player cards array */
        function isPossibleToStart() {
            const emptyCard = squadCards.find(
                (squadCard: SquadCard) => squadCard.card.id === EMPTY_CARD_ID
            );
            emptyCard
                ? setIsPossibleToStartMatch(false)
                : setIsPossibleToStartMatch(true);
        }
        isPossibleToStart();
    });

    const CONTROLS_FIELDS = [
        new Control('0', 'formation', setFormation, [
            '4-4-2',
            '4-2-4',
            '4-2-2-2',
            '4-3-1-2',
            '4-3-3',
            '4-2-3-1',
            '4-3-2-1',
            '4-1-3-2',
            '5-3-2',
            '4-5-2',
        ]),
        new Control('1', 'tactics', setTactic, [
            'attack',
            'defence',
            'balanced',
        ]),
        new Control('2', 'captain', setCaptain, [
            'Captain 1',
            'Captain 2',
            'Captain 3',
        ]),
    ];

    /** shows matchFinder component */
    const showMatchFinder = () => {
        dispatch(startSearchingMatch(true));
        window.scrollTo({
            top: 0,
            behavior: 'smooth',
        });
    };

    return (
        <div className="field-controls-area__wrapper">
            <Clubs />
            {isPossibleToStartMatch &&
                <input
                    type="button"
                    value="Play"
                    className="field-controls-area__play"
                    onClick={showMatchFinder}
                />
            }
            <div className="field-controls-area">
                <h2 className="field-controls-area__title">information</h2>
                {CONTROLS_FIELDS.map((item, index) =>
                    <FieldControl key={index} props={item} />
                )}
            </div>
        </div>
    );
};
