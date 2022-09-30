// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';
import { useSelector } from 'react-redux';
import Slider from 'react-slick';

import { PlayerCard } from '@/app/components/common/PlayerCard';

import { RootState } from '@/app/store';

import { Card } from '@/card';

import box from '@static/img/StorePage/BoxContent/box.png';

import './index.scss';

export const LootboxKeeping: React.FC<{
    handleLootboxSelection: Dispatch<SetStateAction<boolean>>;
}> = ({ handleLootboxSelection }) => {
    const cards = useSelector((state: RootState) => state.lootboxReducer.lootbox);

    /** variables that describe indexes of first and last cards,
     *  that will be shown when lootbox is openned */
    const REGULAR_CARDS_AMOUNT: number = 5;
    const FIRST_PAGE_START: number = 0;
    const FIRST_PAGE_END: number = 4;
    const SECOND_PAGE_END: number = 8;

    const settings = {
        adaptiveHeight: true,
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 1,
        slidesToScroll: 1,
    };

    return (
        <div className="box-keeping">
            <div className="box-keeping__wrapper">
                <Slider {...settings} className="slider">
                    <div>
                        <div className="box-keeping__card__wrapper">
                            {cards.slice(FIRST_PAGE_START, FIRST_PAGE_END).map((card: Card, index: number) =>
                                <PlayerCard key={index} className="box-keeping__card" id={card.id} />
                            )}
                        </div>
                    </div>
                    <div>
                        <div className="box-keeping__card__wrapper">
                            {cards.slice(FIRST_PAGE_END, SECOND_PAGE_END).map((card: Card, index: number) =>
                                <PlayerCard key={index} className="box-keeping__card" id={card.id} />
                            )}
                        </div>
                    </div>
                    {cards.length > REGULAR_CARDS_AMOUNT &&
                        <div>
                            <div className="box-keeping__card__wrapper">
                                {cards.slice(SECOND_PAGE_END).map((card: Card, index: number) =>
                                    <PlayerCard key={index} className="box-keeping__card" id={card.id} />
                                )}
                            </div>
                        </div>
                    }
                </Slider>
                <div className="box-keeping__button__wrapper">
                    <button className="box-keeping__button" onClick={() => handleLootboxSelection(true)}>
                        Keep all
                    </button>
                </div>
                <div className="box-keeping__box__wrapper">
                    <img className="box-keeping__box" src={box} alt="box" />
                </div>
            </div>
        </div>
    );
};
