// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { changeActiveClub, setFormation, startSearchingMatch } from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';
import { amountColumnsElements, Control } from '@/app/types/club';
import { Formations, SquadCard } from '@/club';
import { DropdownStyle } from '@/app/internal/dropdownStyle';
import { FieldDropdown } from './FieldDropdown';

import arrowIcon from '@static/img/FieldPage/arrow.svg';
import arrowActiveIcon from '@static/img/FieldPage/arrow-active.svg';

import './index.scss';

export const FieldControlsArea: React.FC = () => {
    const dispatch = useDispatch();

    const EMPTY_CARD_ID = '00000000-0000-0000-0000-000000000000';

    const [currentOption, setCurrentOption] = useState<null | Control>(null);
    const [optionVisibility, changeVisibility] = useState(false);
    const [activeComposition, setActiveComposition] = useState<string>('Composition 1');
    const [isPossibleToStartMatch, setIsPossibleToStartMatch] = useState<boolean>(true);

    const squadCards = useSelector((state: RootState) => state.clubsReducer.activeClub.squadCards);
    const formation = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);
    const clubs = useSelector((state: RootState) => state.clubsReducer.clubs);
    const activeClubId = useSelector((state: RootState) => state.clubsReducer.activeClub.id);

    const checkActiveElement = (item: Control) => item.title === currentOption?.title && optionVisibility;

    const CONTROLS_FIELDS = [
        new Control(
            '1',
            'club',
            changeActiveClub,
            clubs,
            amountColumnsElements['four-elements'],
            activeClubId,
            'id',
            'name',
            'club'
        ),
        new Control(
            '2',
            'squad',
            setActiveComposition,
            ['Composition 1', 'Composition 2', 'Composition 3', 'Composition 4', 'Composition 5'],
            amountColumnsElements['five-elements'],
            activeComposition
        ),
        new Control(
            '3',
            'formation',
            setFormation,
            ['4-4-2', '4-2-4', '4-2-2-2', '4-3-1-2', '4-3-3', '4-2-3-1', '4-3-2-1', '4-1-3-2', '5-3-2', '4-5-1'],
            amountColumnsElements['five-elements'],
            Formations[formation]
        ),
    ];

    const setCurrentControlsAreaOption = (item: Control) => {
        setCurrentOption(item);

        if (item.title !== currentOption?.title && optionVisibility) {
            changeVisibility(false);
        }
        changeVisibility((prev) => !prev);
    };

    /** shows matchFinder component */
    const showMatchFinder = () => {
        dispatch(startSearchingMatch(true));
        window.scrollTo({
            top: 0,
            behavior: 'smooth',
        });
    };

    const isDropdownActive = useMemo(
        () => currentOption !== null && optionVisibility,
        [currentOption, optionVisibility]
    );

    useEffect(() => {
        /** Function checks field cards and compare it with player cards array */
        function isPossibleToStart() {
            const emptyCard = squadCards.find((squadCard: SquadCard) => squadCard.card.id === EMPTY_CARD_ID);
            emptyCard ? setIsPossibleToStartMatch(false) : setIsPossibleToStartMatch(true);
        }
        isPossibleToStart();
    });

    return (
        <div className="field-controls">
            <div className="field-controls__wrapper">
                <div className="field-controls__settings">
                    {CONTROLS_FIELDS.map((item, index) =>
                        <div className="field-controls__settings__item" key={item.title}>
                            <div
                                className="field-controls__settings__item__heading"
                                onClick={() => setCurrentControlsAreaOption(item)}
                            >
                                <h4 className="field-controls__settings__item__title">{item.title}</h4>

                                <img
                                    className="field-controls__settings__item__image"
                                    src={checkActiveElement(item) ? arrowActiveIcon : arrowIcon}
                                    alt="triangle img"
                                    id={`triangle-${item.id}`}
                                    style={
                                        checkActiveElement(item)
                                            ? { transform: new DropdownStyle(true).triangleRotate }
                                            : {}
                                    }
                                />
                            </div>
                        </div>
                    )}
                </div>
                {isDropdownActive && <FieldDropdown option={currentOption} />}
            </div>
            <input
                type="button"
                value="Play"
                className="field-controls__play"
                onClick={showMatchFinder}
                disabled={isPossibleToStartMatch}
            />
        </div>
    );
};
