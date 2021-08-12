// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


import { FootballerCardIllustrationsDiagram }
    from '@components/FootballerCardPage/FootballerCardIllustrationsDiagram';

import { Card } from '@/app/types/fotballerCard';

import './index.scss';

export const FootballerCardIllustrationsDiagramsArea: React.FC<{ card: Card }> = ({ card }) => {
    const FIRST_CARD_INDEX = 0;
    const diagramData = card.diagram;

    return (
        <div className="footballer-card-illustrations-diagram-area">
            {diagramData.map(item =>
                <FootballerCardIllustrationsDiagram
                    key={item.id}
                    props={item}
                />,
            )}
        </div>
    );
};
