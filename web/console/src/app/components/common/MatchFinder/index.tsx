// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';


import { Timer } from './Timer';
import { AutoCloseTimer } from './AutoCloseTimer';

import { RootState } from '@/app/store';
import { startSearchingMatch } from '@/app/store/actions/clubs';

import './index.scss';

const MatchFinder: React.FC = () => {
    const dispatch = useDispatch();
    const { isSearchingMatch } = useSelector((state: RootState) => state.clubsReducer);

    /** closes MatchFinder component */
    const closeMatchFinder = () => {
        dispatch(startSearchingMatch(false));
    };

    /** TODO: it will set after websocket implementation */
    const [isMatchFound, setIsMatchFound] = useState<boolean>(false);

    /** TODO: it will set after websocket implementation */
    const [title, setTitle] = useState<string>('LOOKING FOR A MATCH');

    return isSearchingMatch ? <section className={isMatchFound ? 'match-finder__wrapper' : ''}>
        <div className="match-finder">
            <h1 className="match-finder__title">
                {title}
            </h1>
            {isMatchFound ? <AutoCloseTimer /> : <Timer />}
            <div className="match-finder__form">
                <input
                    className="match-finder__form__accept"
                    value="Accept"
                    type="button"
                />
                <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={closeMatchFinder}
                />
            </div>
        </div>
    </section> : <></>;
};

export default MatchFinder;
