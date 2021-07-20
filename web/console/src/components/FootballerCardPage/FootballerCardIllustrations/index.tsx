/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { FootballerCardIllustrationsDiagramsArea } from '../FootballerCardIllustrationsDiagramsArea';
import { FootballerCardIllustrationsRadar } from '../FootballerCardIllustrationsRadar';

import './index.scss';
import { Card } from '../../../store/reducers/footballerCard';
import { PlayerCard } from '../../PlayerCard';

export const FootballerCardIllustrations: React.FC<{card: Card}> = ({ card }) =>
    <div className="footballer-card-illustrations">
        <div className="footballer-card-illustrations__card">
            <PlayerCard card={card} parentClassName="footballer-card-illustrations__card"/>
        </div>
        <FootballerCardIllustrationsRadar card={card} />
        <FootballerCardIllustrationsDiagramsArea card={card} />
    </div>;
