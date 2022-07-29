// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { Paginator } from '@components/common/Paginator';
import { PlayerCard } from '@components/common/PlayerCard';
import { FilterField } from '@/app/components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';

import { RootState } from '@/app/store';
import {
    fieldCards,
    getCurrentFieldCardsQueryParameters,
    createFieldCardsQueryParameters,
} from '@/app/store/actions/cards';
import { addCard, cardSelectionVisibility } from '@/app/store/actions/clubs';
import { CardEditIdentificators } from '@/api/club';
import { Card, CardsPage, CardsQueryParametersField } from '@/card';
import { Squad, SquadCard } from '@/club';

import './index.scss';

export const FieldCardSelection = () => {
    const dispatch = useDispatch();
    const squad: Squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);
    const squadCards: SquadCard[] = useSelector((state: RootState) => state.clubsReducer.activeClub.squadCards);

    const { cards, page }: CardsPage = useSelector((state: RootState) => state.cardsReducer.cardsPage);
    const { currentFieldCardsPage } = useSelector((state: RootState) => state.cardsReducer);
    const club = useSelector((state: RootState) => state.clubsReducer);

    const Y_SCROLL_POINT = 200;
    const X_SCROLL_POINT = 0;
    const DELAY = 10;

    const squadCardsIds = squadCards.map((card) => card.card.id);

    const fieldCardsQueryParameters = getCurrentFieldCardsQueryParameters();

    /** Add card to field, and hide card selection component */
    function addCardOnField(cardId: string, isSelected: boolean) {
        if (!isSelected) {
            dispatch(addCard(new CardEditIdentificators(squad.clubId, squad.id, cardId, club.options.chosedCard)));
            dispatch(cardSelectionVisibility(false));
            setTimeout(() => {
                window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
            }, DELAY);
        }
    }

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** Submits search by cards query parameters. */
    const submitSearch = async(cardsQueryParameters: CardsQueryParametersField[]) => {
        createFieldCardsQueryParameters(cardsQueryParameters);
        await dispatch(fieldCards(DEFAULT_PAGE_INDEX));
    };

    return (
        <div id="cardList" className="card-selection">
            <FilterField>
                <FilterByVersion submitSearch={submitSearch} cardsQueryParameters={fieldCardsQueryParameters} />
                <FilterByStats cardsQueryParameters={fieldCardsQueryParameters} submitSearch={submitSearch} />
                <FilterByPrice />
                <FilterByStatus />
            </FilterField>
            <div className="card-selection__list">
                {cards.map((card: Card, index: number) => {
                    const isSelected = squadCardsIds.includes(card.id);

                    return (
                        <div
                            key={index}
                            className={isSelected ? 'card-selection__card__selected' : 'card-selection__card'}
                            onClick={() => addCardOnField(card.id, isSelected)}
                        >
                            <div className={isSelected ? 'card-selection__card__selected__label' : ''}></div>
                            <PlayerCard id={card.id} className={'card-selection__card__picture'} />
                        </div>
                    );
                })}
            </div>
            <Paginator getCardsOnPage={fieldCards} itemsCount={page.totalCount} selectedPage={currentFieldCardsPage} />
        </div>
    );
};
