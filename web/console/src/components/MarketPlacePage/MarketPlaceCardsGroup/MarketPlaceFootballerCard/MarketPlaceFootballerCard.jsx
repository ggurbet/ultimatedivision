/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { PropTypes } from 'prop-types';
import { NavLink } from 'react-router-dom';

import './MarketPlaceFootballerCard.scss';

export const MarketPlaceFootballerCard = ({ card }) => {

    return (
        <div className="marketplace-playerCard">
            <img className="marketplace-playerCard__background-type"
                src={card.mainInfo.backgroundType}
                alt="Player background type" />
            <img className="marketplace-playerCard__face-picture"
                src={card.mainInfo.facePicture}
                alt="Player face" />
            <NavLink to="/marketplace/card">
                <span className="marketplace-playerCard__name">
                    {card.overalInfo.name}
                </span>
            </NavLink>
            <ul className="marketplace-playerCard__list">
                {card.stats.map(
                    (property, index) => {
                        return (
                            <li
                                className="marketplace-playerCard__list__item"
                                key={index}>
                                {
                                    /**
                                    * get only average value of player's game property
                                    */
                                    `${property.average} ${property.title.slice(0,3)}`
                                }
                            </li>
                        );
                    }
                )}
            </ul>
            <div className="marketplace-playerCard__price">
                <img className="marketplace-playerCard__price__picture"
                    src={card.mainInfo.pricePicture}
                    alt="Player price" />
                <span className="marketplace-playerCard__price__current">
                    {card.mainInfo.price}
                </span>
                <img className="marketplace-playerCard__price__status"
                    src={card.mainInfo.priceStatus}
                    alt="Price status" />
            </div>
        </div>
    );
};

MarketPlaceFootballerCard.propTypes = {
    card: PropTypes.object.isRequired
};
