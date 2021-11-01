// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Link } from 'react-router-dom';

import { PlayerCard } from '@components/common/PlayerCard';

import currentBid
    from '@static/img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/bid.svg';
import priceGoldIcon
    from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';

import { Lot } from '@/marketplace';

import './index.scss';

export const MarketPlaceFootballerCard: React.FC<{ lot: Lot; place?: string }> = ({ lot }) =>
    <div
        className="marketplace-playerCard"
    >
        <Link
            className="marketplace-playerCard__link"
            to={`/lot/${lot.id}`}
        >
            <PlayerCard
                card={lot.card}
                parentClassName={'marketplace-playerCard'}
            />
            {/** TODO: fetch datas from back-end. Now it is just statis images */}
            <div className="marketplace-playerCard__price">
                <img className="marketplace-playerCard__price__picture"
                    src={priceGoldIcon}
                    alt="Player price" />
                <span className="marketplace-playerCard__price__current">
                    {lot.currentPrice}
                </span>
                <img className="marketplace-playerCard__price__status"
                    src={currentBid}
                    alt="Price status" />
            </div>
        </Link>
    </div >;
