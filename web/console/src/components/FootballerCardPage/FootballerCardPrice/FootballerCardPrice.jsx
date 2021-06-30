/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import './FootballerCardPrice.scss';
import { Doughnut } from 'react-chartjs-2';
import { useSelector } from 'react-redux';

import currency from '../../../img/FootballerCardPage/currency.png';

/* eslint-disable */
export const FootballerCardPrice = () => {

    const priceData = useSelector(state => state.footballerCard[0]);
    const prpValue = priceData.price.prp.value;

    return (
        <div className="footballer-card-price">
            <div className="footballer-card-price__wrapper">
                <div className="footballer-card-price__diagram">
                    <p className="footballer-card-price__diagram-value">{`PRP: ${prpValue}%`}</p>
                    <Doughnut
                        data={{
                            datasets: [{
                                data: [prpValue, (100 - prpValue)],
                                backgroundColor: [
                                    `${priceData.price.color}`,
                                    '#5E5EAA'
                                ],
                                borderColor: [
                                    'transparent'
                                ],
                                cutout: '80%',
                                rotation: 90,
                                esponsive: true,
                                maintainAspectRatio: true
                            }],
                        }}
                    />
                </div>
                <div className="footballer-card-price__info-area">
                    <h2 className="footballer-card-price__price">
                        <>
                            {priceData.price.price.value}
                            <img
                                className="footballer-card-price__price-currency"
                                src={currency}
                                alt="currency img"
                            />
                        </>
                    </h2>
                    <div className="footballer-card-price__additional-info">
                        <div>
                            Price updated: <span
                                className="footballer-card-price__value"
                            >
                                {priceData.price.updated.value} mins ago
                            </span>
                        </div>
                        <div>
                            PR: <span
                                className="footballer-card-price__value"
                            >
                                {priceData.price.pr.value}
                            </span>
                        </div>
                    </div>

                </div>
            </div>

        </div>
    );
};
