// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { MarketplaceClient } from '@/api/marketplace';
import { RootState } from '@/app/store';
import { BidsMakeOfferTransaction } from '@/casper/types';
import { Marketplaces } from '@/marketplace/service';
import WalletService from '@/wallet/service';

import closePopup from '@static/img/FootballerCardPage/close-popup.svg';

import './index.scss';

type PlaceBidTypes = {
    setIsOpenPlaceBidPopup: (isOpenPlaceBidPopup: boolean) => void;
    setCardBid: (cardBid: number) => void;
    setCurrentCardBid: (currentCardBid: number) => void;
    cardBid: number;
};

export const PlaceBid: React.FC<PlaceBidTypes> = ({ setIsOpenPlaceBidPopup, setCardBid, setCurrentCardBid, cardBid }) => {
    const user = useSelector((state: RootState) => state.usersReducer.user);
    const { lot } = useSelector((state: RootState) => state.marketplaceReducer);

    const marketplaceClient = new MarketplaceClient();
    const marketplaceService = new Marketplaces(marketplaceClient);

    /** Sets current bid to state and close popup. */
    const handlePlaceCurrentBid = async() => {
        const makeOfferData = await marketplaceService.placeBid(lot.cardId, cardBid);

        const walletService = new WalletService(user);

        const marketplaceMakeOfferTransaction = new BidsMakeOfferTransaction(
            makeOfferData.address,
            makeOfferData.rpcNodeAddress,
            makeOfferData.tokenId,
            makeOfferData.contractHash,
            makeOfferData.tokenContractHash,
            cardBid
        );

        await walletService.makeOffer(marketplaceMakeOfferTransaction);

        setCurrentCardBid(cardBid);
        setIsOpenPlaceBidPopup(false);
    };

    return (
        <div className="pop-up__place-bid">
            <div className="pop-up__place-bid__wrapper">
                <img
                    className="pop-up__place-bid__close"
                    src={closePopup}
                    alt="Close popup"
                    onClick={() => setIsOpenPlaceBidPopup(false)}
                />
                <span className="pop-up__place-bid__title">Plase a bid</span>
                <div className="pop-up__place-bid__block">
                    <input
                        className="pop-up__place-bid__input"
                        min={0}
                        type="number"
                        onChange={(e) => setCardBid(Number(e.target.value))}
                        value={cardBid}
                    />
                </div>
                <button className="pop-up__place-bid__btn" onClick={handlePlaceCurrentBid}>
                    <span className="pop-up__place-bid__btn-text">BID</span>
                </button>
            </div>
        </div>
    );
};
