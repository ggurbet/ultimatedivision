// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { Link } from 'react-router-dom';
import { useSelector } from 'react-redux';

import { Lot } from '@/marketplace';
import { RootState } from '@/app/store';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import WalletService from '@/wallet/service';
import { ToastNotifications } from '@/notifications/service';

import { PlayerCard } from '@components/common/PlayerCard';
import { MarketplaceTimer } from '@components/MarketPlace/MarketplaceTimer';

import './index.scss';

const ONE_COIN = 1;

export const MarketPlaceFootballerCard: React.FC<{ lot: Lot; handleShowModal: (lot: Lot) => void; place?: string }> =
    ({ lot, handleShowModal }) => {
        const [isEndTime, setIsEndTime] = useState(false);

        const user = useSelector((state: RootState) => state.usersReducer.user);

        const marketplaceClient = new MarketplaceClient();
        const marketplaceService = new Marketplaces(marketplaceClient);

        /** buys a nft */
        const buyNowButton = async() => {
            try {
                const walletService = new WalletService(user);
                const offerData = await marketplaceService.offer(lot.cardId);

                await walletService.buyListing(offerData);
            } catch (e) {
                ToastNotifications.somethingWentsWrong();
            }
        };

        return <div className="marketplace-playerCard">
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


