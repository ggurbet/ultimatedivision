// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';
import { MyCard } from '@/app/components/Club/ClubCardsArea/MyCard';
import { RootState } from '@/app/store';
import { boxStyle } from '@/app/utils/lootboxStyle';
import boxLight from '@static/img/StorePage/BoxContent/boxLight.svg';
import ribbons from '@static/img/StorePage/BoxContent/ribbons.svg';

import './index.scss';


export const LootboxOpening = () => {
    const FIRST_CARD = 0;
    const cards = useSelector((state: RootState) => state.lootboxReducer.lootbox);

    const box = boxStyle(cards.length);

    return (
        <div className="box-animation">
            <div className="box-animation__box-container">
                <img
                    src={box.body}
                    alt="box body"
                    className={`box-animation__box-body ${cards.length > 5 && "box-animation__box-body__cool"}`}
                />
                <img
                    src={box.cover}
                    alt="box cover"
                    className="box-animation__box-cover"
                />
                <img
                    src={boxLight}
                    alt="shadow image"
                    className="box-animation__light"
                />
                <img
                    src={ribbons}
                    alt="ribbons"
                    className="box-animation__ribbons"
                />
            </div>
            <div className="box-animation__card-container">
                <div className="box-animation__card-container-backlight">
                    <MyCard card={cards[FIRST_CARD]} />
                </div>
            </div>
        </div>
    );
};
