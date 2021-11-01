// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardIllustrationsDiagram }
    from '@/app/components/common/Card/CardIllustrationsDiagram';

import { Card } from '@/card';

import './index.scss';

export const FootballerCardIllustrationsDiagramsArea: React.FC<{ card: Card }> = ({ card }) => {
    const diagramData = card.diagramArea;

    return (
        <div className="footballer-card-illustrations-diagram-area">
            {diagramData.map((item, index) =>
                <FootballerCardIllustrationsDiagram
                    key={index}
                    props={item}
                />,
            )}
        </div>
    );
};
