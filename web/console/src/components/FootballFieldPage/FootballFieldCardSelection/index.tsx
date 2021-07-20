/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../../store';
import { addCard } from '../../../store/reducers/footballField';

import { FilterField } from './FilterField';
import { PlayerCard } from '../../PlayerCard';


import './index.scss';
import { Paginator } from '../../Paginator';

export const FootballFieldCardSelection = () => {
    const cardList = useSelector((state: RootState) => state.cardReducer);
    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer);

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            <div className="card-selection__list">
                {cardList.map((card, index) =>
                    <a key={index} href="#playingArea" className="card-selection__card"
                        onClick={() => dispatch(addCard(card, fieldSetup.options.chosedCard))}
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
