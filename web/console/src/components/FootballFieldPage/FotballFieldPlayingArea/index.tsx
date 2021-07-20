/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { useEffect, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';

import { FootballFieldInformation } from '../FootballFieldInformation'
import { PlayingAreaFootballerCard } from '../PlayingAreaFootballerCard';

import { RootState } from '../../../store';
import { choseCardPosition, setDragStart, setDragTarget, exchangeCards, removeCard }
    from '../../../store/reducers/footballField';

import './index.scss';

export const FootballFieldPlayingArea: React.FC = () => {
    const formation = useSelector((state: RootState) => state.fieldReducer.options.formation);
    const dragStartIndex = useSelector((state: RootState) => state.fieldReducer.options.dragStart);

    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer);

    /** MouseMove event Position */
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
    /** This var created to not allow mouseUpEvent without Dragging before it */
    const [isDragging, handleDrag] = useState(false);
    /** outer padding of playingArea */
    const [outerOffset, handleOffset] = useState({ x: 0, y: 0 });

    /** with getBoundingClientRect() we gettins outer padding of playingArea on any screen width and scale */
    useEffect(() => {
        const playingArea = document.getElementById('playingArea');
        if (playingArea) {
            const position = playingArea.getBoundingClientRect();
            handleOffset({ x: position.x + 60, y: position.y + 100 });
        }
    },[])
    const useMousePosition = (ev: any) => {
        setMousePosition({ x: ev.pageX, y: ev.pageY });
    };

    /** getting dragged card index and changing state to allow mouseUp */
    function dragStart(e: any, index: number = 0): void {
        handleDrag(true);
        dispatch(setDragStart(index));
    }

    /** getting second drag index  and exchanging with first index*/
    function onMouseUp(e: any, index: number = 0): void {
        e.stopPropagation();

        if (isDragging && dragStartIndex) {
            dispatch(setDragTarget(index));
            dispatch(exchangeCards(dragStartIndex, fieldSetup.options.dragTarget));
        }
        dispatch(setDragTarget())
        dispatch(setDragStart());
        handleDrag(false);
    }

    /** when we release card not on target it just brings it on start position*/
    function mouseUpOnArea(e: any) {
        e.stopPropagation();
        dispatch(setDragStart());
    }


/** deleting card when release beyond playing area */
    function removeFromArea() {
        if (isDragging) {
            dispatch(removeCard(dragStartIndex));
            dispatch(setDragStart());
            handleDrag(false);
        }
    }

    const { x, y } = mousePosition;

    return (
        <div
            className="playing-area__wrapper"
            onMouseMove={(ev) => useMousePosition(ev)}
            onMouseUp={removeFromArea}
            style={isDragging ? { cursor: 'not-allowed' } : {}}
        >
            <div className="playing-area" id="playingArea">
                <div
                    style={dragStartIndex ? { cursor: 'grabbing' } : {}}
                    className={`playing-area__${formation}`}
                    onMouseUp={mouseUpOnArea}
                >
                    {fieldSetup.cardsList.map((card, index) => {
                        const data = card.cardData;
                        const equality = (dragStartIndex === index);
                        //TO DO: change style by some class to change style in card
                        return (
                            <a
                                style={
                                    equality
                                        ? { left: (x - outerOffset.x), top: (y - 330), zIndex: 5, pointerEvents: 'none' }
                                        : undefined
                                }
                                href={data ? undefined : '#cardList'}
                                key={index}
                                className={`playing-area__${formation}__${data ? 'card' : 'empty-card'}`}
                                onClick={(e) => dispatch(choseCardPosition(index))}
                                onDragStart={(e) => dragStart(e, index)}
                                onMouseUp={(e) => onMouseUp(e, index)}
                                draggable={true}
                            >
                                {
                                    data && <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                }
                            </a>
                        );
                    })}
                </div>
            </div>
            <FootballFieldInformation />
        </div>
    );
};
