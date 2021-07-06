import React from 'react';
import './FootballFieldCardSelection.scss';

import { useSelector } from 'react-redux';
import { RootState } from '../../../store';
import { PlayingAreaFootballerCard }
    from '../FotballFieldPlayingArea/PlayingAreaFootballerCard/PlayingAreaFootballerCard';

export const FootballFieldCardSelection = () => {
    const cardList = useSelector((state: RootState) => state.cardReducer);

    return (
        <div id="cardList" className="card-selection">
            {cardList.map(card => (
                <a href="#playingArea" className="card-selection__card">
                    <PlayingAreaFootballerCard card={card} />
                </a>
            ))}
        </div>
    )
}
