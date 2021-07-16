// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { FootballerCardIllustrationsDiagram }
    from '../FootballerCardIllustrationsDiagram';

import { RootState } from '../../../store';

import './index.scss';

export const FootballerCardIllustrationsDiagramsArea: React.FC = () => {
    const FIRST_CARD_INDEX = 0;
    const diagramData = useSelector((state: RootState) =>
        state.cardReducer[FIRST_CARD_INDEX].diagram);

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
