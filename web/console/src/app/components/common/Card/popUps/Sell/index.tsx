// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useDispatch, useSelector } from 'react-redux';
import { useEffect, useState } from 'react';

import { RootState } from '@/app/store';
import { createLot } from '@/app/store/actions/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { setCurrentUser } from '@/app/store/actions/users';
import { ToastNotifications } from '@/notifications/service';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import { MarketCreateLotTransaction } from '@/casper/types';
import WalletService from '@/wallet/service';

import closePopup from '@static/img/FootballerCardPage/close-popup.svg';

import './index.scss';

type SellTypes = {
    setIsOpenSellPopup: (isOpenSellPopup: boolean) => void;
};

/** Mock lot creating stats */
const DEFAULT_MIN_BID = 3;
const DEFAULT_MAX_BID = 5;
const MOCK_PERIOD = 3;
const MOCK_PERIOD__TRANSACTION = 300000;

export const Sell: React.FC<SellTypes> = ({ setIsOpenSellPopup }) => {
    const dispatch = useDispatch();
    const [minBidPrice, setMinBidPrice] = useState(DEFAULT_MIN_BID);
    const [maxBidPrice, setMaxBidPrice] = useState(DEFAULT_MAX_BID);

    const marketplaceClient = new MarketplaceClient();
    const marketplaceService = new Marketplaces(marketplaceClient);

    const user = useSelector((state: RootState) => state.usersReducer.user);
    const { card } = useSelector((state: RootState) => state.cardsReducer);

    /** creates lot */
    const setCreatedLot = async() => {
        try {
            const transactionData = await marketplaceService.lotData(card.id);
            const walletService = new WalletService(user);

            const convertedMinBidPrice = Number(minBidPrice);
            const convertedMaxBidPrice = Number(maxBidPrice);

            const marketplaceLotTransaction = new MarketCreateLotTransaction(
                transactionData.address,
                transactionData.addressNodeServer,
                transactionData.tokenId,
                transactionData.contractHash,
                convertedMinBidPrice,
                MOCK_PERIOD__TRANSACTION,
                convertedMaxBidPrice,
            );

            await walletService.createLot(marketplaceLotTransaction);

            dispatch(createLot(new CreatedLot(card.id, 'card', convertedMinBidPrice, convertedMaxBidPrice, MOCK_PERIOD)));

            setIsOpenSellPopup(false);
        }
        catch (e) {
            ToastNotifications.somethingWentsWrong();
        }
    };

    /** changes min price */
    const handleMinPriceChanges = (e: any) => {
        setMinBidPrice(e.target.value);
    };

    /** changes max price */
    const handleMaxPriceChanges = (e: any) => {
        setMaxBidPrice(e.target.value);
    };

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    useEffect(() => {
        setUser();
    }, []);


    return (
        <div className="pop-up__sell">
            <div className="pop-up__sell__wrapper">
                <img
                    className="pop-up__sell__close"
                    src={closePopup}
                    alt="Close popup"
                    onClick={() => setIsOpenSellPopup(false)}
                />
                <span className="pop-up__sell__title">SELL</span>
                <span className="pop-up__sell__price-title">Minimum Price</span>
                <div className="pop-up__sell__price-block">
                    <input
                        className="pop-up__sell__input"
                        type="number"
                        min={1000}
                        onChange={(e) => handleMinPriceChanges(e)}
                        value={minBidPrice}
                    />
                </div>
                <span className="pop-up__sell__price-title">Max Price</span>
                <div className="pop-up__sell__price-block">
                    <input
                        className="pop-up__sell__input"
                        type="number"
                        max={10000}
                        onChange={(e) => handleMaxPriceChanges(e)}
                        value={maxBidPrice}
                    />
                </div>
                <span className="pop-up__sell__price-title">Auction time</span>
                <div className="pop-up__sell__auction-block">
                    <button className="auction-hours">3H</button>
                    <button className="auction-hours">24H</button>
                    <button className="auction-hours">72H</button>
                </div>
                <button className="pop-up__sell__btn" onClick={() => setCreatedLot()}>
                    <span className="pop-up__sell__btn-text">BID</span>
                </button>
            </div>
        </div>);
};
