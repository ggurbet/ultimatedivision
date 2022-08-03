// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

export const MintingArea: React.FC<{ isInactive: boolean; time: string }> = ({ isInactive, time }) =>
    <div className="minting-area">
        {isInactive &&
            <div className="minting-area__timer">
                <p className="minting-area__timer__text">RESTOCK IN</p>
                <span className="minting-area__timer__time">{time}</span>
            </div>
        }
        <button className="minting-area__button" disabled={isInactive}>
            <span className="minting-area__button__text">MINT</span>
            <span className="minting-area__button__value">100 coin</span>
        </button>
    </div>;

