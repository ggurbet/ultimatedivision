// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import currency from '@static/img/FootballerCardPage/currency.svg';
import closePopup from '@static/img/FootballerCardPage/close-popup.svg';

import './index.scss';

type PlaceBidTypes = {
    setIsOpenPlaceBidPopup: (isOpenPlaceBidPopup: boolean) => void;
    setCardBid: (cardBid: number) => void;
    setCurrentCardBid: (currentCardBid: number) => void;
    cardBid: number;
};

export const PlaceBid: React.FC<PlaceBidTypes> = ({ setIsOpenPlaceBidPopup, setCardBid, setCurrentCardBid, cardBid }) => {
    /** Sets current bid to state and close popup. */
    const handlePlaceCurrentBid = () => {
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
                    <img
                        className="pop-up__place-bid__coin"
                        src={currency}
                        alt="currency img"
                    />
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
