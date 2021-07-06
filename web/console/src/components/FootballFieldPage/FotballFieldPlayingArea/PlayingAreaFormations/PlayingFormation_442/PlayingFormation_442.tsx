/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingFormation_442.scss';
import { FootballField } from '../../../../../types/footballField';

export const PlayingFormation_442: React.FC<{ props: FootballField }> = ({ props }) => {

    return (
        <div className="playing-formation-442">
            {props.cardsList.map(card => (
                <div
                    key={card.id}
                    className="playing-formation-442__card"
                >
                </div>
            ))}
        </div>
    )
}
