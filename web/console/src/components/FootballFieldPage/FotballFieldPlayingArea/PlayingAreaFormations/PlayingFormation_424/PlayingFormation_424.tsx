/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingFormation_424.scss';
import { FootballField } from '../../../../../types/footballField';

export const PlayingFormation_424: React.FC<{props: FootballField}> = ({props}) => {
    return (
        <div className="playing-formation-424">
            {props.cardsList.map(card => (
                <div
                    key={card.id}
                    className="playing-formation-424__card"
                >
                </div>
            ))}
        </div>
    )
}
