// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { CardEditIdentificators } from '@/api/club';
import { RootState } from '@/app/store';
import {
    cardSelectionVisibility,
    changeCardPosition,
    choosePosition,
    setDragStart,
    swapCards,
} from '@/app/store/actions/clubs';
import { Club, FormationsType, Squad, SquadCard } from '@/club';
import { FootballerCard } from '../../FootballerCard';

export const FieldCardsFromation: React.FC<{
    club: Club;
    currentCard: Element | null;
    isCardDefined: (params: string) => boolean;
    isDragging: boolean;
    handleDrag: Dispatch<SetStateAction<boolean>>;
    mousePosition: { x: number; y: number };
    playingAreaPosition: { x: number; y: number };
    dragStartIndex: number | null;
    setCurrentCard: (currentCard: Element | null) => void;
}> = ({
    club,
    isCardDefined,
    dragStartIndex,
    setCurrentCard,
    currentCard,
    isDragging,
    handleDrag,
    mousePosition,
    playingAreaPosition,
}) => {
    const dispatch = useDispatch();

    const formation: FormationsType = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);
    const squad: Squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

    const DEFAULT_VALUE: number = 0;
    const X_SCROLL_POINT: number = 0;
    const Y_SCROLL_POINT: number = 1800;
    const DELAY: number = 100;

    /** Add card position, and shows card selection */
    function handleClick(index: number, e: React.MouseEvent<HTMLDivElement>) {
        const target = e.target as Element;

        if (target.className.includes('empty')) {
            dispatch(choosePosition(index));
            dispatch(cardSelectionVisibility(true));
            setTimeout(() => {
                window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
            }, DELAY);
        }
    }

    /** getting dragged card index and changing state to allow mouseUp */
    function dragStart(e: React.MouseEvent<HTMLDivElement>, index: number = DEFAULT_VALUE): void {
        handleDrag(true);
        dispatch(setDragStart(index));
    }
    /** getting second drag index  and exchanging with first index*/
    function onMouseUp(e: React.MouseEvent<HTMLDivElement>, index: number = DEFAULT_VALUE): void {
        e.stopPropagation();
        if (isDragging && dragStartIndex !== null) {
            const cards = club.squadCards;
            isCardDefined(cards[index].card.id)
                ? dispatch(
                    swapCards(
                        new CardEditIdentificators(squad.clubId, squad.id, cards[dragStartIndex].card.id, index),
                        new CardEditIdentificators(squad.clubId, squad.id, cards[index].card.id, dragStartIndex)
                    )
                )
                : dispatch(
                    changeCardPosition(
                        new CardEditIdentificators(squad.clubId, squad.id, cards[dragStartIndex].card.id, index)
                    )
                );
        }

        dispatch(setDragStart());
        handleDrag(false);
    }

    /** when we release card not on target it just brings it on start position*/
    function mouseUpOnArea(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        dispatch(setDragStart());
    }

    return (
        <div
            style={dragStartIndex ? { cursor: 'grabbing' } : {}}
            className={`playing-area__${formation}`}
            onMouseUp={mouseUpOnArea}
        >
            {club.squadCards.map((fieldCard: SquadCard, index: number) => {
                const isDefined = isCardDefined(fieldCard.card.id);
                const isDragging = dragStartIndex === index;

                return (
                    <div
                        style={
                            isDragging
                                ? {
                                    left: mousePosition.x - playingAreaPosition.x,
                                    top: mousePosition.y - playingAreaPosition.y,
                                    transform: 'translate(-55%, -50%)',
                                    zIndex: 5,
                                    pointerEvents: 'none',
                                }
                                : undefined
                        }
                        key={index}
                        className={`playing-area__${formation}__${isDefined ? 'card' : 'empty-card'}`}
                        onClick={(e) => handleClick(index, e)}
                        onDragStart={(e) => dragStart(e, index)}
                        onMouseUp={(e) => onMouseUp(e, index)}
                        draggable={true}
                    >
                        {isDefined &&
                            <FootballerCard
                                card={fieldCard.card}
                                index={index}
                                place={'PlayingArea'}
                                setCurrentCard={setCurrentCard}
                                currentCard={currentCard}
                            />
                        }
                    </div>
                );
            })}
        </div>
    );
};
