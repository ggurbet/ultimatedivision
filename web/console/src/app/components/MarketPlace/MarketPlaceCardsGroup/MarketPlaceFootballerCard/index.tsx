// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { Link } from 'react-router-dom';

import { PlayerCard } from '@components/common/PlayerCard';
import { MarketplaceTimer } from '@components/MarketPlace/MarketplaceTimer';
import { Lot } from '@/marketplace';

import './index.scss';

const ONE_COIN = 1;

export const MarketPlaceFootballerCard: React.FC<{ lot: Lot; handleShowModal: (lot: Lot) => void; place?: string }> =
    ({ lot, handleShowModal }) => {
        /** TODO: add function entity */
        const buyNowButton = () => { };
        const bidButton = () => { };

        const [isEndTime, setIsEndTime] = useState(false);

        return <div className="marketplace-playerCard"
            onClick={() => bidButton()}>
            <Link
                className="marketplace-playerCard__link"
                to={`/lot/${lot.cardId}`}
            >
                <PlayerCard
                    id={lot.cardId}
                    className={'marketplace-playerCard__image'}
                />
            </Link>
            <div className="marketplace-playerCard__info">
                <div className="marketplace-playerCard__text">
                    <p className="marketplace-playerCard__text__info"> Max Bid</p>
                    <span className="marketplace-playerCard__price">
                        {lot.maxPrice} {lot.maxPrice > ONE_COIN ? 'coins' : 'coin'}
                    </span>
                </div>
                <button className="marketplace-playerCard__button" onClick={() => handleShowModal(lot)}>
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
            {isEndTime ?
                <div className="marketplace-playerCard__timer">0 : 0 : 0</div> :
                <MarketplaceTimer lot={lot} setIsEndTime={setIsEndTime} isEndTime={isEndTime} className="marketplace-playerCard__timer" />
            }
        </div>;
    };


