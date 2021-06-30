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

    const priceData = useSelector(state => state.fotballerCardPrice);
    const fields = priceData.fields;
    const prpValue = priceData.fields.prp.value;

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
                                    `${priceData.color}`,
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
                            {fields.price.value} <img src={currency} alt="currency img" />
                        </>
                    </h2>
                    <div className="footballer-card-price__additional-info">
                        <div>
                            Price updated: <span
                                className="footballer-card-price__value"
                            >
                                {fields.updated.value} mins ago
                            </span>
                        </div>
                        <div>
                            PR: <span
                                className="footballer-card-price__value"
                            >
                                {fields.pr.value}
                            </span>
                        </div>
                    </div>

                </div>
            </div>

        </div>
    );
};
