/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { DragEvent } from 'react';
import './PlayingFormation_433.scss';
import { FootballField } from '../../../../../types/footballField';
import { useDispatch } from 'react-redux';
import { choseCardPosition }
    from '../../../../../store/reducers/footballField';
import { PlayingAreaFootballerCard }
    from '../../../FootballFieldCardSelection/PlayingAreaFootballerCard/PlayingAreaFootballerCard';
import { exchangeCards }
    from '../../../../../store/reducers/footballField';
import { useState } from 'react';

export const PlayingFormation_433: React.FC<{ props: FootballField }> = ({ props }) => {
    const dispatch = useDispatch();

    const [currentPosition, handleDrag] = useState(-1);
    const [dragTarget, handleDragTarget] = useState(-1);

    function dragOverHandler(e: DragEvent<HTMLDivElement>, index: number) {
        e.preventDefault();
        handleDragTarget(index);
    };

    function dropHandler(e: DragEvent<HTMLDivElement>) {
        dispatch(exchangeCards(currentPosition, dragTarget));
    };

    return (
        <div className="playing-formation-433">
            {props.cardsList.map((card, index) => {
                const data = card.cardData;
                return (
                    <div
                        key={index}
                        className="playing-formation-433__card box"
                        draggable={true}
                        onDragOver={e => dragOverHandler(e, index)}
                        onMouseDown={(e: React.MouseEvent) => handleDrag(index)}
                        onDrop={e => dropHandler(e)}
                    >
                        {
                            data
                                ? <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                : <a
                                    onClick={() => dispatch(choseCardPosition(index))}
                                    href="#cardList"
                                    className="playing-formation-433__link"
                                >
                                </a>
                        }
                    </div>
                )
            })}
        </div>
    )
}