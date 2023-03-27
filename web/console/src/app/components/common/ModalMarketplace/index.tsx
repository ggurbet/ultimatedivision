// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';

import { Lot } from '@/marketplace';
import { PlayerCard } from '../PlayerCard';

import CloseModal from '@/app/static/img/MarketPlacePage/marketplaceModal/close.svg';

import './index.scss';

const ONE_COIN = 1;

export const ModalMarketPlace: React.FC<{ lot: Lot; setShowModal: Dispatch<SetStateAction<boolean>> }> =
    ({ lot, setShowModal }) => {
        /** TODO: add function entity */
        const buyNowButton = () => { };
        const bidButton = () => { };

        return <div className="marketplace-modal">
            <div className="marketplace-modal__wrapper">
                <div>
                    <button onClick={() => setShowModal(false)} className="marketplace-modal__close-button">
                        <img src={CloseModal} alt="modal close button" className="marketplace-modal__close-button__img" />
                    </button>
                </div>
                <h2 className="marketplace-modal__title">{lot.card.playerName}</h2>
                <div className="marketplace-modal__content">
                    <PlayerCard id={lot.card.id} className="marketplace-modal__card" />
                    <div className="marketplace-modal__lot">
                        <div className="marketplace-modal__bid">
                            <input className="marketplace-modal__bid__input" placeholder="place a bid|" type="number" />
                            <button onClick={() => bidButton()} className="marketplace-modal__button marketplace-modal__button__bid">
                                    bid
                            </button>
                            <span className="marketplace-modal__bid__label">or</span>
                        </div>
                        <div className="marketplace-modal__buy-now" >
                            <div className="marketplace-modal__timer">
                                {/** TODO: change to real data. */}
                                3 : 30 : 12
                            </div>
                              <div className="marketplace-modal__buy-now__label__mobile">
                                    for
                                    <span>
                                        {lot.currentPrice} {lot.currentPrice > ONE_COIN ? 'coins' : 'coin'}
                                    </span>
                            </div>
                            <button onClick={() => buyNowButton()} className="marketplace-modal__button marketplace-modal__button__buy-now">
                                buy now
                                <div className="marketplace-modal__buy-now__label">
                                    for
                                    <span>
                                        {lot.currentPrice} {lot.currentPrice > ONE_COIN ? 'coins' : 'coin'}
                                    </span>
                                </div>
                            </button>   
                        </div>
                        <div className="marketplace-modal__timer__mobile">
                            {/** TODO: change to real data. */}
                            3 : 30 : 12
                        </div>
                    </div>
                </div>
            </div>
        </div>;
    };


