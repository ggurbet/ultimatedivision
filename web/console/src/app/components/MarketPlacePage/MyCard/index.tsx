//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { PlayerCard } from '@components/PlayerCard';

import { Link } from 'react-router-dom';
import { RouteConfig } from '@/app/routes';

import { Card } from '@/app/store/reducers/footballerCard';

import './index.scss';

export const MyCard: React.FC<{ card: Card }> = ({ card }) =>
    <div
        className="marketplace-myCard"
    >
        <Link
            className="marketplace-myCard__link"
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
