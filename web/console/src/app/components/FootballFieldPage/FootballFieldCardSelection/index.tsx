//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../../store';
import { addCard, cardSelectionVisibility } from '@/app/store/reducers/footballField';

import { FilterField } from
    '@components/FootballFieldPage/FootballFieldCardSelection/FilterField';
import { PlayerCard } from '@components/PlayerCard';

import './index.scss';
import { Paginator } from '@components/Paginator';
import { Card } from '@/app/store/reducers/footballerCard';

export const FootballFieldCardSelection = () => {
    const cardList = useSelector((state: RootState) => state.cardReducer);
    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer);

    /** Add card to field, and hide card selection component */
    function handleClick(card: Card, index: number) {
        dispatch(addCard(card, index));
        dispatch(cardSelectionVisibility(false));
    }

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            <div className="card-selection__list">
                {cardList.map((card, index) =>
                    <a key={index} href="#playingArea" className="card-selection__card"
                        onClick={() => handleClick(card, fieldSetup.options.chosedCard)}
                    >
                        <PlayerCard
                            card={card}
                            parentClassName={'card-selection__card'}
                        />
                    </a>,
                )}
            </div>
            <Paginator itemCount={cardList.length} />
        </div>
    );
};
