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
                UD DAO fund will be allocated to support the development of the game's ecosystem and invite outside developers.
                The fund will be incorporated under decentralised governance model and will be managed through voting by the
                UDT token holders.
                <br /><br />
                Voting will take place to determine short-term goals of the project. Members can propose their ideas and get
                funding based on community voting. The entire UD planning and management will gradually evolve into DAO
            </p>
        </div>
    )
}

export default Fund;
