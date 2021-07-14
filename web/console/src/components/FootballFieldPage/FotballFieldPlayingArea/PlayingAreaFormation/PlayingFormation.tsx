/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { DragEvent } from 'react';
import './PlayingFormation.scss';
import { useDispatch, useSelector } from 'react-redux';
import { FootballField } from '../../../../types/footballField';
import { choseCardPosition, setDragStart, setDragTarget }
    from '../../../../store/reducers/footballField';
import { PlayingAreaFootballerCard }
    from '../../FootballFieldCardSelection/PlayingAreaFootballerCard/PlayingAreaFootballerCard';
import { exchangeCards }
    from '../../../../store/reducers/footballField';
import { RootState } from '../../../../store';

export const PlayingFormation: React.FC<{ props: FootballField; formation: string }> = ({ props, formation }) => {
    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer.options);
    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLAnchorElement>) {
        e.preventDefault();
    };
    /** exchange player cards implemnentation:
     *  set drag target and exchange dragStart and dragTarget  */
    function dropHandler(e: DragEvent<HTMLAnchorElement>, index: number) {
        dispatch(setDragTarget(index));
        dispatch(exchangeCards(fieldSetup.dragStart, fieldSetup.dragTarget));
    };

    return (
        <div className={`playing-formation-${formation}`}>
            {props.cardsList.map((card, index) => {
                const data = card.cardData;

                return (
                    <a
                        href={data ? undefined : '#cardList'}
                        key={index}
                        className={`playing-formation-${formation}__${data ? 'card' : 'empty-card'}`}
                        onClick={() => dispatch(choseCardPosition(index))}
                        draggable={true}
                        onDragStart={() => dispatch(setDragStart(index))}
                        onDragOver={dragOverHandler}
                        onDrop={e => dropHandler(e, index)}
                    >
                        {
                            data
                                ? <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                : null
                        }
                    </a>
                );
            })}
        </div>
    );
};
