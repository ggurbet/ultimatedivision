/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { DragEvent } from 'react';
import './FootballField.scss';

import { useDispatch, useSelector } from 'react-redux';
import { FootballFieldPlayingArea }
    from '../FotballFieldPlayingArea/FootballFieldPlayingArea';
import { FootballFieldInformation }
    from '../FootballFieldInformation/FootballFieldInformation';
import { FootballFieldCardSelection }
    from '../FootballFieldCardSelection/FootballFieldCardSelection';
import { RootState } from '../../../store';
import { removeCard } from '../../../store/reducers/footballField';

export const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    const dragItemPosition = useSelector((state: RootState) => state.fieldReducer.options.dragStart);
    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
    };

    /** TO DO: ADD TYPE FOR Event */
    function drop(e: any) {
        if (e.target.className === 'football-field__wrapper') {
            dispatch(removeCard(dragItemPosition));
        }
    };

    return (
        <div className="football-field"
            onDrop={e => drop(e)}
            onDragOver={e => dragOverHandler(e)}
        >
            <h1 className="football-field__title">Football Field</h1>
            <div className="football-field__wrapper"
            >
                <FootballFieldPlayingArea />
                <FootballFieldInformation />
            </div>
            <FootballFieldCardSelection />
        </div>
    );
};
