// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '../../../store';

import './index.scss';

export const FootballerCardInformation: React.FC = () => {
    const FIRST_CARD_INDEX = 0;

    const overalInfo = useSelector((state: RootState) =>
        state.cardReducer[FIRST_CARD_INDEX].overalInfo);

    return (
        <div className="footballer-card-information">
            {overalInfo.map((item, index) =>
                <div className="footballer-card-information__item"
                    key={index}
                >
                    <div className="footballer-card-information__item-title">{item.label}</div>
                    <div className="footballer-card-information__item-value">
                        <>
                            {item.value}
                            <img
                                className="footballer-card-information__item-icon"
                                src={item.icon}
                                alt={item.icon && 'item icon'}
                            />
                        </>
                    </div>
                </div>,
            )}
        </div>
    );
};
