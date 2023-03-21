// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Link } from 'react-router-dom';

import { PlayerCard } from '@components/common/PlayerCard';
import { Lot } from '@/marketplace';

import './index.scss';

const ONE_COIN = 1;

export const MarketPlaceFootballerCard: React.FC<{ lot: Lot; place?: string }> =
    ({ lot }) => {
        /** TODO: add function entity */
        const buyNowButton = () => { };
        const bidButton = () => { };

        return <div className="marketplace-playerCard">
            <Link
                className="marketplace-playerCard__link"
                to={`/lot/${lot.id}`}
            >
                <PlayerCard
                    id={lot.card.id}
                    className={'marketplace-playerCard__image'}
                />
                <div className="marketplace-playerCard__info">
                    <div className="marketplace-playerCard__text">
                        <p className="marketplace-playerCard__text__info"> Max Bid</p>
                        <span className="marketplace-playerCard__price" onClick={() => bidButton()}>
                            {lot.maxPrice} {lot.maxPrice > ONE_COIN ? 'coins' : 'coin'}
                        </span>
                    </div>
                    <button className="marketplace-playerCard__button">
                        Bid
                    </button>
                </div>
                <div className="marketplace-playerCard__info">
                    <div className="marketplace-playerCard__text">
                        <p className="marketplace-playerCard__text__info">Current bid</p>
                        <span className="marketplace-playerCard__price">
                            {lot.currentPrice} {lot.currentPrice > ONE_COIN ? 'coins' : 'coin'}
                        </span>
                    </div>
                    <button className="marketplace-playerCard__button" onClick={() => buyNowButton()}>
                        Buy now
                    </button>
                </div>
                {/** TODO: change to real data. */}
                <div className="marketplace-playerCard__timer">
                    3 : 10 : 15
                </div>
            </Link>
        </div>;
    };


