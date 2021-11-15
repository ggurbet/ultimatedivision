// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

export const MatchFinder: React.FC = () =>
    /** TODO: reworks this component by fetch datas. Right now uses just mock datas. */
    <div className="match-finder">
        <h1 className="match-finder__title">
            LOOKING FOR A MATCH
        </h1>
        <div className="match-finder__timer">
            <span className="match-finder__timer__text">
                53:03
            </span>
        </div>
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
            />
        </div>
    </div>;
