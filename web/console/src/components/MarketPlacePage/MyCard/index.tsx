/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { PlayerCard } from '../../PlayerCard';

import { Card } from '../../../store/reducers/footballerCard';

import './index.scss';

export const MyCard: React.FC<{ card: Card; place?: string }> = ({ card, place }) =>
    <div
        className="marketplace-myCard"
    >
        <img
            className="marketplace-myCard__confirm-icon"
            src={card.mainInfo.confirmIcon}
            alt="Confirm icon"
        />
        <img
            className="marketplace-myCard__price-gold"
            src={card.mainInfo.priceGoldIcon}
            alt="Price icon"
        />
        <PlayerCard
            card={card}
            parentClassName={'marketplace-myCard'}
        />
    </div>;
