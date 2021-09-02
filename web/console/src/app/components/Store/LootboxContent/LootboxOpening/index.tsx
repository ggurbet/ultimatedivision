// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MyCard } from '@/app/components/Club/ClubCardsArea/MyCard';
import { RootState } from '@/app/store';
import boxBody from '@static/img/StorePage/BoxContent/boxBody.svg';
import boxCover from '@static/img/StorePage/BoxContent/boxCover.svg';
import boxLight from '@static/img/StorePage/BoxContent/boxLight.svg';
import ribbons from '@static/img/StorePage/BoxContent/ribbons.svg';
import { useSelector } from 'react-redux';

import './index.scss';


export const LootboxOpening = () => {
    // TODO: will be replaced by backend data
    // eslint-disable-next-line
    const card = useSelector((state: RootState) => state.cardsReducer.clubCards[0]);

    return (
        <div className="box-animation">
            <div className="box-animation__box-wrapper">
                <img
                    src={boxBody}
                    alt="box body"
                    className="box-animation__box-body"
                />
                <img
                    src={boxCover}
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
            <div className="box-animation__card-wrapper">
                <div className="box-animation__card-wrapper-backlight">
                    {card &&
                        <MyCard card={card} />
                    }
                </div>
            </div>
        </div>
    );
};
