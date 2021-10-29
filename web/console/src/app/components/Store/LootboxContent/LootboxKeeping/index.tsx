// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';
import { RootState } from '@/app/store';
import { useSelector } from 'react-redux';
import { MyCard } from '@/app/components/Club/ClubCardsArea/MyCard';
import { boxStyle } from '@/app/utils/lootboxStyle';
import boxLight from '@static/img/StorePage/BoxContent/boxLight.svg';
import ribbons from '@static/img/StorePage/BoxContent/ribbons.svg';

import './index.scss';

export const LootboxKeeping: React.FC<{ handleOpening: Dispatch<SetStateAction<boolean>> }> = ({ handleOpening }) => {
    const cards = useSelector((state: RootState) => state.lootboxReducer.lootbox);
    const box = boxStyle(cards.length);
    /** variables that describe indexes of first and last cards,
     *  that will be shown when lootbox is openned */
    const FIRST_CARD_INDEX: number = 0;
    const LAST_CARD_INDEX: number = 4;

    return (
        <div className="box-keeping">
            <div className="box-keeping__wrapper">
                <h1 className="box-keeping__title">
                    Card
                </h1>
                <div className="box-keeping__card-wrapper">
                    {cards.slice(FIRST_CARD_INDEX, LAST_CARD_INDEX).map((card, index) =>
                        <div className="box-keeping__card">
                            <MyCard card={card} key={index} />
                        </div>
                    )}
                </div>
                <div className="box-keeping__button-wrapper">
                    <button className="box-keeping__button"
                        onClick={() => handleOpening(false)}>
                        Keep all
                    </button>
                </div>
                <div className="box-keeping__box-wrapper">
                    <img
                        src={boxLight}
                        alt="box light"
                        className="box-keeping__box-light"
                    />
                    <img
                        src={ribbons}
                        alt="ribbons"
                        className="box-keeping__ribbons"
                    />
                    <img
                        className="box-keeping__box-body"
                        src={box.body}
                        alt="box"
                    />
                </div>
            </div>
        </div>
    );
};
