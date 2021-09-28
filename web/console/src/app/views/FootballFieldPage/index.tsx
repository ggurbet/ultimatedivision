// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DragEvent } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FootballFieldCardSelection } from
    '@components/FootballField/FootballFieldCardSelection';
import { FootballFieldPlayingArea } from
    '@components/FootballField/FotballFieldPlayingArea';

import { RootState } from '@/app/store';
import { getClub, removeCard } from '@/app/store/actions/club';

import './index.scss';

const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(getClub());
    const fieldSetup = useSelector((state: RootState) => state.clubReducer.options);
    const cardSelectionVisibility = useSelector((state: RootState) => state.clubReducer.options.showCardSeletion);


    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
    };

    /** TO DO: ADD TYPE FOR Event */
    function drop(e: any) {
        if (e.target.className === 'football-field__wrapper') {
            dispatch(removeCard(fieldSetup.dragStart));
        }
    };

    return (
        <div className="football-field"
            onDrop={e => drop(e)}
            onDragOver={e => dragOverHandler(e)}
        >
            <h1 className="football-field__title">Football Field</h1>
            <FootballFieldPlayingArea />
            <div
                style={{ height: cardSelectionVisibility ? 'unset' : '0' }}
                className="football-field__wrapper"
            >
                < FootballFieldCardSelection />
            </div>
        </div>
    );
};

export default FootballField;
