import React from 'react';
import './index.scss';

import { useSelector } from 'react-redux';
import { RootState } from '../../../store';
import { PlayingAreaFootballerCard } from './PlayingAreaFootballerCard';
import { FilterField } from './FilterField';

export const FootballFieldCardSelection = () => {
    const cardList = useSelector((state: RootState) => state.cardReducer);

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            {cardList.map((card, index) =>
                <a key={index} href="#playingArea" className="card-selection__card">
                    <PlayingAreaFootballerCard card={card} />
                </a>
            )}
        </div>
    );
};
