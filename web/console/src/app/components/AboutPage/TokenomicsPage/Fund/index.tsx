// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import './index.scss';

const Fund: React.FC = () => {
    return (
        <div className="ud-fund">
            <h1 className="ud-fund__title">
                UD Fund
            </h1>
            <p className="ud-fund__description">
                A special fund will be allocated to support the development of the game’s ecosystem and invite outside developers.
                The fund will be managed by the UDT holder in a decentralised fashion.
            </p>
        </div>
    )
}

export default Fund;
