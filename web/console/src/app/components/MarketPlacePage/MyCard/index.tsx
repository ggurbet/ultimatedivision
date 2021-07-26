/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { PlayerCard } from '@playerCard';

import { Link } from 'react-router-dom';
import { RouteConfig } from '@routes';

import { Card } from '@store/reducers/footballerCard';

import './index.scss';

export const MyCard: React.FC<{ card: Card }> = ({ card }) =>
    <div
        className="marketplace-myCard"
    >
        <Link
            style={{ textDecoration: 'none' }}
            to={{
                pathname: RouteConfig.FootballerCard.path,
                state: {
                    card,
                },
            }}
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
        </Link>
    </div>;
