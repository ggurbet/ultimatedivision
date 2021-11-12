// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';
import { useSelector } from 'react-redux';
import Slider from 'react-slick';

import { MyCard } from '@/app/components/Club/ClubCardsArea/MyCard';

import boxLight from '@static/img/StorePage/BoxContent/boxLight.svg';
import ribbons from '@static/img/StorePage/BoxContent/ribbons.svg';

import { RootState } from '@/app/store';
import { boxStyle } from '@/app/utils/lootboxStyle';

import './index.scss';

export const LootboxKeeping: React.FC<{ handleOpening: Dispatch<SetStateAction<boolean>> }> = ({ handleOpening }) => {
    const cards = useSelector((state: RootState) => state.lootboxReducer.lootbox);
    const box = boxStyle(cards.length);
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
                <h1 className="box-keeping__title">
                    Card
                </h1>
                <Slider {...settings} className="slider">
                    <div>
                        <div className="box-keeping__card-wrapper">
                            {cards.slice(FIRST_PAGE_START, FIRST_PAGE_END).map((card, index) =>
                                <div className="box-keeping__card">
                                    <MyCard card={card} key={index} />
                                </div>
                            )}
                        </div>
                    </div>
                    <div>
                        <div className="box-keeping__card-wrapper">
                            {cards.slice(FIRST_PAGE_END, SECOND_PAGE_END).map((card, index) =>
                                <div className="box-keeping__card">
                                    <MyCard card={card} key={index} />
                                </div>
                            )}
                        </div>
                    </div>
                    {
                        cards.length > REGULAR_CARDS_AMOUNT &&
                        <div>
                            <div className="box-keeping__card-wrapper">
                                {cards.slice(SECOND_PAGE_END).map((card, index) =>
                                    <div className="box-keeping__card">
                                        <MyCard card={card} key={index} />
                                    </div>
                                )}
                            </div>
                        </div>
                    }
                </Slider>
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
