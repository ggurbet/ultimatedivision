// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { PlaceBid } from '@components/common/Card/popUps/PlaceBid';
import { MarketplaceTimer } from '@/app/components/MarketPlace/MarketplaceTimer';
import { RootState } from '@/app/store';

import './index.scss';

/** Initial bid value. */
const INITIAL_BID: number = 0;

export const BidArea = () => {
    const { lot } = useSelector((state: RootState) => state.marketplaceReducer);

    const [isOpenPlaceBidPopup, setIsOpenPlaceBidPopup] = useState<boolean>(false);
    const [cardBid, setCardBid] = useState<number>(INITIAL_BID);
    const [currentCardBid, setCurrentCardBid] = useState<number>(INITIAL_BID);
    const [isEndTime, setIsEndTime] = useState(false);

    /** Handle opening of a place bids pop-up. */
    const handleOpenPlaceBidPopup = () => {
        setIsOpenPlaceBidPopup(true);
    };

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
                <button className="buy-now">
                    <span className="buy-now__text">Buy now</span>
                    <span className="buy-now__value">{currentCardBid} COIN</span>
                </button>
            </div>
        </div>
    </div>;
};
