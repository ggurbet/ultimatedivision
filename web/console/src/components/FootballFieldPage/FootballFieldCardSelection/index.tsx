/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */
import { useSelector } from 'react-redux';

import { FilterField } from './FilterField';
import { PlayingAreaFootballerCard } from './PlayingAreaFootballerCard';

import { RootState } from '../../../store';

import './index.scss';

export const FootballFieldCardSelection = () => {
    const cardList = useSelector((state: RootState) => state.cardReducer);

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            {cardList.map((card, index) =>
                <a key={index} href="#playingArea" className="card-selection__card">
                    <PlayingAreaFootballerCard card={card} />
                </a>,
            )}
        </div>
    );
};
