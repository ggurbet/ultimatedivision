// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import closePopup from '@static/img/FootballerCardPage/close-popup.svg';

import './index.scss';

type SellTypes = {
    setIsOpenSellPopup: (isOpenSellPopup: boolean) => void;
};

export const Sell: React.FC<SellTypes> = ({ setIsOpenSellPopup }) =>
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
                    onChange={() => {}}
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
                    onChange={() => {}}
                    value="100"
                />
            </div>
            <span className="pop-up__sell__price-title">Auction time</span>
            <div className="pop-up__sell__auction-block">
                <button className="auction-hours">3H</button>
                <button className="auction-hours">24H</button>
                <button className="auction-hours">72H</button>
            </div>
            <button className="pop-up__sell__btn" onClick={() => setIsOpenSellPopup(false)}>
                <span className="pop-up__sell__btn-text">BID</span>
            </button>
        </div>
    </div>;


