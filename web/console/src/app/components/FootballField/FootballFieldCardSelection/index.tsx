// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '@/app/store';
import { addCard, cardSelectionVisibility } from '@/app/store/actions/club';

import { listOfCards } from '@/app/store/actions/cards';

import { FilterField } from
    '@components/FootballField/FootballFieldCardSelection/FilterField';
import { PlayerCard } from '@components/common/PlayerCard';
import { Paginator } from '@components/common/Paginator';
import { Card } from '@/card';


import './index.scss';

export const FootballFieldCardSelection = () => {
    const { cards, page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);
    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.clubReducer);

    const Y_SCROLL_POINT = 200;
    const X_SCROLL_POINT = 0;
    const DELAY = 10;

    /** Add card to field, and hide card selection component */
    function handleClick(card: Card, index: number) {
        dispatch(addCard(card, index));
        dispatch(cardSelectionVisibility(false));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
    }

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            <div className="card-selection__list">
                {cards.map((card: Card, index: number) =>
                    <div
                        key={index}
                        className="card-selection__card"
                        onClick={() => handleClick(card, fieldSetup.options.chosedCard)}
                    >
                        <PlayerCard
                            card={card}
                            parentClassName={'card-selection__card'}
                        />
                    </div>,
                )}
            </div>
            <Paginator
                getCardsOnPage={listOfCards}
                pagesCount={page.pageCount}
                selectedPage={page.currentPage}
            />
        </div>
    );
};
