// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { PlaceBid } from '@components/common/Card/popUps/PlaceBid';
import { MarketplaceTimer } from '@/app/components/MarketPlace/MarketplaceTimer';
import { RootState } from '@/app/store';
import WalletService from '@/wallet/service';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import { ToastNotifications } from '@/notifications/service';

import './index.scss';

export const BidArea = () => {
    const { lot } = useSelector((state: RootState) => state.marketplaceReducer);

    const [isOpenPlaceBidPopup, setIsOpenPlaceBidPopup] = useState<boolean>(false);
    const [cardBid, setCardBid] = useState<number>(lot.currentPrice);
    const [currentCardBid, setCurrentCardBid] = useState<number>(lot.currentPrice);
    const [isEndTime, setIsEndTime] = useState(false);

    const user = useSelector((state: RootState) => state.usersReducer.user);

    const marketplaceClient = new MarketplaceClient();
    const marketplaceService = new Marketplaces(marketplaceClient);

    /** Handle opening of a place bids pop-up. */
    const handleOpenPlaceBidPopup = () => {
        setIsOpenPlaceBidPopup(true);
    };

    /** buys an nft */
    const buyNow = async() => {
        try {
            const walletService = new WalletService(user);
            const offerData = await marketplaceService.offer(lot.cardId);

            await walletService.buyListing(offerData);
        } catch (e) {
            ToastNotifications.somethingWentsWrong();
        }
    };

    useEffect(() => {
        setCardBid(lot.currentPrice);
        setCurrentCardBid(lot.currentPrice);
    }, [lot]);

    return <div className="footballer-card-price">
        {isOpenPlaceBidPopup &&
            <PlaceBid
                setCurrentCardBid={setCurrentCardBid}
                setIsOpenPlaceBidPopup={setIsOpenPlaceBidPopup}
                setCardBid={setCardBid}
                cardBid={cardBid}
            />
        }
        <div className="footballer-card-price__wrapper">
            <div className="footballer-card-price__info-area">
                <div className="footballer-card-price__bid">
                    <div className="bid">
                        <span className="bid__title">Current bid:</span>
                        <span className="bid__value">{currentCardBid}</span>
                    </div>
                </div>
                <div className="footballer-card-price__auction">
                    <span className="auction-title">
                        Auction expires in:
                    </span>
                    {isEndTime ?
                        <div className="auction-expire-time"> 0 : 0 : 0 </div> :
                        <MarketplaceTimer lot={lot} setIsEndTime={setIsEndTime} isEndTime={isEndTime} className="auction-expire-time" />}
                </div>
            </div>
            <div className="footballer-card-price__buttons">
                <button className="place-bid" onClick={handleOpenPlaceBidPopup}>
                    <span className="place-bid__text">Plase a bid</span>
                </button>
                <button className="buy-now" onClick={() => buyNow()}>
                    <span className="buy-now__text">Buy now</span>
                    <span className="buy-now__value">{currentCardBid} COIN</span>
                </button>
            </div>
        </div>
    </div>;
};
