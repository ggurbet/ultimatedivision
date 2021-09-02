// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { PlayerCard } from '@components/common/PlayerCard';

/** TODO: replace it by class fields */
import confirmIcon from '@static/img/MarketPlacePage/MyCard/ok.svg';
import priceGoldIcon from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';

import { Link } from 'react-router-dom';
import { RouteConfig } from '@/app/router';

import { Card } from '@/card';

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
                src={confirmIcon}
                alt="Confirm icon"
            />
            <img
                className="marketplace-myCard__price-gold"
                src={priceGoldIcon}
                alt="Price icon"
            />
            <PlayerCard
                card={card}
                parentClassName={'marketplace-myCard'}
            />
        </Link>
    </div>;
