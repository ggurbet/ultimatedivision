// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FieldControlsArea } from '@components/Field/FieldControlsArea';
import { FootballerCard } from '@components/Field/FootballerCard';

import { CardEditIdentificators } from '@/api/club';
import { RootState } from '@/app/store';
import { Card } from '@/card';
import { SquadCard } from '@/club';
import {
    cardSelectionVisibility,
    changeCardPosition,
    choosePosition,
    deleteCard,
    setDragStart,
    swapCards,
} from '@/app/store/actions/clubs';

import './index.scss';

export const FieldPlayingArea: React.FC = () => {
    const dispatch = useDispatch();

    const cards = useSelector((state: RootState) => state.cardsReducer.cardsPage.cards);
    const formation = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);
    const dragStartIndex = useSelector((state: RootState) => state.clubsReducer.options.dragStart);
    const club = useSelector((state: RootState) => state.clubsReducer.activeClub);
    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

    /** MouseMove event Position */
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
    /** This var created to not allow mouseUpEvent without Dragging before it */
    const [isDragging, handleDrag] = useState(false);
    /** Playing area position */
    const [playingAreaPosition, setplayingAreaPosition] = useState({ x: 0, y: 0 });

    const DEFAULT_VALUE = 0;
    const X_SCROLL_POINT = 0;
    const Y_SCROLL_POINT = 1200;
    const DELAY = 100;

    /** Gets playing area position */
    useEffect(() => {
        const playingArea = document.getElementById('playingArea');
        playingArea &&
            setplayingAreaPosition({
                x: playingArea.offsetLeft,
                y: playingArea.offsetTop,
            });
    }, []);
    const useMousePosition = (ev: any) => {
        setMousePosition({ x: ev.pageX, y: ev.pageY });
    };

    /** returns card data for card */
    function getCard(id: string) {
        return cards.find((card: Card) => card.id === id);
    }

    /** Add card position, and shows card selection */
    function handleClick(index: number) {
        dispatch(choosePosition(index));
        dispatch(cardSelectionVisibility(true));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
    }

    /** getting dragged card index and changing state to allow mouseUp */
    function dragStart(
        e: React.MouseEvent<HTMLDivElement>,
        index: number = DEFAULT_VALUE
    ): void {
        handleDrag(true);
        dispatch(setDragStart(index));
    }
    /** getting second drag index  and exchanging with first index*/
    function onMouseUp(
        e: React.MouseEvent<HTMLDivElement>,
        index: number = DEFAULT_VALUE
    ): void {
        e.stopPropagation();
        if (isDragging && dragStartIndex !== null) {
            const cards = club.squadCards;
            getCard(cards[index].cardId) ?
                dispatch(swapCards(
                    new CardEditIdentificators(squad.clubId, squad.id, cards[dragStartIndex].cardId, index),
                    new CardEditIdentificators(squad.clubId, squad.id, cards[index].cardId, dragStartIndex)
                ))
                :
                dispatch(changeCardPosition(
                    new CardEditIdentificators(squad.clubId, squad.id, cards[dragStartIndex].cardId, index),
                ));
        }

        dispatch(setDragStart());
        handleDrag(false);
    }

    /** when we release card not on target it just brings it on start position*/
    function mouseUpOnArea(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        dispatch(setDragStart());
    }

    /** deleting card when release beyond playing area */
    function removeFromArea() {
        if (isDragging && dragStartIndex) {
            dispatch(deleteCard(
                new CardEditIdentificators(squad.clubId, squad.id, club.squadCards[dragStartIndex].cardId, dragStartIndex))
            );
        }
        dispatch(setDragStart());
        handleDrag(false);
    }

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
                    {club.squadCards.map(
                        (fieldCard: SquadCard, index: number) => {
                            const card = getCard(fieldCard.cardId);
                            const isDragging = dragStartIndex === index;

                            return (
                                <div
                                    style={
                                        isDragging ? {
                                            left: mousePosition.x - playingAreaPosition.x,
                                            top: mousePosition.y - playingAreaPosition.y,
                                            transform: 'translate(-55%, -50%)',
                                            zIndex: 5,
                                            pointerEvents: 'none',
                                        }
                                            : undefined
                                    }
                                    key={index}
                                    className={`playing-area__${formation}__${card ? 'card' : 'empty-card'
                                    }`}
                                    onClick={() => handleClick(index)}
                                    onDragStart={(e) => dragStart(e, index)}
                                    onMouseUp={(e) => onMouseUp(e, index)}
                                    draggable={true}
                                >
                                    {card &&
                                        <FootballerCard
                                            card={card}
                                            index={index}
                                            place={'PlayingArea'}
                                        />
                                    }
                                </div>
                            );
                        })}
                </div>
                <div className={`playing-area__${formation}-shadows`}>
                    {club.squadCards.map(
                        (fieldCard: SquadCard, index: number) => {
                            const card = getCard(fieldCard.cardId);

                            return (
                                <div
                                    className={`playing-area__${formation}-shadows__card`}
                                    key={index}
                                >
                                    {card &&
                                        <img
                                            src={card.style.shadow}
                                            alt="card shadow"
                                            className={`playing-area__${formation}-shadows__shadow`}
                                        />
                                    }
                                </div>
                            );
                        }
                    )}
                </div>
            </div>
            <FieldControlsArea />
        </div>
    );
};
