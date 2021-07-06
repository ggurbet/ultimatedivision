/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingFormation_424.scss';
import { FootballField } from '../../../../../types/footballField';
import { useDispatch } from 'react-redux';
import { choseCardPosition }
    from '../../../../../store/reducers/footballField';
import { PlayingAreaFootballerCard }
    from '../../PlayingAreaFootballerCard/PlayingAreaFootballerCard';

export const PlayingFormation_424: React.FC<{ props: FootballField }> = ({ props }) => {
    const dispatch = useDispatch();

    return (
        <div className="playing-formation-424">
            {props.cardsList.map((card, index) => {
                const data = card.cardData;
                return (
                    <div
                        onClick={() => dispatch(choseCardPosition(index.toString()))}
                        key={index}
                        className="playing-formation-424__card"
                    >
                        {
                            data
                                ? <PlayingAreaFootballerCard card={data} place={'PlayingArea'} />
                                : <a
                                    href="#cardList"
                                    className="playing-formation-424__link"
                                >
                                </a>
                        }
                    </div>
                )
            })}
        </div>
    )
}
