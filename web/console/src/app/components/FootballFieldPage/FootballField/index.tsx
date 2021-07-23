/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { DragEvent } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FootballFieldCardSelection } from '../FootballFieldCardSelection';
import { FootballFieldPlayingArea } from '../FotballFieldPlayingArea';

import { RootState } from '../../../store';
import { removeCard } from '../../../store/reducers/footballField';

import './index.scss';

const FootballField: React.FC = () => {
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
            <FootballFieldPlayingArea />
            <FootballFieldCardSelection />
        </div>
    );
};

export default FootballField;
