// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';

import { RootState } from '@/app/store';
import { useSelector } from 'react-redux';

import { MyCard } from '@/app/components/Club/ClubCardsArea/MyCard';

import boxBody from '@static/img/StorePage/BoxContent/boxBody.svg';
import boxLight from '@static/img/StorePage/BoxContent/boxLight.svg';
import ribbons from '@static/img/StorePage/BoxContent/ribbons.svg';
import './index.scss';

export const LootboxKeeping: React.FC<{ handleOpening: Dispatch<SetStateAction<boolean>> }> = ({ handleOpening }) => {
    // TODO: replace by backend data
    /* eslint-disable-next-line */
    const cards = useSelector((state: RootState) => state.cardsReducer.club.slice(0, 5));

    return (
        <div className="box-keeping">
            <h1 className="box-keeping__title">
                Card
            </h1>
            <div className="box-keeping__card-wrapper">
                {cards.map((card, index) =>
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
                <img src={boxBody} alt="box" />
            </div>
        </div>
    );
};
