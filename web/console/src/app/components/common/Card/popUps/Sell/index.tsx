// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';

import { RootState } from '@/app/store';
import { createLot } from '@/app/store/actions/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { setCurrentUser } from '@/app/store/actions/users';
import { ToastNotifications } from '@/notifications/service';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import { MarketCreateLotTransaction } from '@/marketplace';
import WalletService from '@/wallet/service';

import closePopup from '@static/img/FootballerCardPage/close-popup.svg';

import './index.scss';

type SellTypes = {
    setIsOpenSellPopup: (isOpenSellPopup: boolean) => void;
};

/** Mock lot creating stats */
const MOCK_MIN_BID = 3000;
const MOCK_MAX_BID = 800;
const MOCK_PERIOD = 300000;
const MOCK_REDEMPTION_PRRICE = 30000;

export const Sell: React.FC<SellTypes> = ({ setIsOpenSellPopup }) => {
    const dispatch = useDispatch();

    const marketplaceClient = new MarketplaceClient();
    const marketplaceService = new Marketplaces(marketplaceClient);

    const user = useSelector((state: RootState) => state.usersReducer.user);
    const { card } = useSelector((state: RootState) => state.cardsReducer);

    const setCreatedLot = async() => {
        const transactionData = await marketplaceService.lotData(card.id);
        const walletService = new WalletService(user);

        const marketplaceLotTransaction =
            new MarketCreateLotTransaction(
                transactionData.address,
                transactionData.addressNodeServer,
                transactionData.tokenId,
                transactionData.contractHash,
                MOCK_MIN_BID,
                MOCK_REDEMPTION_PRRICE,
                MOCK_MAX_BID,
            );

        await walletService.createLot(marketplaceLotTransaction);

        dispatch(createLot(new CreatedLot(card.id, 'card', MOCK_MIN_BID, MOCK_MAX_BID, MOCK_PERIOD)));
        setIsOpenSellPopup(false);
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


    return(
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
                        min={0}
                        // TODO: Need add logic to listener and value.
                        onChange={() => { }}
                        value="0"
                    />
                </div>
                <span className="pop-up__sell__price-title">Buy Now Price</span>
                <div className="pop-up__sell__price-block">
                    <input
                        className="pop-up__sell__input"
                        type="number"
                        min={0}
                        // TODO: Need add logic to listener and value.
                        onChange={() => { }}
                        value="100"
                    />
                </div>
                <span className="pop-up__sell__price-title">Auction time</span>
                <div className="pop-up__sell__auction-block">
                    <button className="auction-hours">3H</button>
                    <button className="auction-hours">24H</button>
                    <button className="auction-hours">72H</button>
                </div>
                <button className="pop-up__sell__btn" onClick={() => setCreatedLot() }>
                    <span className="pop-up__sell__btn-text">BID</span>
                </button>
            </div>
        </div>);
};
