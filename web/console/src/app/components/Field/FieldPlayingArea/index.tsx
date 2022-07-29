// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FieldControlsArea } from '@components/Field/FieldControlsArea';

import { CardEditIdentificators } from '@/api/club';
import { RootState } from '@/app/store';
import { Club, FormationsType, Squad, SquadCard } from '@/club';
import { deleteCard, setDragStart } from '@/app/store/actions/clubs';
import { FieldCardsFromation } from './FieldCardsFormation';
import { FieldCardsShadows } from './FieldCardsShadows';

import footballField from '@static/img/FieldPage/football_field.png';

import './index.scss';
import { Card } from '@/card';

export const FieldPlayingArea: React.FC = () => {
    const dispatch = useDispatch();

    const cards: Card[] = useSelector((state: RootState) => state.cardsReducer.cardsPage.cards);
    const formation: FormationsType = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);
    const dragStartIndex: number | null = useSelector((state: RootState) => state.clubsReducer.options.dragStart);
    const club: Club = useSelector((state: RootState) => state.clubsReducer.activeClub);
    const squad: Squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

    const [currentCard, setCurrentCard] = useState<Element | null>(null);
    /** MouseMove event Position */
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
    /** This var created to not allow mouseUpEvent without Dragging before it */
    const [isDragging, handleDrag] = useState(false);
    /** Playing area position */
    const [playingAreaPosition, setplayingAreaPosition] = useState({
        x: 0,
        y: 0,
    });

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

    /** Compares card id with default id */
    function isCardDefined(id: string) {
        const defaultId = '00000000-0000-0000-0000-000000000000';

        return id !== defaultId;
    }

    /** deleting card when release beyond playing area */
    function removeFromArea() {
        if (isDragging && dragStartIndex !== null) {
            dispatch(
                deleteCard(
                    new CardEditIdentificators(
                        squad.clubId,
                        squad.id,
                        club.squadCards[dragStartIndex].card.id,
                        dragStartIndex
                    )
                )
            );
        }
        dispatch(setDragStart());
        handleDrag(false);
    }

    /** Show/hide delete block, preventing scroll to cardSelection. */
    const handleVisibility = (e: React.MouseEvent<HTMLInputElement>): void => {
        e.stopPropagation();

        const target = e.target as Element;

        if (target && target.id !== currentCard?.id) {
            setCurrentCard(target);

            return;
        }

        setCurrentCard(null);
    };

    return (
        <div
            className="playing-area__wrapper"
            onMouseMove={(ev) => useMousePosition(ev)}
            onMouseUp={removeFromArea}
            style={isDragging ? { cursor: 'not-allowed' } : {}}
            onClick={handleVisibility}
        >
            <FieldControlsArea />
            <div className="playing-area" id="playingArea">
                <FieldCardsFromation
                    club={club}
                    currentCard={currentCard}
                    isCardDefined={isCardDefined}
                    isDragging={isDragging}
                    handleDrag={handleDrag}
                    mousePosition={mousePosition}
                    playingAreaPosition={playingAreaPosition}
                    dragStartIndex={dragStartIndex}
                    setCurrentCard={setCurrentCard}
                />
                <div className={`playing-area__${formation}-shadows`}>
                    <FieldCardsShadows club={club} isCardDefined={isCardDefined} />
                </div>
                <img src={footballField} className="playing-area__field" />
            </div>
        </div>
    );
};
