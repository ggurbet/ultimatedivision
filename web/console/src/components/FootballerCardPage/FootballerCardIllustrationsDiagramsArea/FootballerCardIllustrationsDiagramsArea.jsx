// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import './FootballerCardIllustrationsDiagramsArea.scss';

import { FootballerCardIllustrationsDiagram }
    from '../FootballerCardIllustrationsDiagram/FootballerCardIllustrationsDiagram';

export const FootballerCardIllustrationsDiagramsArea = () => {

    const diagramData = [
        {
            id: 1,
            name: 'Physical',
            min: 100,
            max: 800,
            value: 688
        },
        {
            id: 2,
            name: 'Mental',
            min: 100,
            max: 500,
            value: 364
        },
        {
            id: 3,
            name: 'Skill',
            min: 400,
            max: 1500,
            value: 1120
        },
        {
            id: 4,
            name: 'Chem. Style',
            min: 100,
            max: 300,
            value: 200
        },
        {
            id: 5,
            name: 'Base stats',
            min: 100,
            max: 600,
            value: 464
        },
        {
            id: 6,
            name: 'In game stats',
            min: 100,
            max: 2800,
            value: 2258
        },
    ];

    return (
        <div className="footballer-card-illustrations-diagram-area">
            {diagramData.map(item => (
                <FootballerCardIllustrationsDiagram
                    key={item.id}
                    {...item}
                />
            ))}
        </div>
    );
};
